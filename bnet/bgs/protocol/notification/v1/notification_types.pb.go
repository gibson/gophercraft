// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: notification_types.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/Gophercraft/core/bnet/bgs/protocol"
	v1 "github.com/Gophercraft/core/bnet/bgs/protocol/account/v1"
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

type Target struct {
	Identity             *TargetIdentity `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Type                 *string         `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e08de6de44a1482, []int{0}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetIdentity() *TargetIdentity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Target) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

type TargetIdentity struct {
	Account              *v1.AccountId         `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	GameAccount          *v1.GameAccountHandle `protobuf:"bytes,2,opt,name=game_account,json=gameAccount" json:"game_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TargetIdentity) Reset()         { *m = TargetIdentity{} }
func (m *TargetIdentity) String() string { return proto.CompactTextString(m) }
func (*TargetIdentity) ProtoMessage()    {}
func (*TargetIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e08de6de44a1482, []int{1}
}

func (m *TargetIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetIdentity.Unmarshal(m, b)
}
func (m *TargetIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetIdentity.Marshal(b, m, deterministic)
}
func (m *TargetIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetIdentity.Merge(m, src)
}
func (m *TargetIdentity) XXX_Size() int {
	return xxx_messageInfo_TargetIdentity.Size(m)
}
func (m *TargetIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_TargetIdentity proto.InternalMessageInfo

func (m *TargetIdentity) GetAccount() *v1.AccountId {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *TargetIdentity) GetGameAccount() *v1.GameAccountHandle {
	if m != nil {
		return m.GameAccount
	}
	return nil
}

type Subscription struct {
	Target               []*Target    `protobuf:"bytes,1,rep,name=target" json:"target,omitempty"`
	Subscriber           *v1.Identity `protobuf:"bytes,2,opt,name=subscriber" json:"subscriber,omitempty"`
	DeliveryRequired     *bool        `protobuf:"varint,3,opt,name=delivery_required,json=deliveryRequired" json:"delivery_required,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e08de6de44a1482, []int{2}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetTarget() []*Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Subscription) GetSubscriber() *v1.Identity {
	if m != nil {
		return m.Subscriber
	}
	return nil
}

// Deprecated: Do not use.
func (m *Subscription) GetDeliveryRequired() bool {
	if m != nil && m.DeliveryRequired != nil {
		return *m.DeliveryRequired
	}
	return false
}

type Notification struct {
	SenderId             *protocol.EntityId    `protobuf:"bytes,1,opt,name=sender_id,json=senderId" json:"sender_id,omitempty"`
	TargetId             *protocol.EntityId    `protobuf:"bytes,2,req,name=target_id,json=targetId" json:"target_id,omitempty"`
	Type                 *string               `protobuf:"bytes,3,req,name=type" json:"type,omitempty"`
	Attribute            []*protocol.Attribute `protobuf:"bytes,4,rep,name=attribute" json:"attribute,omitempty"`
	SenderAccountId      *protocol.EntityId    `protobuf:"bytes,5,opt,name=sender_account_id,json=senderAccountId" json:"sender_account_id,omitempty"`
	TargetAccountId      *protocol.EntityId    `protobuf:"bytes,6,opt,name=target_account_id,json=targetAccountId" json:"target_account_id,omitempty"`
	SenderBattleTag      *string               `protobuf:"bytes,7,opt,name=sender_battle_tag,json=senderBattleTag" json:"sender_battle_tag,omitempty"`
	TargetBattleTag      *string               `protobuf:"bytes,8,opt,name=target_battle_tag,json=targetBattleTag" json:"target_battle_tag,omitempty"`
	ForwardingIdentity   *v1.Identity          `protobuf:"bytes,10,opt,name=forwarding_identity,json=forwardingIdentity" json:"forwarding_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e08de6de44a1482, []int{3}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetSenderId() *protocol.EntityId {
	if m != nil {
		return m.SenderId
	}
	return nil
}

func (m *Notification) GetTargetId() *protocol.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *Notification) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Notification) GetAttribute() []*protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *Notification) GetSenderAccountId() *protocol.EntityId {
	if m != nil {
		return m.SenderAccountId
	}
	return nil
}

