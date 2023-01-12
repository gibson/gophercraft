// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: presence_types.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/Gophercraft/core/bnet/bgs/protocol"
	v1 "github.com/Gophercraft/core/bnet/bgs/protocol/channel/v1"
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

type FieldOperation_OperationType int32

const (
	FieldOperation_SET   FieldOperation_OperationType = 0
	FieldOperation_CLEAR FieldOperation_OperationType = 1
)

var FieldOperation_OperationType_name = map[int32]string{
	0: "SET",
	1: "CLEAR",
}

var FieldOperation_OperationType_value = map[string]int32{
	"SET":   0,
	"CLEAR": 1,
}

func (x FieldOperation_OperationType) Enum() *FieldOperation_OperationType {
	p := new(FieldOperation_OperationType)
	*p = x
	return p
}

func (x FieldOperation_OperationType) String() string {
	return proto.EnumName(FieldOperation_OperationType_name, int32(x))
}

func (x *FieldOperation_OperationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldOperation_OperationType_value, data, "FieldOperation_OperationType")
	if err != nil {
		return err
	}
	*x = FieldOperation_OperationType(value)
	return nil
}

func (FieldOperation_OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d60f47575e6b3dd5, []int{3, 0}
}

type RichPresenceLocalizationKey struct {
	Program              *uint32  `protobuf:"fixed32,1,req,name=program" json:"program,omitempty"`
	Stream               *uint32  `protobuf:"fixed32,2,req,name=stream" json:"stream,omitempty"`
	LocalizationId       *uint32  `protobuf:"varint,3,req,name=localization_id,json=localizationId" json:"localization_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RichPresenceLocalizationKey) Reset()         { *m = RichPresenceLocalizationKey{} }
func (m *RichPresenceLocalizationKey) String() string { return proto.CompactTextString(m) }
func (*RichPresenceLocalizationKey) ProtoMessage()    {}
func (*RichPresenceLocalizationKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60f47575e6b3dd5, []int{0}
}

func (m *RichPresenceLocalizationKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RichPresenceLocalizationKey.Unmarshal(m, b)
}
func (m *RichPresenceLocalizationKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RichPresenceLocalizationKey.Marshal(b, m, deterministic)
}
func (m *RichPresenceLocalizationKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RichPresenceLocalizationKey.Merge(m, src)
}
func (m *RichPresenceLocalizationKey) XXX_Size() int {
	return xxx_messageInfo_RichPresenceLocalizationKey.Size(m)
}
func (m *RichPresenceLocalizationKey) XXX_DiscardUnknown() {
	xxx_messageInfo_RichPresenceLocalizationKey.DiscardUnknown(m)
}

var xxx_messageInfo_RichPresenceLocalizationKey proto.InternalMessageInfo

func (m *RichPresenceLocalizationKey) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *RichPresenceLocalizationKey) GetStream() uint32 {
	if m != nil && m.Stream != nil {
		return *m.Stream
	}
	return 0
}

func (m *RichPresenceLocalizationKey) GetLocalizationId() uint32 {
	if m != nil && m.LocalizationId != nil {
		return *m.LocalizationId
	}
	return 0
}

type FieldKey struct {
	Program              *uint32  `protobuf:"varint,1,req,name=program" json:"program,omitempty"`
	Group                *uint32  `protobuf:"varint,2,req,name=group" json:"group,omitempty"`
	Field                *uint32  `protobuf:"varint,3,req,name=field" json:"field,omitempty"`
	UniqueId             *uint64  `protobuf:"varint,4,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldKey) Reset()         { *m = FieldKey{} }
func (m *FieldKey) String() string { return proto.CompactTextString(m) }
func (*FieldKey) ProtoMessage()    {}
func (*FieldKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60f47575e6b3dd5, []int{1}
}

func (m *FieldKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldKey.Unmarshal(m, b)
}
func (m *FieldKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldKey.Marshal(b, m, deterministic)
}
func (m *FieldKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldKey.Merge(m, src)
}
func (m *FieldKey) XXX_Size() int {
	return xxx_messageInfo_FieldKey.Size(m)
}
func (m *FieldKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldKey.DiscardUnknown(m)
}

var xxx_messageInfo_FieldKey proto.InternalMessageInfo

func (m *FieldKey) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *FieldKey) GetGroup() uint32 {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return 0
}

