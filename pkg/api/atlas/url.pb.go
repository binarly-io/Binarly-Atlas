// Code generated by protoc-gen-go. DO NOT EDIT.
// source: url.proto

package atlas

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

// URL represents abstract url.
type URL struct {
	// URL
	Url                  string   `protobuf:"bytes,100,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URL) Reset()      { *m = URL{} }
func (*URL) ProtoMessage() {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2def92d9e1aa990, []int{0}
}

func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (m *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(m, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*URL)(nil), "atlas.URL")
}

func init() { proto.RegisterFile("url.proto", fileDescriptor_a2def92d9e1aa990) }

var fileDescriptor_a2def92d9e1aa990 = []byte{
	// 70 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2d, 0xca, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x56, 0x12, 0xe7, 0x62,
	0x0e, 0x0d, 0xf2, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x91, 0x48, 0x51, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0x31, 0x93, 0xd8, 0xc0, 0xca, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xd8,
	0x94, 0x91, 0x33, 0x00, 0x00, 0x00,
}
