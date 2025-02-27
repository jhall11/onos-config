// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/northbound/proto/admin.proto

// Package admin defines the administrative and diagnostic gRPC interfaces.

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// NetworkChangesRequest is a message for specifying GetNetworkChanges query parameters.
type NetworkChangesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkChangesRequest) Reset()         { *m = NetworkChangesRequest{} }
func (m *NetworkChangesRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkChangesRequest) ProtoMessage()    {}
func (*NetworkChangesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{0}
}

func (m *NetworkChangesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkChangesRequest.Unmarshal(m, b)
}
func (m *NetworkChangesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkChangesRequest.Marshal(b, m, deterministic)
}
func (m *NetworkChangesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkChangesRequest.Merge(m, src)
}
func (m *NetworkChangesRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkChangesRequest.Size(m)
}
func (m *NetworkChangesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkChangesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkChangesRequest proto.InternalMessageInfo

// ConfigChange is a descriptor of a submitted configuration change targeted as a single device.
type ConfigChange struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigChange) Reset()         { *m = ConfigChange{} }
func (m *ConfigChange) String() string { return proto.CompactTextString(m) }
func (*ConfigChange) ProtoMessage()    {}
func (*ConfigChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{1}
}

func (m *ConfigChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigChange.Unmarshal(m, b)
}
func (m *ConfigChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigChange.Marshal(b, m, deterministic)
}
func (m *ConfigChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigChange.Merge(m, src)
}
func (m *ConfigChange) XXX_Size() int {
	return xxx_messageInfo_ConfigChange.Size(m)
}
func (m *ConfigChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigChange.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigChange proto.InternalMessageInfo

func (m *ConfigChange) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConfigChange) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// NetChange is a descriptor of a configuration change submitted via gNMI.
type NetChange struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	User                 string               `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Changes              []*ConfigChange      `protobuf:"bytes,4,rep,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NetChange) Reset()         { *m = NetChange{} }
func (m *NetChange) String() string { return proto.CompactTextString(m) }
func (*NetChange) ProtoMessage()    {}
func (*NetChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{2}
}

func (m *NetChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetChange.Unmarshal(m, b)
}
func (m *NetChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetChange.Marshal(b, m, deterministic)
}
func (m *NetChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetChange.Merge(m, src)
}
func (m *NetChange) XXX_Size() int {
	return xxx_messageInfo_NetChange.Size(m)
}
func (m *NetChange) XXX_DiscardUnknown() {
	xxx_messageInfo_NetChange.DiscardUnknown(m)
}

var xxx_messageInfo_NetChange proto.InternalMessageInfo

func (m *NetChange) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *NetChange) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetChange) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *NetChange) GetChanges() []*ConfigChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

// RegisterRequest carries data for registering or unregistering a YANG model.
type RegisterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{3}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

// RegisterResponse carries status of YANG model registration or unregistration.
type RegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{4}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

// RollbackRequest carries the name of a network config to rollback. If there
// are subsequent changes to any of the devices in that config, the rollback will
// be rejected. If no name is given the last network change will be rolled back.
type RollbackRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RollbackRequest) Reset()         { *m = RollbackRequest{} }
func (m *RollbackRequest) String() string { return proto.CompactTextString(m) }
func (*RollbackRequest) ProtoMessage()    {}
func (*RollbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{5}
}

