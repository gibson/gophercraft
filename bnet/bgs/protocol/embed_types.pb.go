// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: embed_types.proto

package protocol

import (
	fmt "fmt"
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

type EmbedImage struct {
	Url                  *string  `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Width                *uint32  `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height               *uint32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmbedImage) Reset()         { *m = EmbedImage{} }
func (m *EmbedImage) String() string { return proto.CompactTextString(m) }
func (*EmbedImage) ProtoMessage()    {}
func (*EmbedImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{0}
}

func (m *EmbedImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmbedImage.Unmarshal(m, b)
}
func (m *EmbedImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmbedImage.Marshal(b, m, deterministic)
}
func (m *EmbedImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbedImage.Merge(m, src)
}
func (m *EmbedImage) XXX_Size() int {
	return xxx_messageInfo_EmbedImage.Size(m)
}
func (m *EmbedImage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbedImage.DiscardUnknown(m)
}

var xxx_messageInfo_EmbedImage proto.InternalMessageInfo

func (m *EmbedImage) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *EmbedImage) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *EmbedImage) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type Provider struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{1}
}

func (m *Provider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Provider.Unmarshal(m, b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return xxx_messageInfo_Provider.Size(m)
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type EmbedHTML struct {
	Content              *string  `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Width                *uint32  `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height               *uint32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmbedHTML) Reset()         { *m = EmbedHTML{} }
func (m *EmbedHTML) String() string { return proto.CompactTextString(m) }
func (*EmbedHTML) ProtoMessage()    {}
func (*EmbedHTML) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{2}
}

func (m *EmbedHTML) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmbedHTML.Unmarshal(m, b)
}
func (m *EmbedHTML) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmbedHTML.Marshal(b, m, deterministic)
}
func (m *EmbedHTML) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbedHTML.Merge(m, src)
}
func (m *EmbedHTML) XXX_Size() int {
	return xxx_messageInfo_EmbedHTML.Size(m)
}
func (m *EmbedHTML) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbedHTML.DiscardUnknown(m)
}

var xxx_messageInfo_EmbedHTML proto.InternalMessageInfo

func (m *EmbedHTML) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *EmbedHTML) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *EmbedHTML) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type EmbedInfo struct {
	Title                *string     `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Type                 *string     `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	OriginalUrl          *string     `protobuf:"bytes,3,opt,name=original_url,json=originalUrl" json:"original_url,omitempty"`
	Thumbnail            *EmbedImage `protobuf:"bytes,4,opt,name=thumbnail" json:"thumbnail,omitempty"`
	Provider             *Provider   `protobuf:"bytes,5,opt,name=provider" json:"provider,omitempty"`
	Description          *string     `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Html                 *EmbedHTML  `protobuf:"bytes,8,opt,name=html" json:"html,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EmbedInfo) Reset()         { *m = EmbedInfo{} }
func (m *EmbedInfo) String() string { return proto.CompactTextString(m) }
func (*EmbedInfo) ProtoMessage()    {}
func (*EmbedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a65fcf161f3afa9f, []int{3}
}

func (m *EmbedInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmbedInfo.Unmarshal(m, b)
}
func (m *EmbedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmbedInfo.Marshal(b, m, deterministic)
}
func (m *EmbedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbedInfo.Merge(m, src)
}
func (m *EmbedInfo) XXX_Size() int {
	return xxx_messageInfo_EmbedInfo.Size(m)
}
func (m *EmbedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EmbedInfo proto.InternalMessageInfo

func (m *EmbedInfo) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *EmbedInfo) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *EmbedInfo) GetOriginalUrl() string {
	if m != nil && m.OriginalUrl != nil {
		return *m.OriginalUrl
	}
	return ""
}

func (m *EmbedInfo) GetThumbnail() *EmbedImage {
	if m != nil {
		return m.Thumbnail
	}
	return nil
}

func (m *EmbedInfo) GetProvider() *Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *EmbedInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *EmbedInfo) GetHtml() *EmbedHTML {
	if m != nil {
		return m.Html
	}
	return nil
}

func init() {
	proto.RegisterType((*EmbedImage)(nil), "bgs.protocol.EmbedImage")
	proto.RegisterType((*Provider)(nil), "bgs.protocol.Provider")
	proto.RegisterType((*EmbedHTML)(nil), "bgs.protocol.EmbedHTML")
	proto.RegisterType((*EmbedInfo)(nil), "bgs.protocol.EmbedInfo")
}

func init() { proto.RegisterFile("embed_types.proto", fileDescriptor_a65fcf161f3afa9f) }