func (m *FieldKey) GetField() uint32 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *FieldKey) GetUniqueId() uint64 {
	if m != nil && m.UniqueId != nil {
		return *m.UniqueId
	}
	return 0
}

type Field struct {
	Key                  *FieldKey         `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value                *protocol.Variant `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60f47575e6b3dd5, []int{2}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetKey() *FieldKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Field) GetValue() *protocol.Variant {
	if m != nil {
		return m.Value
	}
	return nil
}

type FieldOperation struct {
	Field                *Field                        `protobuf:"bytes,1,req,name=field" json:"field,omitempty"`
	Operation            *FieldOperation_OperationType `protobuf:"varint,2,opt,name=operation,enum=bgs.protocol.presence.v1.FieldOperation_OperationType,def=0" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FieldOperation) Reset()         { *m = FieldOperation{} }
func (m *FieldOperation) String() string { return proto.CompactTextString(m) }
func (*FieldOperation) ProtoMessage()    {}
func (*FieldOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60f47575e6b3dd5, []int{3}
}

func (m *FieldOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldOperation.Unmarshal(m, b)
}
func (m *FieldOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldOperation.Marshal(b, m, deterministic)
}
func (m *FieldOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldOperation.Merge(m, src)
}
func (m *FieldOperation) XXX_Size() int {
	return xxx_messageInfo_FieldOperation.Size(m)
}
func (m *FieldOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FieldOperation proto.InternalMessageInfo

const Default_FieldOperation_Operation FieldOperation_OperationType = FieldOperation_SET

func (m *FieldOperation) GetField() *Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *FieldOperation) GetOperation() FieldOperation_OperationType {
	if m != nil && m.Operation != nil {
		return *m.Operation
	}
	return Default_FieldOperation_Operation
}

type PresenceState struct {
	EntityId             *protocol.EntityId `protobuf:"bytes,1,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	FieldOperation       []*FieldOperation  `protobuf:"bytes,2,rep,name=field_operation,json=fieldOperation" json:"field_operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PresenceState) Reset()         { *m = PresenceState{} }
func (m *PresenceState) String() string { return proto.CompactTextString(m) }
func (*PresenceState) ProtoMessage()    {}
func (*PresenceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60f47575e6b3dd5, []int{4}
}

func (m *PresenceState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresenceState.Unmarshal(m, b)
}
func (m *PresenceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresenceState.Marshal(b, m, deterministic)
}
func (m *PresenceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresenceState.Merge(m, src)
}
func (m *PresenceState) XXX_Size() int {
	return xxx_messageInfo_PresenceState.Size(m)
}
func (m *PresenceState) XXX_DiscardUnknown() {
	xxx_messageInfo_PresenceState.DiscardUnknown(m)
}

var xxx_messageInfo_PresenceState proto.InternalMessageInfo

func (m *PresenceState) GetEntityId() *protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *PresenceState) GetFieldOperation() []*FieldOperation {
	if m != nil {
		return m.FieldOperation
	}
	return nil
}

type ChannelState struct {
	EntityId             *protocol.EntityId `protobuf:"bytes,1,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	FieldOperation       []*FieldOperation  `protobuf:"bytes,2,rep,name=field_operation,json=fieldOperation" json:"field_operation,omitempty"`
	Healing              *bool              `protobuf:"varint,3,opt,name=healing" json:"healing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChannelState) Reset()         { *m = ChannelState{} }
func (m *ChannelState) String() string { return proto.CompactTextString(m) }
func (*ChannelState) ProtoMessage()    {}
func (*ChannelState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60f47575e6b3dd5, []int{5}
}

func (m *ChannelState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelState.Unmarshal(m, b)
}
func (m *ChannelState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelState.Marshal(b, m, deterministic)
}
func (m *ChannelState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelState.Merge(m, src)
}
func (m *ChannelState) XXX_Size() int {
	return xxx_messageInfo_ChannelState.Size(m)
}
func (m *ChannelState) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelState.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelState proto.InternalMessageInfo

func (m *ChannelState) GetEntityId() *protocol.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *ChannelState) GetFieldOperation() []*FieldOperation {
	if m != nil {
		return m.FieldOperation
	}
	return nil
}