func (m *RollbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollbackRequest.Unmarshal(m, b)
}
func (m *RollbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollbackRequest.Marshal(b, m, deterministic)
}
func (m *RollbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollbackRequest.Merge(m, src)
}
func (m *RollbackRequest) XXX_Size() int {
	return xxx_messageInfo_RollbackRequest.Size(m)
}
func (m *RollbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RollbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RollbackRequest proto.InternalMessageInfo

func (m *RollbackRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RollbackRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type RollbackResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RollbackResponse) Reset()         { *m = RollbackResponse{} }
func (m *RollbackResponse) String() string { return proto.CompactTextString(m) }
func (*RollbackResponse) ProtoMessage()    {}
func (*RollbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{6}
}

func (m *RollbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollbackResponse.Unmarshal(m, b)
}
func (m *RollbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollbackResponse.Marshal(b, m, deterministic)
}
func (m *RollbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollbackResponse.Merge(m, src)
}
func (m *RollbackResponse) XXX_Size() int {
	return xxx_messageInfo_RollbackResponse.Size(m)
}
func (m *RollbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RollbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RollbackResponse proto.InternalMessageInfo

func (m *RollbackResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// DeviceInfo is a record of various device-pertinent information.
type DeviceInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	User                 string   `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	CaPath               string   `protobuf:"bytes,7,opt,name=caPath,proto3" json:"caPath,omitempty"`
	CertPath             string   `protobuf:"bytes,8,opt,name=certPath,proto3" json:"certPath,omitempty"`
	KeyPath              string   `protobuf:"bytes,9,opt,name=keyPath,proto3" json:"keyPath,omitempty"`
	Plain                bool     `protobuf:"varint,10,opt,name=plain,proto3" json:"plain,omitempty"`
	Insecure             bool     `protobuf:"varint,11,opt,name=insecure,proto3" json:"insecure,omitempty"`
	Timeout              int64    `protobuf:"varint,12,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceInfo) Reset()         { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()    {}
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{7}
}

func (m *DeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceInfo.Unmarshal(m, b)
}
func (m *DeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceInfo.Marshal(b, m, deterministic)
}
func (m *DeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfo.Merge(m, src)
}
func (m *DeviceInfo) XXX_Size() int {
	return xxx_messageInfo_DeviceInfo.Size(m)
}
func (m *DeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfo proto.InternalMessageInfo

func (m *DeviceInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DeviceInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *DeviceInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DeviceInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *DeviceInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DeviceInfo) GetCaPath() string {
	if m != nil {
		return m.CaPath
	}
	return ""
}

func (m *DeviceInfo) GetCertPath() string {
	if m != nil {
		return m.CertPath
	}
	return ""
}

func (m *DeviceInfo) GetKeyPath() string {
	if m != nil {
		return m.KeyPath
	}
	return ""
}

func (m *DeviceInfo) GetPlain() bool {
	if m != nil {
		return m.Plain
	}
	return false
}

func (m *DeviceInfo) GetInsecure() bool {
	if m != nil {
		return m.Insecure
	}
	return false
}

func (m *DeviceInfo) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// DeviceResponse carries the status of the add/remove operation.
type DeviceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceResponse) Reset()         { *m = DeviceResponse{} }
func (m *DeviceResponse) String() string { return proto.CompactTextString(m) }
func (*DeviceResponse) ProtoMessage()    {}
func (*DeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{8}
}

func (m *DeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceResponse.Unmarshal(m, b)
}
func (m *DeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceResponse.Marshal(b, m, deterministic)
}
func (m *DeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceResponse.Merge(m, src)
}
func (m *DeviceResponse) XXX_Size() int {
	return xxx_messageInfo_DeviceResponse.Size(m)
}
func (m *DeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceResponse proto.InternalMessageInfo

// GetDevicesRequest carries devices query information.
type GetDevicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDevicesRequest) Reset()         { *m = GetDevicesRequest{} }
func (m *GetDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*GetDevicesRequest) ProtoMessage()    {}
func (*GetDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91573d9ad3811b0, []int{9}
}

func (m *GetDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDevicesRequest.Unmarshal(m, b)
}
func (m *GetDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDevicesRequest.Marshal(b, m, deterministic)
}
func (m *GetDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDevicesRequest.Merge(m, src)
}
func (m *GetDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_GetDevicesRequest.Size(m)
}
func (m *GetDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDevicesRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NetworkChangesRequest)(nil), "proto.NetworkChangesRequest")
	proto.RegisterType((*ConfigChange)(nil), "proto.ConfigChange")
	proto.RegisterType((*NetChange)(nil), "proto.NetChange")
	proto.RegisterType((*RegisterRequest)(nil), "proto.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "proto.RegisterResponse")
	proto.RegisterType((*RollbackRequest)(nil), "proto.RollbackRequest")
	proto.RegisterType((*RollbackResponse)(nil), "proto.RollbackResponse")
	proto.RegisterType((*DeviceInfo)(nil), "proto.DeviceInfo")
	proto.RegisterType((*DeviceResponse)(nil), "proto.DeviceResponse")
	proto.RegisterType((*GetDevicesRequest)(nil), "proto.GetDevicesRequest")
}

func init() { proto.RegisterFile("pkg/northbound/proto/admin.proto", fileDescriptor_a91573d9ad3811b0) }

var fileDescriptor_a91573d9ad3811b0 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xba, 0x6e, 0xdd, 0xee, 0xca, 0xd6, 0x7a, 0x6c, 0xb3, 0x22, 0x24, 0xaa, 0x3c, 0xf5,
	0x01, 0xda, 0xa9, 0xbc, 0x20, 0x21, 0x04, 0x63, 0x20, 0xe0, 0x81, 0x81, 0x02, 0xfb, 0x00, 0x37,
	0xb9, 0x4b, 0xa3, 0x36, 0x76, 0xb0, 0xdd, 0x4e, 0xfb, 0x0d, 0xfe, 0x81, 0x1f, 0xe0, 0x43, 0xf8,
	0x26, 0xe4, 0xd8, 0x4e, 0xd6, 0x95, 0x17, 0xc4, 0x53, 0x7c, 0xce, 0xbd, 0xf7, 0xdc, 0x93, 0xeb,
	0x6b, 0x18, 0x94, 0xf3, 0x6c, 0xcc, 0x85, 0xd4, 0xb3, 0xa9, 0x58, 0xf2, 0x74, 0x5c, 0x4a, 0xa1,
	0xc5, 0x98, 0xa5, 0x45, 0xce, 0x47, 0xd5, 0x99, 0x6c, 0x57, 0x9f, 0xf0, 0x71, 0x26, 0x44, 0xb6,
	0x40, 0x9b, 0x30, 0x5d, 0x5e, 0x8f, 0x75, 0x5e, 0xa0, 0xd2, 0xac, 0x28, 0x6d, 0x5e, 0x74, 0x0a,
	0xc7, 0x97, 0xa8, 0x6f, 0x84, 0x9c, 0x5f, 0xcc, 0x18, 0xcf, 0x50, 0xc5, 0xf8, 0x7d, 0x89, 0x4a,
	0x47, 0x13, 0xe8, 0x5e, 0x08, 0x7e, 0x9d, 0x67, 0x96, 0x27, 0x07, 0xd0, 0xca, 0x53, 0x1a, 0x0c,
	0x82, 0xe1, 0x5e, 0xdc, 0xca, 0x53, 0x42, 0xa0, 0x3d, 0x63, 0x6a, 0x46, 0x5b, 0x15, 0x53, 0x9d,
	0xa3, 0x1f, 0x01, 0xec, 0x5d, 0xa2, 0x76, 0x15, 0x23, 0x68, 0x9b, 0x6e, 0x55, 0xcd, 0xfe, 0x24,
	0x1c, 0x59, 0x2b, 0x23, 0x6f, 0x65, 0xf4, 0xcd, 0x5b, 0x89, 0xab, 0x3c, 0xa3, 0xc8, 0x59, 0x81,
	0x5e, 0xd1, 0x9c, 0x0d, 0xb7, 0x54, 0x28, 0xe9, 0x96, 0xe5, 0xcc, 0x99, 0x3c, 0x85, 0x4e, 0x62,
	0xbd, 0xd2, 0xf6, 0x60, 0x6b, 0xb8, 0x3f, 0x39, 0xb2, 0x9a, 0xa3, 0xbb, 0x7e, 0x63, 0x9f, 0x13,
	0xf5, 0xe1, 0x30, 0xc6, 0x2c, 0x57, 0x1a, 0xa5, 0xff, 0x37, 0x02, 0xbd, 0x86, 0x52, 0xa5, 0xe0,
	0x0a, 0xa3, 0x57, 0x70, 0x18, 0x8b, 0xc5, 0x62, 0xca, 0x92, 0xb9, 0x4b, 0xab, 0x0d, 0x05, 0x77,
	0x0c, 0x51, 0xe8, 0x24, 0xa2, 0x28, 0x90, 0x6b, 0xe7, 0xd3, 0xc3, 0xe8, 0x09, 0xf4, 0x1a, 0x01,
	0x2b, 0x6a, 0xb2, 0x0b, 0x54, 0x8a, 0x65, 0x5e, 0xc4, 0xc3, 0xe8, 0x57, 0x0b, 0xe0, 0x2d, 0xae,
	0xf2, 0x04, 0x3f, 0xf2, 0x6b, 0xb1, 0x31, 0x5d, 0x0a, 0x1d, 0x96, 0xa6, 0x12, 0x95, 0xf2, 0x6d,
	0x1c, 0x24, 0x27, 0xb0, 0xa3, 0x99, 0xcc, 0x50, 0xbb, 0x99, 0x38, 0x64, 0x2a, 0x56, 0x28, 0x55,
	0x2e, 0x38, 0x6d, 0xdb, 0x0a, 0x07, 0xeb, 0x19, 0x6e, 0xdf, 0x99, 0x61, 0x08, 0xbb, 0x25, 0x53,
	0xea, 0x46, 0xc8, 0x94, 0xee, 0x54, 0x7c, 0x8d, 0x4d, 0x87, 0x84, 0x7d, 0x61, 0x7a, 0x46, 0x3b,
	0xb6, 0x83, 0x45, 0xa6, 0x26, 0x41, 0xa9, 0xab, 0xc8, 0xae, 0xad, 0xf1, 0xd8, 0x74, 0x9f, 0xe3,
	0x6d, 0x15, 0xda, 0xb3, 0xdd, 0x1d, 0x24, 0x0f, 0x61, 0xbb, 0x5c, 0xb0, 0x9c, 0x53, 0x18, 0x04,
	0xc3, 0xdd, 0xd8, 0x02, 0xa3, 0x95, 0x73, 0x85, 0xc9, 0x52, 0x22, 0xdd, 0xaf, 0x02, 0x35, 0x36,
	0x5a, 0x66, 0x1f, 0xc4, 0x52, 0xd3, 0xee, 0x20, 0x18, 0x6e, 0xc5, 0x1e, 0x46, 0x3d, 0x38, 0xb0,
	0x33, 0xab, 0x6f, 0xed, 0x08, 0xfa, 0xef, 0x51, 0x5b, 0xd2, 0xaf, 0xee, 0xe4, 0x67, 0x0b, 0xba,
	0xe7, 0xe6, 0x2d, 0x7c, 0x45, 0x69, 0x02, 0xe4, 0x35, 0x3c, 0xf0, 0xf7, 0xfd, 0x49, 0xa4, 0xb8,
	0x20, 0x27, 0x6e, 0x63, 0xee, 0x2d, 0x46, 0x78, 0xba, 0xc1, 0xbb, 0x8b, 0x7c, 0x03, 0x87, 0x57,
	0x5c, 0xfe, 0x9f, 0xc6, 0xbb, 0xca, 0xeb, 0xfa, 0x6b, 0x23, 0x8f, 0x5c, 0xf6, 0x5f, 0x1f, 0x61,
	0xd8, 0x6b, 0xa2, 0x36, 0x72, 0x16, 0x90, 0x0f, 0x70, 0xec, 0xf7, 0x6c, 0xad, 0xa8, 0x31, 0xb4,
	0xbe, 0xc6, 0x8d, 0xa1, 0x7b, 0xdb, 0x39, 0xf9, 0x1d, 0xc0, 0x89, 0xdf, 0xc1, 0x15, 0x72, 0x2d,
	0xe4, 0xad, 0x9f, 0xd8, 0x4b, 0xe8, 0x9f, 0xa7, 0xe9, 0x67, 0x79, 0x55, 0xa6, 0x4c, 0xa3, 0x4d,
	0x22, 0x7d, 0x27, 0xd4, 0xec, 0x6d, 0x78, 0xbc, 0x46, 0xd5, 0xbf, 0xfa, 0x1c, 0xba, 0x31, 0x16,
	0x62, 0xf5, 0xef, 0x95, 0x2f, 0x00, 0x9a, 0x0b, 0x25, 0xd4, 0x25, 0x6d, 0xdc, 0x71, 0xb8, 0xa9,
	0x78, 0x16, 0x4c, 0x77, 0x2a, 0xee, 0xd9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xeb, 0x95,
	0xf0, 0x1f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	// RegisterModel adds the specified YANG model to the list of supported models.
	RegisterModel(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// UnregisterModel removes the specified YANG model from the list of supported models.
	UnregisterModel(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// GetNetworkChanges returns a stream of network changes submitted via gNMI.
	GetNetworkChanges(ctx context.Context, in *NetworkChangesRequest, opts ...grpc.CallOption) (AdminService_GetNetworkChangesClient, error)
	// RollbackNetworkChange rolls back the specified network change (or the latest one).
	RollbackNetworkChange(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*RollbackResponse, error)
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) RegisterModel(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.AdminService/RegisterModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UnregisterModel(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.AdminService/UnregisterModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetNetworkChanges(ctx context.Context, in *NetworkChangesRequest, opts ...grpc.CallOption) (AdminService_GetNetworkChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AdminService_serviceDesc.Streams[0], "/proto.AdminService/GetNetworkChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminServiceGetNetworkChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AdminService_GetNetworkChangesClient interface {
	Recv() (*NetChange, error)
	grpc.ClientStream
}

type adminServiceGetNetworkChangesClient struct {
	grpc.ClientStream
}

func (x *adminServiceGetNetworkChangesClient) Recv() (*NetChange, error) {
	m := new(NetChange)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminServiceClient) RollbackNetworkChange(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*RollbackResponse, error) {
	out := new(RollbackResponse)
	err := c.cc.Invoke(ctx, "/proto.AdminService/RollbackNetworkChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	// RegisterModel adds the specified YANG model to the list of supported models.
	RegisterModel(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// UnregisterModel removes the specified YANG model from the list of supported models.
	UnregisterModel(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// GetNetworkChanges returns a stream of network changes submitted via gNMI.
	GetNetworkChanges(*NetworkChangesRequest, AdminService_GetNetworkChangesServer) error
	// RollbackNetworkChange rolls back the specified network change (or the latest one).
	RollbackNetworkChange(context.Context, *RollbackRequest) (*RollbackResponse, error)
}

// UnimplementedAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (*UnimplementedAdminServiceServer) RegisterModel(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterModel not implemented")
}
func (*UnimplementedAdminServiceServer) UnregisterModel(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterModel not implemented")
}
func (*UnimplementedAdminServiceServer) GetNetworkChanges(req *NetworkChangesRequest, srv AdminService_GetNetworkChangesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNetworkChanges not implemented")
}
func (*UnimplementedAdminServiceServer) RollbackNetworkChange(ctx context.Context, req *RollbackRequest) (*RollbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackNetworkChange not implemented")
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_RegisterModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RegisterModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AdminService/RegisterModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RegisterModel(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UnregisterModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UnregisterModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AdminService/UnregisterModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UnregisterModel(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetNetworkChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NetworkChangesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServiceServer).GetNetworkChanges(m, &adminServiceGetNetworkChangesServer{stream})
}

type AdminService_GetNetworkChangesServer interface {
	Send(*NetChange) error
	grpc.ServerStream
}

type adminServiceGetNetworkChangesServer struct {
	grpc.ServerStream
}

func (x *adminServiceGetNetworkChangesServer) Send(m *NetChange) error {
	return x.ServerStream.SendMsg(m)
}

func _AdminService_RollbackNetworkChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RollbackNetworkChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AdminService/RollbackNetworkChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RollbackNetworkChange(ctx, req.(*RollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterModel",
			Handler:    _AdminService_RegisterModel_Handler,
		},
		{
			MethodName: "UnregisterModel",
			Handler:    _AdminService_UnregisterModel_Handler,
		},
		{
			MethodName: "RollbackNetworkChange",
			Handler:    _AdminService_RollbackNetworkChange_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNetworkChanges",
			Handler:       _AdminService_GetNetworkChanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/northbound/proto/admin.proto",
}

// DeviceInventoryServiceClient is the client API for DeviceInventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceInventoryServiceClient interface {
	// AddOrUpdateDevice adds a new or updates an existing new device in the device inventory.
	// The address field is required.
	AddOrUpdateDevice(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*DeviceResponse, error)
	// RemoveDevice removes a device to the device inventory.
	// Just the address field is required.
	RemoveDevice(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*DeviceResponse, error)
	// GetDevices returns a stream of devices in the device inventory.
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (DeviceInventoryService_GetDevicesClient, error)
}

type deviceInventoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceInventoryServiceClient(cc *grpc.ClientConn) DeviceInventoryServiceClient {
	return &deviceInventoryServiceClient{cc}
}

func (c *deviceInventoryServiceClient) AddOrUpdateDevice(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*DeviceResponse, error) {
	out := new(DeviceResponse)
	err := c.cc.Invoke(ctx, "/proto.DeviceInventoryService/AddOrUpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceInventoryServiceClient) RemoveDevice(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*DeviceResponse, error) {
	out := new(DeviceResponse)
	err := c.cc.Invoke(ctx, "/proto.DeviceInventoryService/RemoveDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceInventoryServiceClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (DeviceInventoryService_GetDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DeviceInventoryService_serviceDesc.Streams[0], "/proto.DeviceInventoryService/GetDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &deviceInventoryServiceGetDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DeviceInventoryService_GetDevicesClient interface {
	Recv() (*DeviceInfo, error)
	grpc.ClientStream
}

type deviceInventoryServiceGetDevicesClient struct {
	grpc.ClientStream
}

func (x *deviceInventoryServiceGetDevicesClient) Recv() (*DeviceInfo, error) {
	m := new(DeviceInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DeviceInventoryServiceServer is the server API for DeviceInventoryService service.
type DeviceInventoryServiceServer interface {
	// AddOrUpdateDevice adds a new or updates an existing new device in the device inventory.
	// The address field is required.
	AddOrUpdateDevice(context.Context, *DeviceInfo) (*DeviceResponse, error)
	// RemoveDevice removes a device to the device inventory.
	// Just the address field is required.
	RemoveDevice(context.Context, *DeviceInfo) (*DeviceResponse, error)
	// GetDevices returns a stream of devices in the device inventory.
	GetDevices(*GetDevicesRequest, DeviceInventoryService_GetDevicesServer) error
}

// UnimplementedDeviceInventoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceInventoryServiceServer struct {
}

func (*UnimplementedDeviceInventoryServiceServer) AddOrUpdateDevice(ctx context.Context, req *DeviceInfo) (*DeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateDevice not implemented")
}
func (*UnimplementedDeviceInventoryServiceServer) RemoveDevice(ctx context.Context, req *DeviceInfo) (*DeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDevice not implemented")
}
func (*UnimplementedDeviceInventoryServiceServer) GetDevices(req *GetDevicesRequest, srv DeviceInventoryService_GetDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}

func RegisterDeviceInventoryServiceServer(s *grpc.Server, srv DeviceInventoryServiceServer) {
	s.RegisterService(&_DeviceInventoryService_serviceDesc, srv)
}

func _DeviceInventoryService_AddOrUpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceInventoryServiceServer).AddOrUpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceInventoryService/AddOrUpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceInventoryServiceServer).AddOrUpdateDevice(ctx, req.(*DeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceInventoryService_RemoveDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceInventoryServiceServer).RemoveDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceInventoryService/RemoveDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceInventoryServiceServer).RemoveDevice(ctx, req.(*DeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceInventoryService_GetDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeviceInventoryServiceServer).GetDevices(m, &deviceInventoryServiceGetDevicesServer{stream})
}

type DeviceInventoryService_GetDevicesServer interface {
	Send(*DeviceInfo) error
	grpc.ServerStream
}

type deviceInventoryServiceGetDevicesServer struct {
	grpc.ServerStream
}

func (x *deviceInventoryServiceGetDevicesServer) Send(m *DeviceInfo) error {
	return x.ServerStream.SendMsg(m)
}

var _DeviceInventoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DeviceInventoryService",
	HandlerType: (*DeviceInventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrUpdateDevice",
			Handler:    _DeviceInventoryService_AddOrUpdateDevice_Handler,
		},
		{
			MethodName: "RemoveDevice",
			Handler:    _DeviceInventoryService_RemoveDevice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDevices",
			Handler:       _DeviceInventoryService_GetDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/northbound/proto/admin.proto",
}