var fileDescriptor_a65fcf161f3afa9f = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0x5f, 0x4b, 0xfb, 0x30,
	0x14, 0xa5, 0xfb, 0xf7, 0x5b, 0xef, 0xf6, 0x03, 0x0d, 0x32, 0xf3, 0x24, 0x75, 0x4f, 0x03, 0xa1,
	0x1d, 0x43, 0xfc, 0x00, 0x82, 0xe0, 0x60, 0x82, 0x54, 0x7d, 0xf1, 0x65, 0xf4, 0x4f, 0xd6, 0x06,
	0xd2, 0x24, 0xa4, 0xb7, 0x8a, 0x9f, 0xc3, 0x2f, 0x2c, 0xc9, 0x5a, 0x37, 0xc1, 0x17, 0xdf, 0xce,
	0xb9, 0x39, 0xf7, 0xdc, 0x93, 0x03, 0xa7, 0xac, 0x4a, 0x59, 0xbe, 0xc5, 0x0f, 0xcd, 0xea, 0x50,
	0x1b, 0x85, 0x8a, 0x4c, 0xd3, 0xa2, 0x85, 0x99, 0x12, 0xf3, 0x0d, 0xc0, 0x9d, 0x95, 0xac, 0xab,
	0xa4, 0x60, 0xe4, 0x04, 0xfa, 0x8d, 0x11, 0xd4, 0x0b, 0xbc, 0x85, 0x1f, 0x5b, 0x48, 0xce, 0x60,
	0xf8, 0xce, 0x73, 0x2c, 0x69, 0x2f, 0xf0, 0x16, 0xff, 0xe3, 0x3d, 0x21, 0x33, 0x18, 0x95, 0x8c,
	0x17, 0x25, 0xd2, 0xbe, 0x1b, 0xb7, 0x6c, 0x7e, 0x01, 0xe3, 0x47, 0xa3, 0xde, 0x78, 0xce, 0x0c,
	0x21, 0x30, 0x90, 0x49, 0xc5, 0x5a, 0x33, 0x87, 0xe7, 0x4f, 0xe0, 0xbb, 0x6b, 0xf7, 0xcf, 0x0f,
	0x1b, 0x42, 0xe1, 0x5f, 0xa6, 0x24, 0x32, 0x89, 0xad, 0xa6, 0xa3, 0x7f, 0x3c, 0xfa, 0xd9, 0x6b,
	0x5d, 0xd7, 0x72, 0xa7, 0xec, 0x2e, 0x72, 0x14, 0xdd, 0xdd, 0x3d, 0xb1, 0x61, 0x6c, 0x07, 0xce,
	0xd0, 0x8f, 0x1d, 0x26, 0x97, 0x30, 0x55, 0x86, 0x17, 0x5c, 0x26, 0x62, 0x6b, 0x7f, 0xdd, 0x77,
	0x6f, 0x93, 0x6e, 0xf6, 0x62, 0x04, 0xb9, 0x01, 0x1f, 0xcb, 0xa6, 0x4a, 0x65, 0xc2, 0x05, 0x1d,
	0x04, 0xde, 0x62, 0xb2, 0xa2, 0xe1, 0x71, 0x7f, 0xe1, 0xa1, 0xbc, 0xf8, 0x20, 0x25, 0x2b, 0x18,
	0xeb, 0xb6, 0x07, 0x3a, 0x74, 0x6b, 0xb3, 0x9f, 0x6b, 0x5d, 0x4b, 0xf1, 0xb7, 0x8e, 0x04, 0x30,
	0xc9, 0x59, 0x9d, 0x19, 0xae, 0x91, 0x2b, 0x49, 0x47, 0xfb, 0x34, 0x47, 0x23, 0x72, 0x05, 0x83,
	0x12, 0x2b, 0x41, 0xc7, 0xce, 0xf1, 0xfc, 0x97, 0x20, 0xb6, 0xd7, 0xd8, 0x89, 0x6e, 0xaf, 0x5f,
	0x57, 0x05, 0xc7, 0xb2, 0x49, 0xc3, 0x4c, 0x55, 0x51, 0xdd, 0x68, 0x66, 0xf4, 0x72, 0x89, 0x51,
	0xa1, 0x74, 0xc9, 0x4c, 0x66, 0x92, 0x1d, 0x46, 0xa9, 0x64, 0x18, 0xa5, 0x45, 0x1d, 0x75, 0x2e,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xe2, 0xb9, 0x6c, 0x30, 0x02, 0x00, 0x00,
}