func (m *ChannelState) GetHealing() bool {
	if m != nil && m.Healing != nil {
		return *m.Healing
	}
	return false
}

var E_ChannelState_Presence = &proto.ExtensionDesc{
	ExtendedType:  (*v1.ChannelState)(nil),
	ExtensionType: (*ChannelState)(nil),
	Field:         101,
	Name:          "bgs.protocol.presence.v1.ChannelState.presence",
	Tag:           "bytes,101,opt,name=presence",
	Filename:      "presence_types.proto",
}

func init() {
	proto.RegisterEnum("bgs.protocol.presence.v1.FieldOperation_OperationType", FieldOperation_OperationType_name, FieldOperation_OperationType_value)
	proto.RegisterType((*RichPresenceLocalizationKey)(nil), "bgs.protocol.presence.v1.RichPresenceLocalizationKey")
	proto.RegisterType((*FieldKey)(nil), "bgs.protocol.presence.v1.FieldKey")
	proto.RegisterType((*Field)(nil), "bgs.protocol.presence.v1.Field")
	proto.RegisterType((*FieldOperation)(nil), "bgs.protocol.presence.v1.FieldOperation")
	proto.RegisterType((*PresenceState)(nil), "bgs.protocol.presence.v1.PresenceState")
	proto.RegisterExtension(E_ChannelState_Presence)
	proto.RegisterType((*ChannelState)(nil), "bgs.protocol.presence.v1.ChannelState")
}

func init() { proto.RegisterFile("presence_types.proto", fileDescriptor_d60f47575e6b3dd5) }

