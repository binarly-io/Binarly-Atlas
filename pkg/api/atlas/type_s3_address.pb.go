// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_s3_address.proto

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

// S3Address represents S3 and MinIO address
type S3Address struct {
	// Bucket name
	Bucket string `protobuf:"bytes,100,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Object name
	Object               string   `protobuf:"bytes,200,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S3Address) Reset()         { *m = S3Address{} }
func (m *S3Address) String() string { return proto.CompactTextString(m) }
func (*S3Address) ProtoMessage()    {}
func (*S3Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_3606c6ec6b03573f, []int{0}
}

func (m *S3Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3Address.Unmarshal(m, b)
}
func (m *S3Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3Address.Marshal(b, m, deterministic)
}
func (m *S3Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3Address.Merge(m, src)
}
func (m *S3Address) XXX_Size() int {
	return xxx_messageInfo_S3Address.Size(m)
}
func (m *S3Address) XXX_DiscardUnknown() {
	xxx_messageInfo_S3Address.DiscardUnknown(m)
}

var xxx_messageInfo_S3Address proto.InternalMessageInfo

func (m *S3Address) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *S3Address) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func init() {
	proto.RegisterType((*S3Address)(nil), "atlas.S3Address")
}

func init() { proto.RegisterFile("type_s3_address.proto", fileDescriptor_3606c6ec6b03573f) }

var fileDescriptor_3606c6ec6b03573f = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0xa9, 0x2c, 0x48,
	0x8d, 0x2f, 0x36, 0x8e, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x56, 0xb2, 0xe1, 0xe2, 0x0c, 0x36, 0x76, 0x84,
	0xc8, 0x08, 0x89, 0x71, 0xb1, 0x25, 0x95, 0x26, 0x67, 0xa7, 0x96, 0x48, 0xa4, 0x28, 0x30, 0x6a,
	0x70, 0x06, 0x41, 0x79, 0x42, 0xe2, 0x5c, 0x6c, 0xf9, 0x49, 0x59, 0xa9, 0xc9, 0x25, 0x12, 0x27,
	0x18, 0x21, 0x12, 0x10, 0x6e, 0x12, 0x1b, 0xd8, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe3, 0xfa, 0x0f, 0xfb, 0x64, 0x00, 0x00, 0x00,
}
