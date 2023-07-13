// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: club_member_id.proto

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

type MemberId struct {
	AccountId            *v1.AccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UniqueId             *uint64       `protobuf:"varint,2,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MemberId) Reset()         { *m = MemberId{} }
func (m *MemberId) String() string { return proto.CompactTextString(m) }
func (*MemberId) ProtoMessage()    {}
func (*MemberId) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdef93e55a672430, []int{0}
}

func (m *MemberId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberId.Unmarshal(m, b)
}
func (m *MemberId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberId.Marshal(b, m, deterministic)
}
func (m *MemberId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberId.Merge(m, src)
}
func (m *MemberId) XXX_Size() int {
	return xxx_messageInfo_MemberId.Size(m)
}
func (m *MemberId) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberId.DiscardUnknown(m)
}

var xxx_messageInfo_MemberId proto.InternalMessageInfo

func (m *MemberId) GetAccountId() *v1.AccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *MemberId) GetUniqueId() uint64 {
	if m != nil && m.UniqueId != nil {
		return *m.UniqueId
	}
	return 0
}

func init() {
	proto.RegisterType((*MemberId)(nil), "bgs.protocol.club.v1.MemberId")
}

func init() { proto.RegisterFile("club_member_id.proto", fileDescriptor_cdef93e55a672430) }

var fileDescriptor_cdef93e55a672430 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x15, 0xc4, 0xd0, 0x9a, 0x01, 0x29, 0x74, 0x40, 0x65, 0xa9, 0xba, 0xd0, 0xc9, 0x6e,
	0x18, 0x11, 0x4b, 0xd9, 0x32, 0xb0, 0x74, 0x64, 0xb1, 0xfc, 0xe7, 0x70, 0x8c, 0x12, 0x9f, 0xb1,
	0xcf, 0x11, 0x7c, 0x7b, 0x14, 0x42, 0x24, 0xd8, 0xde, 0xfd, 0xee, 0xde, 0xbd, 0x3b, 0xb6, 0x31,
	0x7d, 0xd1, 0x72, 0x80, 0x41, 0x43, 0x92, 0xde, 0xf2, 0x98, 0x90, 0xb0, 0xde, 0x68, 0x97, 0x67,
	0x69, 0xb0, 0xe7, 0xd3, 0x08, 0x1f, 0x9b, 0xed, 0xbd, 0xeb, 0x51, 0xab, 0x5e, 0xc2, 0x27, 0x41,
	0xc8, 0x1e, 0x43, 0x16, 0x03, 0xe4, 0xac, 0x1c, 0x48, 0x8c, 0x34, 0xd5, 0xb3, 0x67, 0x7b, 0xa3,
	0x8c, 0xc1, 0x12, 0x48, 0xd2, 0x57, 0x84, 0x05, 0x5e, 0xa7, 0x68, 0xfe, 0x82, 0xfd, 0x3b, 0x5b,
	0xbd, 0xfc, 0xe4, 0xb6, 0xb6, 0x3e, 0x31, 0xb6, 0x78, 0xbc, 0xbd, 0xad, 0x76, 0xd5, 0xe1, 0xea,
	0x61, 0xcf, 0xff, 0x5d, 0xf1, 0xdb, 0xe7, 0x63, 0xc3, 0x4f, 0xb3, 0x6c, 0xed, 0x79, 0xad, 0x16,
	0x59, 0xdf, 0xb1, 0x75, 0x09, 0xfe, 0xa3, 0xc0, 0xb4, 0xe1, 0x62, 0x57, 0x1d, 0x2e, 0xcf, 0xab,
	0x19, 0xb4, 0xf6, 0xf9, 0xe9, 0xf5, 0xd1, 0x79, 0xea, 0x8a, 0xe6, 0x06, 0x07, 0x91, 0x4b, 0x84,
	0x14, 0x8f, 0x47, 0x12, 0x0e, 0x63, 0x07, 0xc9, 0x24, 0xf5, 0x46, 0x42, 0x07, 0x20, 0xa1, 0x5d,
	0x16, 0x4b, 0xa4, 0x98, 0x1e, 0x17, 0x63, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x3c, 0x0b,
	0x3d, 0x25, 0x01, 0x00, 0x00,
}
