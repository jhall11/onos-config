// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package synchronizer synchronizes configurations down to devices
package synchronizer

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/onosproject/onos-config/pkg/events"
	"github.com/onosproject/onos-config/pkg/southbound"
	"github.com/onosproject/onos-config/pkg/southbound/topocache"
	"github.com/onosproject/onos-config/pkg/store"
	"github.com/onosproject/onos-config/pkg/utils"
	"github.com/openconfig/gnmi/client"
	"github.com/openconfig/gnmi/proto/gnmi"
	"log"
)

// Synchronizer enables proper configuring of a device based on store events and cache of operational data
type Synchronizer struct {
	context.Context
	*store.ChangeStore
	*topocache.Device
	deviceConfigChan     <-chan events.ConfigEvent
	operationalStateChan chan<- events.OperationalStateEvent
	key                  southbound.DeviceID
	query                client.Query
	operationalCache     map[string]string
}

// New Build a new Synchronizer given the parameters, starts the connection with the device and polls the capabilities
func New(context context.Context, changeStore *store.ChangeStore, device *topocache.Device,
	deviceCfgChan <-chan events.ConfigEvent, opStateChan chan<- events.OperationalStateEvent) (*Synchronizer, error) {
	sync := &Synchronizer{
		Context:              context,
		ChangeStore:          changeStore,
		Device:               device,
		deviceConfigChan:     deviceCfgChan,
		operationalStateChan: opStateChan,
		operationalCache:     make(map[string]string),
	}
	log.Println("Connecting to", sync.Device.Addr, "over gNMI")
	target := southbound.Target{}
	key, err := target.ConnectTarget(context, *sync.Device)
	sync.key = key
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(device.Addr, "Connected over gNMI")

	// Get the device capabilities
	capResponse, capErr := target.CapabilitiesWithString(context, "")
	if capErr != nil {
		log.Println(sync.Device.Addr, "Capabilities", err)
		return nil, err
	}

	log.Println(sync.Device.Addr, "Capabilities", capResponse)
	return sync, nil
}

// Devicesync is a go routine that listens out for configuration events specific
// to a device and propagates them downwards through southbound interface
func (sync *Synchronizer) syncNbConfiguration() {

	for deviceConfigEvent := range sync.deviceConfigChan {
		change := sync.ChangeStore.Store[deviceConfigEvent.ChangeID()]
		err := change.IsValid()
		if err != nil {
			log.Println("Event discarded because change is invalid", err)
			continue
		}
		gnmiChange, parseError := change.GnmiChange()

		if parseError != nil {
			log.Println("Parsing error for Gnmi change ", parseError)
			continue
		}

		log.Println("Change formatted to gNMI setRequest", gnmiChange)
		target, err := southbound.GetTarget(sync.key)
		if err != nil {
			log.Println(err)
			continue
		}
		setResponse, err := target.Set(sync.Context, gnmiChange)
		if err != nil {
			log.Println("SetResponse ", err)
			continue
		}
		log.Println(sync.Device.Addr, "SetResponse", setResponse)

	}
}

func (sync Synchronizer) syncOperationalState() error {

	target, err := southbound.GetTarget(sync.key)

	if err != nil {
		log.Println("Can't find target for key", sync.key)
		return err
	}

	notifications := make([]*gnmi.Notification, 0)

	requestState := &gnmi.GetRequest{
		Type: gnmi.GetRequest_STATE,
	}

	responseState, err := target.Get(target.Ctx, requestState)

	if err != nil {
		log.Println("Error in requesting read-only state to target", sync.key, err)
		return err
	}

	notifications = append(notifications, responseState.Notification...)

	requestOperational := &gnmi.GetRequest{
		Type: gnmi.GetRequest_OPERATIONAL,
	}

	responseOperational, err := target.Get(target.Ctx, requestOperational)

	if err != nil {
		log.Println("Error in requesting read-only state to target", sync.key, err)
		return err
	}

	notifications = append(notifications, responseOperational.Notification...)

	subscribePaths := make([][]string, 0)

	for _, notification := range notifications {
		for _, update := range notification.Update {
			pathStr := utils.StrPath(update.Path)
			val := utils.StrVal(update.Val)
			sync.operationalCache[pathStr] = val
			subscribePaths = append(subscribePaths, utils.SplitPath(pathStr))
		}

	}

	options := &southbound.SubscribeOptions{
		UpdatesOnly:       false,
		Prefix:            "",
		Mode:              "stream",
		StreamMode:        "target_defined",
		SampleInterval:    15,
		HeartbeatInterval: 15,
		Paths:             subscribePaths,
		Origin:            "",
	}

	req, err := southbound.NewSubscribeRequest(options)
	if err != nil {
		return err
	}
	return target.Subscribe(sync.Context, req, sync.handler)

}

func (sync *Synchronizer) handler(msg proto.Message) error {

	_, err := southbound.GetTarget(sync.key)
	if err != nil {
		return fmt.Errorf("target not connected %#v", msg)
	}

	resp, ok := msg.(*gnmi.SubscribeResponse)
	if !ok {
		return fmt.Errorf("failed to type assert message %#v", msg)
	}
	switch v := resp.Response.(type) {
	default:
		return fmt.Errorf("unknown response %T: %s", v, v)
	case *gnmi.SubscribeResponse_Error:
		return fmt.Errorf("error in response: %s", v)
	case *gnmi.SubscribeResponse_SyncResponse:
		if sync.query.Type == client.Poll || sync.query.Type == client.Once {
			return client.ErrStopReading
		}
	case *gnmi.SubscribeResponse_Update:
		eventValues := make(map[string]string)
		notification := v.Update
		for _, update := range notification.Update {
			if update.Path == nil {
				return fmt.Errorf("invalid nil path in update: %v", update)
			}
			pathStr := utils.StrPath(update.Path)
			val := utils.StrVal(update.Val)
			eventValues[pathStr] = val
			sync.operationalCache[pathStr] = val
		}
		for _, del := range notification.Delete {
			if del.Elem == nil {
				return fmt.Errorf("invalid nil path in update: %v", del)
			}
			//deletedPaths := delete.Elem
			pathStr := utils.StrPathElem(del.Elem)
			eventValues[pathStr] = ""
			sync.operationalCache[pathStr] = ""
		}
		sync.operationalStateChan <- events.CreateOperationalStateEvent(sync.Device.Target, eventValues)
	}
	return nil
}
