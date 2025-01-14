// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: api/client/v1/channel_id.proto

package v1

import (
	fmt "fmt"
	protocol "github.com/gibson/gophercraft/bnet/bgs/protocol"
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

type ChannelId struct {
	Type                 *uint32             `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Host                 *protocol.ProcessId `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Id                   *uint32             `protobuf:"fixed32,3,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ChannelId) Reset()         { *m = ChannelId{} }
func (m *ChannelId) String() string { return proto.CompactTextString(m) }
func (*ChannelId) ProtoMessage()    {}
func (*ChannelId) Descriptor() ([]byte, []int) {
	return fileDescriptor_25048f65e7320eca, []int{0}
}

func (m *ChannelId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelId.Unmarshal(m, b)
}
func (m *ChannelId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelId.Marshal(b, m, deterministic)
}
func (m *ChannelId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelId.Merge(m, src)
}
func (m *ChannelId) XXX_Size() int {
	return xxx_messageInfo_ChannelId.Size(m)
}
func (m *ChannelId) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelId.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelId proto.InternalMessageInfo

func (m *ChannelId) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ChannelId) GetHost() *protocol.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ChannelId) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ChannelId)(nil), "bgs.protocol.channel.v1.ChannelId")
}

func init() { proto.RegisterFile("api/client/v1/channel_id.proto", fileDescriptor_25048f65e7320eca) }

var fileDescriptor_25048f65e7320eca = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcf, 0xbb, 0x4b, 0xc5, 0x30,
	0x14, 0xc7, 0x71, 0x52, 0x2f, 0x88, 0x11, 0x15, 0xb2, 0xdc, 0xe2, 0x20, 0xc5, 0xa9, 0x20, 0xe4,
	0xdc, 0xba, 0x8b, 0xa0, 0x53, 0x37, 0xe9, 0x28, 0x42, 0x69, 0x1e, 0x26, 0x81, 0x9a, 0x84, 0xe4,
	0xb4, 0xe0, 0x7f, 0x2f, 0xf6, 0x31, 0xdc, 0xed, 0x07, 0x9f, 0xc3, 0x17, 0x0e, 0x7d, 0x18, 0xa2,
	0x03, 0x39, 0x3a, 0xed, 0x11, 0xe6, 0x06, 0xa4, 0x1d, 0xbc, 0xd7, 0x63, 0xef, 0x14, 0x8f, 0x29,
	0x60, 0x60, 0x47, 0x61, 0xf2, 0x3a, 0x65, 0x18, 0xf9, 0xc6, 0x7c, 0x6e, 0xee, 0xef, 0x52, 0x94,
	0x3d, 0xfe, 0x46, 0xbd, 0xf1, 0xe3, 0x17, 0xbd, 0x7a, 0x5f, 0xb9, 0x55, 0x8c, 0xd1, 0xc3, 0xbf,
	0x95, 0xa4, 0x22, 0xf5, 0x4d, 0xb7, 0x6c, 0xf6, 0x44, 0x0f, 0x36, 0x64, 0x2c, 0x8b, 0x8a, 0xd4,
	0xd7, 0xcf, 0x47, 0x7e, 0x56, 0xfe, 0x48, 0x41, 0xea, 0x9c, 0x5b, 0xd5, 0x2d, 0x47, 0xec, 0x96,
	0x16, 0x4e, 0x95, 0x17, 0x15, 0xa9, 0x2f, 0xbb, 0xc2, 0xa9, 0xb7, 0xd7, 0xcf, 0x17, 0xe3, 0xd0,
	0x4e, 0x82, 0xcb, 0xf0, 0x03, 0x79, 0x8a, 0x3a, 0xc5, 0xd3, 0x09, 0xc1, 0x84, 0x68, 0x75, 0x92,
	0x69, 0xf8, 0x46, 0x10, 0x5e, 0x23, 0x08, 0x93, 0x61, 0xaf, 0xee, 0xef, 0xc0, 0xdc, 0xfc, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x59, 0x64, 0xbd, 0xef, 0xe9, 0x00, 0x00, 0x00,
}
