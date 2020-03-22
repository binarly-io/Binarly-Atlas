// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_command.proto

package mservice

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

type CommandType int32

const (
	// Due to first enum value has to be zero in proto3
	CommandType_ACTION_TYPE_RESERVED    CommandType = 0
	CommandType_ACTION_TYPE_UNSPECIFIED CommandType = 10
	CommandType_ACTION_PING             CommandType = 20
	CommandType_ACTION_CONFIG           CommandType = 30
	CommandType_ACTION_METRICS          CommandType = 40
	CommandType_ACTION_SCAN_SCHEDULE    CommandType = 50
	CommandType_ACTION_SCAN_REQUEST     CommandType = 60
)

var CommandType_name = map[int32]string{
	0:  "ACTION_TYPE_RESERVED",
	10: "ACTION_TYPE_UNSPECIFIED",
	20: "ACTION_PING",
	30: "ACTION_CONFIG",
	40: "ACTION_METRICS",
	50: "ACTION_SCAN_SCHEDULE",
	60: "ACTION_SCAN_REQUEST",
}

var CommandType_value = map[string]int32{
	"ACTION_TYPE_RESERVED":    0,
	"ACTION_TYPE_UNSPECIFIED": 10,
	"ACTION_PING":             20,
	"ACTION_CONFIG":           30,
	"ACTION_METRICS":          40,
	"ACTION_SCAN_SCHEDULE":    50,
	"ACTION_SCAN_REQUEST":     60,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}

func (CommandType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e84d1630985fc0a4, []int{0}
}

type Command struct {
	Type                 CommandType `protobuf:"varint,10,opt,name=type,proto3,enum=mservice.CommandType" json:"type,omitempty"`
	Uuid                 *UUID       `protobuf:"bytes,20,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_e84d1630985fc0a4, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetType() CommandType {
	if m != nil {
		return m.Type
	}
	return CommandType_ACTION_TYPE_RESERVED
}

func (m *Command) GetUuid() *UUID {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func init() {
	proto.RegisterEnum("mservice.CommandType", CommandType_name, CommandType_value)
	proto.RegisterType((*Command)(nil), "mservice.Command")
}

func init() {
	proto.RegisterFile("type_command.proto", fileDescriptor_e84d1630985fc0a4)
}

var fileDescriptor_e84d1630985fc0a4 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc0, 0x71, 0x03, 0x45, 0x65, 0x82, 0x71, 0x3b, 0x46, 0x1a, 0x14, 0xa4, 0xf4, 0x14, 0x3d,
	0xe4, 0x10, 0xaf, 0x5e, 0xca, 0x66, 0x5a, 0x17, 0x34, 0x8d, 0xfb, 0x21, 0x7a, 0x0a, 0xda, 0xe6,
	0xd0, 0x43, 0x9a, 0x50, 0x5b, 0xa1, 0x4f, 0xe5, 0x2b, 0x4a, 0x93, 0x68, 0xf7, 0xb2, 0x87, 0xf9,
	0xfd, 0x87, 0x85, 0x01, 0xdc, 0xec, 0xea, 0x22, 0x9f, 0x57, 0x65, 0xf9, 0xb1, 0x5a, 0x44, 0xf5,
	0xba, 0xda, 0x54, 0x78, 0x5a, 0x7e, 0x15, 0xeb, 0xef, 0xe5, 0xbc, 0xb8, 0xea, 0xff, 0x6b, 0xb5,
	0x6a, 0x71, 0xf4, 0x06, 0x27, 0xbc, 0xad, 0xf1, 0x16, 0x7a, 0x7b, 0x0f, 0x60, 0xe8, 0x84, 0x5e,
	0x7c, 0x19, 0xfd, 0xad, 0x45, 0x5d, 0xa0, 0x77, 0x75, 0x21, 0x9b, 0x04, 0x47, 0xd0, 0xdb, 0x6e,
	0x97, 0x8b, 0xc0, 0x1f, 0x3a, 0xa1, 0x1b, 0x7b, 0x87, 0xd4, 0x18, 0x91, 0xc8, 0xc6, 0xee, 0x7e,
	0x1c, 0x70, 0xad, 0x4d, 0x0c, 0xc0, 0x1f, 0x73, 0x2d, 0x66, 0x69, 0xae, 0xdf, 0x33, 0xca, 0x25,
	0x29, 0x92, 0xaf, 0x94, 0xb0, 0x23, 0xbc, 0x86, 0x81, 0x2d, 0x26, 0x55, 0x19, 0x71, 0x31, 0x11,
	0x94, 0x30, 0xc0, 0x73, 0x70, 0x3b, 0xcc, 0x44, 0x3a, 0x65, 0x3e, 0xf6, 0xe1, 0xac, 0x1b, 0xf0,
	0x59, 0x3a, 0x11, 0x53, 0x76, 0x83, 0x08, 0x5e, 0x37, 0x7a, 0x26, 0x2d, 0x05, 0x57, 0x2c, 0xb4,
	0xbe, 0x53, 0x7c, 0xbc, 0x7f, 0x1e, 0x29, 0x31, 0x4f, 0xc4, 0x62, 0x1c, 0xc0, 0x85, 0x2d, 0x92,
	0x5e, 0x0c, 0x29, 0xcd, 0x1e, 0x3e, 0x8f, 0x9b, 0x93, 0xdc, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xd2, 0xcb, 0x9a, 0x28, 0x45, 0x01, 0x00, 0x00,
}
