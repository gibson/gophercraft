// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: challenge_service.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/Gophercraft/core/bnet/bgs/protocol"
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

type ChallengeExternalRequest struct {
	RequestToken         *string  `protobuf:"bytes,1,opt,name=request_token,json=requestToken" json:"request_token,omitempty"`
	PayloadType          *string  `protobuf:"bytes,2,opt,name=payload_type,json=payloadType" json:"payload_type,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChallengeExternalRequest) Reset()         { *m = ChallengeExternalRequest{} }
func (m *ChallengeExternalRequest) String() string { return proto.CompactTextString(m) }
func (*ChallengeExternalRequest) ProtoMessage()    {}
func (*ChallengeExternalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fddc6fc79451af25, []int{0}
}

func (m *ChallengeExternalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeExternalRequest.Unmarshal(m, b)
}
func (m *ChallengeExternalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeExternalRequest.Marshal(b, m, deterministic)
}
func (m *ChallengeExternalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeExternalRequest.Merge(m, src)
}
func (m *ChallengeExternalRequest) XXX_Size() int {
	return xxx_messageInfo_ChallengeExternalRequest.Size(m)
}
func (m *ChallengeExternalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeExternalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeExternalRequest proto.InternalMessageInfo

func (m *ChallengeExternalRequest) GetRequestToken() string {
	if m != nil && m.RequestToken != nil {
		return *m.RequestToken
	}
	return ""
}

func (m *ChallengeExternalRequest) GetPayloadType() string {
	if m != nil && m.PayloadType != nil {
		return *m.PayloadType
	}
	return ""
}

func (m *ChallengeExternalRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ChallengeExternalResult struct {
	RequestToken         *string  `protobuf:"bytes,1,opt,name=request_token,json=requestToken" json:"request_token,omitempty"`
	Passed               *bool    `protobuf:"varint,2,opt,name=passed,def=1" json:"passed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChallengeExternalResult) Reset()         { *m = ChallengeExternalResult{} }
func (m *ChallengeExternalResult) String() string { return proto.CompactTextString(m) }
func (*ChallengeExternalResult) ProtoMessage()    {}
func (*ChallengeExternalResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fddc6fc79451af25, []int{1}
}

func (m *ChallengeExternalResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeExternalResult.Unmarshal(m, b)
}
func (m *ChallengeExternalResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeExternalResult.Marshal(b, m, deterministic)
}
func (m *ChallengeExternalResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeExternalResult.Merge(m, src)
}
func (m *ChallengeExternalResult) XXX_Size() int {
	return xxx_messageInfo_ChallengeExternalResult.Size(m)
}
func (m *ChallengeExternalResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeExternalResult.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeExternalResult proto.InternalMessageInfo

const Default_ChallengeExternalResult_Passed bool = true

func (m *ChallengeExternalResult) GetRequestToken() string {
	if m != nil && m.RequestToken != nil {
		return *m.RequestToken
	}
	return ""
}

func (m *ChallengeExternalResult) GetPassed() bool {
	if m != nil && m.Passed != nil {
		return *m.Passed
	}
	return Default_ChallengeExternalResult_Passed
}

func init() {
	proto.RegisterType((*ChallengeExternalRequest)(nil), "bgs.protocol.challenge.v1.ChallengeExternalRequest")
	proto.RegisterType((*ChallengeExternalResult)(nil), "bgs.protocol.challenge.v1.ChallengeExternalResult")
}

func init() { proto.RegisterFile("challenge_service.proto", fileDescriptor_fddc6fc79451af25) }

var fileDescriptor_fddc6fc79451af25 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4f, 0xc2, 0x40,
	0x14, 0xc7, 0x2d, 0x10, 0xc4, 0x13, 0x63, 0x3c, 0x07, 0x0a, 0x71, 0x40, 0x1c, 0xc4, 0x90, 0xf4,
	0xf8, 0xb1, 0xb9, 0x98, 0xa8, 0x6c, 0x06, 0x4c, 0x61, 0x32, 0x26, 0x4d, 0x5b, 0x1e, 0xa5, 0xb1,
	0xf4, 0xce, 0xbb, 0x2b, 0xb1, 0x89, 0x83, 0x61, 0xf4, 0x5f, 0xf1, 0x2f, 0x64, 0x33, 0x85, 0x6b,
	0x13, 0x13, 0x88, 0xba, 0xdd, 0xfb, 0xf6, 0xfb, 0x3e, 0x9f, 0xf4, 0xa1, 0x8a, 0x3b, 0xb3, 0x83,
	0x00, 0x42, 0x0f, 0x2c, 0x01, 0x7c, 0xe1, 0xbb, 0x60, 0x30, 0x4e, 0x25, 0xc5, 0x55, 0xc7, 0x13,
	0x9b, 0xa7, 0x4b, 0x03, 0x23, 0x6b, 0x19, 0x8b, 0x4e, 0xed, 0x98, 0x33, 0xd7, 0x92, 0x31, 0x03,
	0x55, 0x68, 0xbc, 0x23, 0xfd, 0x2e, 0x2d, 0xf4, 0xdf, 0x24, 0xf0, 0xd0, 0x0e, 0x4c, 0x78, 0x8d,
	0x40, 0x48, 0x7c, 0x81, 0x8e, 0xf8, 0xe6, 0x69, 0x49, 0xfa, 0x02, 0xa1, 0xae, 0xd5, 0xb5, 0xe6,
	0x81, 0x59, 0x56, 0xe1, 0x38, 0xc9, 0xf0, 0x39, 0x2a, 0x33, 0x3b, 0x0e, 0xa8, 0x3d, 0x59, 0x73,
	0xf5, 0xdc, 0xba, 0x73, 0xa8, 0xb2, 0x71, 0xcc, 0x00, 0xeb, 0x68, 0x5f, 0x8d, 0x7a, 0xbe, 0xae,
	0x35, 0xcb, 0x66, 0x3a, 0x36, 0x9e, 0x51, 0x65, 0x8b, 0x5d, 0x44, 0xc1, 0x1f, 0xe5, 0x67, 0xa8,
	0xc8, 0x6c, 0x21, 0x60, 0xb2, 0xd6, 0x96, 0xae, 0x0b, 0x92, 0x47, 0x60, 0xaa, 0xac, 0xfb, 0x95,
	0x43, 0x27, 0x19, 0xfe, 0xc1, 0x17, 0x12, 0x42, 0xe0, 0x78, 0x8e, 0x4e, 0x87, 0x61, 0x2a, 0xcb,
	0x3e, 0xe3, 0x9e, 0xb1, 0xf3, 0x6a, 0xc6, 0xae, 0x0b, 0xd5, 0xaa, 0x3f, 0x97, 0x06, 0x43, 0xcb,
	0xec, 0x8f, 0x1e, 0x87, 0x83, 0x51, 0xbf, 0x51, 0x5c, 0xae, 0x5a, 0xb9, 0x52, 0x1e, 0x73, 0x54,
	0xdd, 0xa2, 0x53, 0x3f, 0xd9, 0xfd, 0x9f, 0x34, 0xd9, 0xf9, 0xdd, 0x59, 0xa8, 0xf5, 0x96, 0xab,
	0xd6, 0x15, 0xba, 0x74, 0x42, 0x90, 0xdb, 0xe0, 0x19, 0x79, 0x40, 0xa5, 0x3f, 0x8d, 0x3f, 0x93,
	0x25, 0xed, 0xf6, 0xfe, 0xe9, 0xc6, 0xf3, 0xe5, 0x2c, 0x72, 0x0c, 0x97, 0xce, 0x89, 0x88, 0x18,
	0x70, 0xd6, 0x6e, 0x4b, 0xe2, 0x51, 0x36, 0x03, 0xee, 0x72, 0x7b, 0x2a, 0x49, 0x42, 0x24, 0x8e,
	0x27, 0x48, 0x4a, 0x25, 0x19, 0x95, 0x2c, 0x3a, 0x1f, 0xda, 0xde, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0xba, 0xb9, 0xc5, 0x95, 0x02, 0x00, 0x00,
}
