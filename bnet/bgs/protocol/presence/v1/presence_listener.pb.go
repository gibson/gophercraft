// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: presence_listener.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gibson/gophercraft/bnet/bgs/protocol"
	v1 "github.com/gibson/gophercraft/bnet/bgs/protocol/account/v1"
	proto "github.com/golang/protobuf/proto"
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

type SubscribeNotification struct {
	SubscriberId         *v1.AccountId    `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId" json:"subscriber_id,omitempty"`
	State                []*PresenceState `protobuf:"bytes,2,rep,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SubscribeNotification) Reset()         { *m = SubscribeNotification{} }
func (m *SubscribeNotification) String() string { return proto.CompactTextString(m) }
func (*SubscribeNotification) ProtoMessage()    {}
func (*SubscribeNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_926ea8405948847c, []int{0}
}

func (m *SubscribeNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeNotification.Unmarshal(m, b)
}
func (m *SubscribeNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeNotification.Marshal(b, m, deterministic)
}
func (m *SubscribeNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeNotification.Merge(m, src)
}
func (m *SubscribeNotification) XXX_Size() int {
	return xxx_messageInfo_SubscribeNotification.Size(m)
}
func (m *SubscribeNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeNotification.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeNotification proto.InternalMessageInfo

func (m *SubscribeNotification) GetSubscriberId() *v1.AccountId {
	if m != nil {
		return m.SubscriberId
	}
	return nil
}

func (m *SubscribeNotification) GetState() []*PresenceState {
	if m != nil {
		return m.State
	}
	return nil
}

type StateChangedNotification struct {
	SubscriberId         *v1.AccountId    `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId" json:"subscriber_id,omitempty"`
	State                []*PresenceState `protobuf:"bytes,2,rep,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StateChangedNotification) Reset()         { *m = StateChangedNotification{} }
func (m *StateChangedNotification) String() string { return proto.CompactTextString(m) }
func (*StateChangedNotification) ProtoMessage()    {}
func (*StateChangedNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_926ea8405948847c, []int{1}
}

func (m *StateChangedNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateChangedNotification.Unmarshal(m, b)
}
func (m *StateChangedNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateChangedNotification.Marshal(b, m, deterministic)
}
func (m *StateChangedNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateChangedNotification.Merge(m, src)
}
func (m *StateChangedNotification) XXX_Size() int {
	return xxx_messageInfo_StateChangedNotification.Size(m)
}
func (m *StateChangedNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_StateChangedNotification.DiscardUnknown(m)
}

var xxx_messageInfo_StateChangedNotification proto.InternalMessageInfo

func (m *StateChangedNotification) GetSubscriberId() *v1.AccountId {
	if m != nil {
		return m.SubscriberId
	}
	return nil
}

func (m *StateChangedNotification) GetState() []*PresenceState {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeNotification)(nil), "bgs.protocol.presence.v1.SubscribeNotification")
	proto.RegisterType((*StateChangedNotification)(nil), "bgs.protocol.presence.v1.StateChangedNotification")
}

func init() { proto.RegisterFile("presence_listener.proto", fileDescriptor_926ea8405948847c) }

var fileDescriptor_926ea8405948847c = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x51, 0xcd, 0x4b, 0xc3, 0x30,
	0x14, 0xa7, 0x15, 0x45, 0x32, 0xbf, 0x88, 0x8a, 0xb5, 0xa7, 0xd1, 0x8b, 0xc3, 0x49, 0xb2, 0xf5,
	0xe0, 0x4d, 0xf1, 0x83, 0x21, 0x03, 0x59, 0xc7, 0x7a, 0xf3, 0x32, 0xda, 0x34, 0xeb, 0x02, 0x33,
	0x09, 0x49, 0x3a, 0xf0, 0xba, 0xa3, 0x7f, 0x80, 0x57, 0xc1, 0xff, 0x72, 0x37, 0xd9, 0xd6, 0xce,
	0x0d, 0xec, 0xbc, 0x7a, 0x7b, 0xbc, 0xf7, 0x7e, 0x1f, 0xef, 0xf7, 0xc0, 0x99, 0x54, 0x54, 0x53,
	0x4e, 0x68, 0x7f, 0xc4, 0xb4, 0xa1, 0x9c, 0x2a, 0x24, 0x95, 0x30, 0x02, 0x3a, 0x71, 0xaa, 0x17,
	0x25, 0x11, 0x23, 0x54, 0x6c, 0xa1, 0x71, 0xd3, 0x3d, 0x59, 0x42, 0xcc, 0x9b, 0xa4, 0xf9, 0x92,
	0x7b, 0x1c, 0x11, 0x22, 0x32, 0x6e, 0xd6, 0x9a, 0x87, 0x4a, 0x92, 0xd5, 0x86, 0xf7, 0x69, 0x81,
	0xd3, 0x30, 0x8b, 0x35, 0x51, 0x2c, 0xa6, 0x1d, 0x61, 0xd8, 0x80, 0x91, 0xc8, 0x30, 0xc1, 0xe1,
	0x13, 0xd8, 0xd7, 0xc5, 0x40, 0xf5, 0x59, 0xe2, 0x58, 0x55, 0xab, 0x56, 0xf1, 0x3d, 0xb4, 0xe6,
	0x23, 0x17, 0x41, 0xe3, 0x26, 0xba, 0x5f, 0x94, 0xed, 0xa4, 0xb7, 0xf7, 0x03, 0x6c, 0x27, 0xf0,
	0x06, 0x6c, 0x6b, 0x13, 0x19, 0xea, 0xd8, 0xd5, 0xad, 0x5a, 0xc5, 0xbf, 0x40, 0x65, 0x87, 0xa0,
	0x6e, 0x5e, 0x87, 0xb3, 0xf5, 0xde, 0x02, 0xe5, 0x7d, 0x59, 0xc0, 0x99, 0x37, 0x1e, 0x87, 0x11,
	0x4f, 0x69, 0xf2, 0x1f, 0x4d, 0xfa, 0x1f, 0x36, 0x38, 0x2a, 0x06, 0xcf, 0xf9, 0xdf, 0x60, 0x04,
	0x2a, 0x01, 0x5f, 0x86, 0x0b, 0x71, 0x39, 0xe7, 0xaf, 0x1f, 0x70, 0xcf, 0xd7, 0x01, 0x9d, 0xa0,
	0xdf, 0x6b, 0x85, 0xdd, 0xa0, 0x13, 0xb6, 0xbc, 0x9d, 0xc9, 0xb4, 0x6e, 0xef, 0x5a, 0x30, 0x05,
	0x07, 0x01, 0x5f, 0x4d, 0x07, 0xfa, 0x1b, 0x54, 0x4a, 0x52, 0xfc, 0x5b, 0xc8, 0x76, 0xaf, 0x27,
	0xd3, 0xfa, 0x15, 0xb8, 0x8c, 0x39, 0x35, 0x9b, 0x63, 0x29, 0xae, 0x7f, 0x9f, 0x1b, 0x7c, 0xb8,
	0x7b, 0xb9, 0x4d, 0x99, 0x19, 0x66, 0x31, 0x22, 0xe2, 0x15, 0xeb, 0x4c, 0x52, 0x25, 0x1b, 0x0d,
	0x83, 0x53, 0x21, 0x87, 0x54, 0x11, 0x15, 0x0d, 0x0c, 0x9e, 0x91, 0xe2, 0x38, 0xd5, 0xb8, 0x20,
	0xc6, 0x05, 0x31, 0x1e, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xca, 0xd5, 0x37, 0x11,
	0x03, 0x00, 0x00,
}
