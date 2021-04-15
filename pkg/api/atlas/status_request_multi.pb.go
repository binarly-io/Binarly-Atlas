// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status_request_multi.proto

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

type StatusRequestMulti struct {
	Domain               *Domain           `protobuf:"bytes,100,opt,name=domain,proto3" json:"domain,omitempty"`
	Mode                 StatusRequestMode `protobuf:"varint,200,opt,name=mode,proto3,enum=atlas.StatusRequestMode" json:"mode,omitempty"`
	Entities             []*StatusRequest  `protobuf:"bytes,300,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatusRequestMulti) Reset()      { *m = StatusRequestMulti{} }
func (*StatusRequestMulti) ProtoMessage() {}
func (*StatusRequestMulti) Descriptor() ([]byte, []int) {
	return fileDescriptor_53d95ec84ec0cd32, []int{0}
}

func (m *StatusRequestMulti) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequestMulti.Unmarshal(m, b)
}
func (m *StatusRequestMulti) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequestMulti.Marshal(b, m, deterministic)
}
func (m *StatusRequestMulti) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequestMulti.Merge(m, src)
}
func (m *StatusRequestMulti) XXX_Size() int {
	return xxx_messageInfo_StatusRequestMulti.Size(m)
}
func (m *StatusRequestMulti) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequestMulti.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequestMulti proto.InternalMessageInfo

func (m *StatusRequestMulti) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *StatusRequestMulti) GetMode() StatusRequestMode {
	if m != nil {
		return m.Mode
	}
	return StatusRequestMode_RESERVED
}

func (m *StatusRequestMulti) GetEntities() []*StatusRequest {
	if m != nil {
		return m.Entities
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusRequestMulti)(nil), "atlas.StatusRequestMulti")
}

func init() { proto.RegisterFile("status_request_multi.proto", fileDescriptor_53d95ec84ec0cd32) }

var fileDescriptor_53d95ec84ec0cd32 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0x8e, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x89, 0xcf, 0x2d, 0xcd, 0x29, 0xc9,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0xe2, 0x49,
	0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x83, 0x08, 0x4a, 0x89, 0xa0, 0x6a, 0x80, 0x8a, 0x4a, 0xa2, 0x1b,
	0x93, 0x9f, 0x92, 0x0a, 0x91, 0x52, 0x9a, 0xcb, 0xc8, 0x25, 0x14, 0x0c, 0x96, 0x0d, 0x82, 0x48,
	0xfa, 0x82, 0xac, 0x10, 0x52, 0xe5, 0x62, 0x83, 0x98, 0x2b, 0x91, 0xa2, 0xc0, 0xa8, 0xc1, 0x6d,
	0xc4, 0xab, 0x07, 0xb6, 0x4d, 0xcf, 0x05, 0x2c, 0x18, 0x04, 0x95, 0x14, 0xd2, 0xe5, 0x62, 0x01,
	0x99, 0x25, 0x71, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x02, 0xaa, 0x0a, 0xd5, 0xc0, 0xfc,
	0x94, 0xd4, 0x20, 0xb0, 0x32, 0x21, 0x43, 0x2e, 0x8e, 0xd4, 0xbc, 0x92, 0xcc, 0x92, 0xcc, 0xd4,
	0x62, 0x89, 0x35, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x22, 0xd8, 0xb4, 0x04, 0xc1, 0x95, 0x25,
	0xb1, 0x81, 0x9d, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x4d, 0x52, 0xc4, 0x0a, 0x01,
	0x00, 0x00,
}