func (m *Notification) GetTargetAccountId() *protocol.EntityId {
	if m != nil {
		return m.TargetAccountId
	}
	return nil
}

func (m *Notification) GetSenderBattleTag() string {
	if m != nil && m.SenderBattleTag != nil {
		return *m.SenderBattleTag
	}
	return ""
}

func (m *Notification) GetTargetBattleTag() string {
	if m != nil && m.TargetBattleTag != nil {
		return *m.TargetBattleTag
	}
	return ""
}

func (m *Notification) GetForwardingIdentity() *v1.Identity {
	if m != nil {
		return m.ForwardingIdentity
	}
	return nil
}

func init() {
	proto.RegisterType((*Target)(nil), "bgs.protocol.notification.v1.Target")
	proto.RegisterType((*TargetIdentity)(nil), "bgs.protocol.notification.v1.TargetIdentity")
	proto.RegisterType((*Subscription)(nil), "bgs.protocol.notification.v1.Subscription")
	proto.RegisterType((*Notification)(nil), "bgs.protocol.notification.v1.Notification")
}

func init() { proto.RegisterFile("notification_types.proto", fileDescriptor_2e08de6de44a1482) }

var fileDescriptor_2e08de6de44a1482 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x65, 0xa7, 0xa4, 0xc9, 0x26, 0xa2, 0x74, 0x2b, 0xc0, 0xaa, 0x38, 0x18, 0x8b, 0x43,
	0x54, 0x21, 0xbb, 0x2d, 0xe2, 0xd6, 0x4b, 0x2c, 0x10, 0xcd, 0x01, 0x0e, 0x4b, 0x4f, 0x5c, 0xac,
	0xb5, 0x77, 0xb3, 0x5d, 0x29, 0xb1, 0xcd, 0x7a, 0x1c, 0x94, 0x87, 0xe1, 0x5d, 0x78, 0x2b, 0xae,
	0x95, 0xbd, 0xeb, 0x7f, 0x87, 0x46, 0xb9, 0x59, 0x33, 0xf3, 0xfb, 0xe6, 0x9b, 0xd9, 0x31, 0x72,
	0xd2, 0x0c, 0xe4, 0x5a, 0x26, 0x14, 0x64, 0x96, 0x46, 0xb0, 0xcf, 0x79, 0xe1, 0xe7, 0x2a, 0x83,
	0x0c, 0xbf, 0x8b, 0x85, 0xf9, 0x4c, 0xb2, 0x8d, 0xdf, 0x2f, 0xf3, 0x77, 0x37, 0x97, 0x17, 0x34,
	0x49, 0xb2, 0x32, 0x85, 0x3e, 0x72, 0xf9, 0x9a, 0x02, 0x28, 0x19, 0x97, 0xc0, 0x07, 0x61, 0xcc,
	0x53, 0x90, 0xb0, 0x1f, 0xc4, 0xce, 0x54, 0x9e, 0xf4, 0x03, 0xde, 0x1a, 0x8d, 0x1f, 0xa8, 0x12,
	0x1c, 0xf0, 0x3d, 0x9a, 0x48, 0xa6, 0x11, 0xc7, 0x72, 0xad, 0xc5, 0xec, 0xf6, 0xa3, 0x7f, 0xc8,
	0x8b, 0xaf, 0xb9, 0x95, 0x61, 0x48, 0x4b, 0x63, 0x8c, 0x4e, 0xaa, 0x16, 0x8e, 0xed, 0x5a, 0x8b,
	0x29, 0xa9, 0xbf, 0xbd, 0xbf, 0x16, 0x7a, 0x39, 0x04, 0xf0, 0x1d, 0x3a, 0x35, 0xd3, 0x98, 0x7e,
	0xde, 0xb0, 0x9f, 0x49, 0x56, 0xad, 0x96, 0xfa, 0x73, 0xc5, 0x48, 0x83, 0xe0, 0xef, 0x68, 0x2e,
	0xe8, 0x96, 0x47, 0x8d, 0x84, 0x5d, 0x4b, 0x5c, 0x3d, 0x2b, 0xf1, 0x8d, 0x6e, 0xb9, 0x91, 0xb9,
	0xa7, 0x29, 0xdb, 0x70, 0x32, 0x13, 0x5d, 0xc8, 0xfb, 0x67, 0xa1, 0xf9, 0xcf, 0x32, 0x2e, 0x12,
	0x25, 0xf3, 0x6a, 0x40, 0x7c, 0x87, 0xc6, 0x50, 0xfb, 0x75, 0x2c, 0x77, 0xb4, 0x98, 0xdd, 0x7e,
	0x38, 0x66, 0x19, 0xc4, 0x30, 0x78, 0x89, 0x50, 0xa1, 0xd5, 0x62, 0xae, 0x8c, 0xb7, 0xf7, 0xcf,
	0x7a, 0x6b, 0x77, 0xd8, 0x83, 0x70, 0x80, 0xce, 0x19, 0xdf, 0xc8, 0x1d, 0x57, 0xfb, 0x48, 0xf1,
	0xdf, 0xa5, 0x54, 0x9c, 0x39, 0x23, 0xd7, 0x5a, 0x4c, 0x42, 0xdb, 0xb1, 0xc8, 0xab, 0x26, 0x49,
	0x4c, 0xce, 0xfb, 0x3f, 0x42, 0xf3, 0x1f, 0x3d, 0x5b, 0xf8, 0x13, 0x9a, 0x16, 0x3c, 0x65, 0x5c,
	0x45, 0x92, 0x99, 0x15, 0xbf, 0x19, 0x7a, 0xf8, 0x5a, 0xb7, 0x5d, 0x31, 0x32, 0xd1, 0x85, 0x2b,
	0x56, 0x41, 0x7a, 0x86, 0x0a, 0xb2, 0x5d, 0xfb, 0x10, 0x04, 0xe6, 0x41, 0xdb, 0x17, 0x1f, 0xb9,
	0x76, 0xf3, 0xe2, 0xf8, 0x33, 0x9a, 0xb6, 0x77, 0xe9, 0x9c, 0xd4, 0x3b, 0x7c, 0x3b, 0x14, 0x5a,
	0x36, 0x69, 0xd2, 0x55, 0xe2, 0x10, 0x9d, 0x1b, 0xd3, 0xcd, 0xa9, 0x4b, 0xe6, 0xbc, 0x38, 0x68,
	0xfe, 0x4c, 0x03, 0xed, 0x91, 0x54, 0x1a, 0x66, 0x86, 0x9e, 0xc6, 0xf8, 0xb0, 0x86, 0x06, 0x3a,
	0x8d, 0xab, 0xd6, 0x47, 0x4c, 0x01, 0x36, 0x3c, 0x02, 0x2a, 0x9c, 0xd3, 0xfa, 0xa2, 0x4d, 0xbf,
	0xb0, 0x8e, 0x3f, 0x50, 0x51, 0xd5, 0x9a, 0x7e, 0xbd, 0xda, 0x89, 0xae, 0xd5, 0x89, 0xae, 0x96,
	0xa0, 0x8b, 0x75, 0xa6, 0xfe, 0x50, 0xc5, 0x64, 0x2a, 0xa2, 0xf6, 0x8f, 0x43, 0xc7, 0x9e, 0x08,
	0xee, 0xe8, 0x26, 0x16, 0x7e, 0xf9, 0x15, 0x0a, 0x09, 0x8f, 0x65, 0xec, 0x27, 0xd9, 0x36, 0x28,
	0xca, 0x9c, 0xab, 0xfc, 0xfa, 0x1a, 0x02, 0x91, 0xe5, 0x8f, 0x5c, 0x25, 0x8a, 0xae, 0x21, 0x88,
	0x53, 0x0e, 0x41, 0x2c, 0x8a, 0xa0, 0x51, 0x0f, 0xfa, 0x27, 0x1c, 0xec, 0x6e, 0x9e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x42, 0xc1, 0x96, 0x1d, 0x94, 0x04, 0x00, 0x00,
}