var fileDescriptor_d60f47575e6b3dd5 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xcd, 0x6e, 0xd3, 0x4e,
	0x10, 0xff, 0x3b, 0x69, 0xfe, 0x49, 0x26, 0x24, 0xad, 0x96, 0xb6, 0xb2, 0xda, 0x03, 0x91, 0x11,
	0x60, 0x09, 0xc9, 0x6e, 0xcd, 0xc7, 0x81, 0x03, 0x02, 0xaa, 0x20, 0x45, 0x54, 0x02, 0xb6, 0x15,
	0x12, 0x5c, 0xa2, 0x8d, 0x33, 0x71, 0x16, 0x1c, 0x7b, 0xd9, 0xac, 0x23, 0xc2, 0x8b, 0x70, 0xe0,
	0x81, 0x78, 0x2d, 0xe4, 0xb5, 0x37, 0xb5, 0x25, 0xaa, 0x72, 0xe4, 0x36, 0x33, 0x3b, 0xbf, 0x2f,
	0x6b, 0x0c, 0xfb, 0x42, 0xe2, 0x0a, 0x93, 0x10, 0x27, 0x6a, 0x23, 0x70, 0xe5, 0x09, 0x99, 0xaa,
	0x94, 0xd8, 0xd3, 0xa8, 0x2c, 0xc3, 0x34, 0xf6, 0xcc, 0x8a, 0xb7, 0x3e, 0x3d, 0x3a, 0x60, 0x4a,
	0x49, 0x3e, 0xcd, 0x54, 0x0d, 0x70, 0x44, 0x30, 0x51, 0x5c, 0x6d, 0x6a, 0xb3, 0xdb, 0xe1, 0x82,
	0x25, 0x09, 0xc6, 0xd5, 0xa1, 0xf3, 0x0d, 0x8e, 0x29, 0x0f, 0x17, 0xef, 0x4a, 0xca, 0xf3, 0x34,
	0x64, 0x31, 0xff, 0xce, 0x14, 0x4f, 0x93, 0x37, 0xb8, 0x21, 0x36, 0xb4, 0x85, 0x4c, 0x23, 0xc9,
	0x96, 0xb6, 0x35, 0x6c, 0xb8, 0x6d, 0x6a, 0x5a, 0x72, 0x08, 0xff, 0xaf, 0x94, 0x44, 0xb6, 0xb4,
	0x1b, 0xfa, 0xa1, 0xec, 0xc8, 0x03, 0xd8, 0x8d, 0x2b, 0x24, 0x13, 0x3e, 0xb3, 0x9b, 0xc3, 0x86,
	0xdb, 0xa7, 0x83, 0xea, 0x78, 0x3c, 0x73, 0x96, 0xd0, 0x79, 0xcd, 0x31, 0x9e, 0xfd, 0x41, 0xa6,
	0x7f, 0x25, 0xb3, 0x0f, 0xad, 0x48, 0xa6, 0x99, 0xd0, 0x2a, 0x7d, 0x5a, 0x34, 0xf9, 0x74, 0x9e,
	0x63, 0x4b, 0xea, 0xa2, 0x21, 0xc7, 0xd0, 0xcd, 0x12, 0xfe, 0x35, 0xc3, 0x5c, 0x74, 0x67, 0x68,
	0xb9, 0x3b, 0xb4, 0x53, 0x0c, 0xc6, 0x33, 0xe7, 0x33, 0xb4, 0xb4, 0x1c, 0x79, 0x0c, 0xcd, 0x2f,
	0xb8, 0xd1, 0x3a, 0xbd, 0xc0, 0xf1, 0xae, 0xfb, 0xb2, 0x9e, 0x31, 0x47, 0xf3, 0x75, 0xf2, 0x10,
	0x5a, 0x6b, 0x16, 0x67, 0xa8, 0x7d, 0xf4, 0x82, 0x83, 0x3a, 0xee, 0x03, 0x93, 0x9c, 0x25, 0x8a,
	0x16, 0x3b, 0xce, 0x2f, 0x0b, 0x06, 0x1a, 0xfe, 0x56, 0xa0, 0xd4, 0x79, 0xc9, 0x13, 0xe3, 0xb8,
	0xd0, 0xbd, 0x73, 0x83, 0xae, 0x89, 0xf4, 0x11, 0xba, 0xa9, 0xe1, 0xb0, 0x1b, 0x43, 0xcb, 0x1d,
	0x04, 0x4f, 0x6f, 0x80, 0x6e, 0x35, 0xbd, 0x6d, 0x75, 0xb9, 0x11, 0xf8, 0xac, 0x79, 0x31, 0xba,
	0xa4, 0x57, 0x6c, 0xce, 0x5d, 0xe8, 0xd7, 0x16, 0x48, 0x1b, 0xf2, 0x95, 0xbd, 0xff, 0x48, 0x17,
	0x5a, 0x67, 0xe7, 0xa3, 0x97, 0x74, 0xcf, 0x72, 0x7e, 0x58, 0xd0, 0x37, 0xb7, 0x71, 0xa1, 0x98,
	0x42, 0xf2, 0x08, 0xba, 0xe5, 0x6d, 0xf1, 0x3c, 0x8c, 0xe5, 0xf6, 0x82, 0xc3, 0xba, 0xa3, 0x91,
	0x7e, 0x1e, 0xcf, 0x68, 0x07, 0xcb, 0x8a, 0xbc, 0x87, 0x5d, 0x9d, 0x67, 0x52, 0x0d, 0xd3, 0x74,
	0x7b, 0x81, 0xfb, 0xb7, 0x61, 0xe8, 0x60, 0x5e, 0xeb, 0x9d, 0x9f, 0x0d, 0xb8, 0x75, 0x56, 0x1c,
	0xf4, 0x3f, 0x65, 0x2c, 0xbf, 0xe5, 0x05, 0xb2, 0x98, 0x27, 0x91, 0xdd, 0x1c, 0x5a, 0x6e, 0x87,
	0x9a, 0x36, 0xe0, 0xd0, 0x31, 0x3c, 0xe4, 0x5e, 0x9d, 0xbf, 0xfc, 0x35, 0x73, 0xfa, 0x6a, 0x28,
	0x1b, 0x75, 0x8e, 0xfb, 0xd7, 0x9b, 0xa9, 0x6e, 0xd3, 0x2d, 0xfd, 0xab, 0x17, 0x9f, 0x9e, 0x47,
	0x5c, 0x2d, 0xb2, 0xa9, 0x17, 0xa6, 0x4b, 0x7f, 0x95, 0x09, 0x94, 0xe2, 0xe4, 0x44, 0xf9, 0x51,
	0x2a, 0x16, 0x28, 0x43, 0xc9, 0xe6, 0xca, 0x9f, 0x26, 0xa8, 0xfc, 0x69, 0xb4, 0xf2, 0x0d, 0xb1,
	0x6f, 0xe0, 0xfe, 0xfa, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x59, 0x26, 0x92, 0x89,
	0x04, 0x00, 0x00,
}