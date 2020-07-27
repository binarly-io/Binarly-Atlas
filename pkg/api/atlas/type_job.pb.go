// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_job.proto

package atlas

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Job struct {
	// Types that are valid to be assigned to TypeOptional:
	//	*Job_Type
	TypeOptional isJob_TypeOptional `protobuf_oneof:"type_optional"`
	// Types that are valid to be assigned to NameOptional:
	//	*Job_Name
	NameOptional isJob_NameOptional `protobuf_oneof:"name_optional"`
	// Types that are valid to be assigned to VersionOptional:
	//	*Job_Version
	VersionOptional isJob_VersionOptional `protobuf_oneof:"version_optional"`
	// Types that are valid to be assigned to UuidOptional:
	//	*Job_Uuid
	UuidOptional isJob_UuidOptional `protobuf_oneof:"uuid_optional"`
	// Types that are valid to be assigned to UuidReferenceOptional:
	//	*Job_UuidReference
	UuidReferenceOptional isJob_UuidReferenceOptional `protobuf_oneof:"uuid_reference_optional"`
	// Types that are valid to be assigned to TimestampOptional:
	//	*Job_Ts
	TimestampOptional    isJob_TimestampOptional `protobuf_oneof:"timestamp_optional"`
	Addresses            []*S3Address            `protobuf:"bytes,700,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9367f4548affdc3, []int{0}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

type isJob_TypeOptional interface {
	isJob_TypeOptional()
}

type Job_Type struct {
	Type int32 `protobuf:"varint,100,opt,name=type,proto3,oneof"`
}

func (*Job_Type) isJob_TypeOptional() {}

func (m *Job) GetTypeOptional() isJob_TypeOptional {
	if m != nil {
		return m.TypeOptional
	}
	return nil
}

func (m *Job) GetType() int32 {
	if x, ok := m.GetTypeOptional().(*Job_Type); ok {
		return x.Type
	}
	return 0
}

type isJob_NameOptional interface {
	isJob_NameOptional()
}

type Job_Name struct {
	Name string `protobuf:"bytes,200,opt,name=name,proto3,oneof"`
}

func (*Job_Name) isJob_NameOptional() {}

func (m *Job) GetNameOptional() isJob_NameOptional {
	if m != nil {
		return m.NameOptional
	}
	return nil
}

func (m *Job) GetName() string {
	if x, ok := m.GetNameOptional().(*Job_Name); ok {
		return x.Name
	}
	return ""
}

type isJob_VersionOptional interface {
	isJob_VersionOptional()
}

type Job_Version struct {
	Version int32 `protobuf:"varint,300,opt,name=version,proto3,oneof"`
}

func (*Job_Version) isJob_VersionOptional() {}

func (m *Job) GetVersionOptional() isJob_VersionOptional {
	if m != nil {
		return m.VersionOptional
	}
	return nil
}

func (m *Job) GetVersion() int32 {
	if x, ok := m.GetVersionOptional().(*Job_Version); ok {
		return x.Version
	}
	return 0
}

type isJob_UuidOptional interface {
	isJob_UuidOptional()
}

type Job_Uuid struct {
	Uuid *UUID `protobuf:"bytes,400,opt,name=uuid,proto3,oneof"`
}

func (*Job_Uuid) isJob_UuidOptional() {}

func (m *Job) GetUuidOptional() isJob_UuidOptional {
	if m != nil {
		return m.UuidOptional
	}
	return nil
}

func (m *Job) GetUuid() *UUID {
	if x, ok := m.GetUuidOptional().(*Job_Uuid); ok {
		return x.Uuid
	}
	return nil
}

type isJob_UuidReferenceOptional interface {
	isJob_UuidReferenceOptional()
}

type Job_UuidReference struct {
	UuidReference *UUID `protobuf:"bytes,500,opt,name=uuid_reference,json=uuidReference,proto3,oneof"`
}

func (*Job_UuidReference) isJob_UuidReferenceOptional() {}

func (m *Job) GetUuidReferenceOptional() isJob_UuidReferenceOptional {
	if m != nil {
		return m.UuidReferenceOptional
	}
	return nil
}

func (m *Job) GetUuidReference() *UUID {
	if x, ok := m.GetUuidReferenceOptional().(*Job_UuidReference); ok {
		return x.UuidReference
	}
	return nil
}

type isJob_TimestampOptional interface {
	isJob_TimestampOptional()
}

type Job_Ts struct {
	Ts *timestamp.Timestamp `protobuf:"bytes,600,opt,name=ts,proto3,oneof"`
}

func (*Job_Ts) isJob_TimestampOptional() {}

func (m *Job) GetTimestampOptional() isJob_TimestampOptional {
	if m != nil {
		return m.TimestampOptional
	}
	return nil
}

func (m *Job) GetTs() *timestamp.Timestamp {
	if x, ok := m.GetTimestampOptional().(*Job_Ts); ok {
		return x.Ts
	}
	return nil
}

func (m *Job) GetAddresses() []*S3Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Job) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Job_Type)(nil),
		(*Job_Name)(nil),
		(*Job_Version)(nil),
		(*Job_Uuid)(nil),
		(*Job_UuidReference)(nil),
		(*Job_Ts)(nil),
	}
}

func init() {
	proto.RegisterType((*Job)(nil), "atlas.Job")
}

func init() {
	proto.RegisterFile("type_job.proto", fileDescriptor_a9367f4548affdc3)
}

var fileDescriptor_a9367f4548affdc3 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x4e, 0xf3, 0x30,
	0x10, 0x86, 0xbf, 0xfc, 0x7d, 0xa8, 0x53, 0xf5, 0x47, 0x56, 0x2b, 0x4c, 0x59, 0x10, 0x75, 0x95,
	0x0d, 0xae, 0xd4, 0x8a, 0x03, 0x80, 0x58, 0x04, 0x96, 0x86, 0xae, 0x2b, 0x87, 0xb8, 0x55, 0x50,
	0x1b, 0x47, 0xb1, 0x8b, 0xc4, 0x96, 0x13, 0x70, 0x10, 0x8e, 0xc0, 0x01, 0x38, 0x02, 0x07, 0xe1,
	0x00, 0xc8, 0x93, 0xa4, 0x51, 0x25, 0x56, 0x91, 0x9f, 0x79, 0xac, 0xc9, 0xfb, 0x1a, 0xfa, 0xe6,
	0xb5, 0x90, 0xab, 0x67, 0x95, 0xb0, 0xa2, 0x54, 0x46, 0x91, 0x40, 0x98, 0xad, 0xd0, 0x93, 0x8b,
	0x8d, 0x52, 0x9b, 0xad, 0x9c, 0x21, 0x4c, 0xf6, 0xeb, 0x99, 0xc9, 0x76, 0x52, 0x1b, 0xb1, 0x2b,
	0x2a, 0x6f, 0x32, 0xc0, 0x7b, 0xfb, 0x7d, 0x96, 0xd6, 0x60, 0x8c, 0x40, 0x2f, 0x56, 0x22, 0x4d,
	0x4b, 0xa9, 0x75, 0x85, 0xa7, 0x6f, 0x1e, 0x78, 0xf7, 0x2a, 0x21, 0x23, 0xf0, 0xad, 0x40, 0xd3,
	0xd0, 0x89, 0x82, 0xf8, 0x1f, 0xc7, 0x13, 0x19, 0x83, 0x9f, 0x8b, 0x9d, 0xa4, 0x5f, 0x4e, 0xe8,
	0x44, 0x9d, 0xd8, 0xe1, 0x78, 0x24, 0xe7, 0x70, 0xf2, 0x22, 0x4b, 0x9d, 0xa9, 0x9c, 0x7e, 0xb8,
	0x78, 0xc1, 0xe5, 0x0d, 0x21, 0x53, 0xf0, 0xed, 0x5a, 0xfa, 0xee, 0x85, 0x4e, 0xd4, 0x9d, 0x77,
	0x19, 0xfe, 0x31, 0x5b, 0x2e, 0xef, 0x6e, 0x63, 0x8f, 0xe3, 0x8c, 0x5c, 0x41, 0xdf, 0x7e, 0x57,
	0xa5, 0x5c, 0xcb, 0x52, 0xe6, 0x4f, 0x92, 0xfe, 0xfc, 0x61, 0xfb, 0xbc, 0x67, 0x2d, 0xde, 0x48,
	0xe4, 0x12, 0x5c, 0xa3, 0xe9, 0xb7, 0x8f, 0xea, 0x84, 0x55, 0x1d, 0xb0, 0xa6, 0x03, 0xf6, 0xd8,
	0x74, 0x10, 0x07, 0xdc, 0x35, 0x9a, 0xcc, 0xa0, 0x53, 0x87, 0x95, 0x9a, 0x7e, 0x06, 0xa1, 0x17,
	0x75, 0xe7, 0xc3, 0x7a, 0xc1, 0xc3, 0xe2, 0xba, 0x1a, 0xf1, 0xd6, 0xb9, 0x19, 0x40, 0x0f, 0x5b,
	0x52, 0x85, 0xc9, 0x54, 0x2e, 0xb6, 0x16, 0xd8, 0xc0, 0x2d, 0x20, 0x30, 0xac, 0x73, 0x1e, 0x49,
	0x18, 0xe6, 0x00, 0xce, 0xe0, 0xf4, 0x38, 0x5d, 0x3b, 0x1a, 0x01, 0x39, 0xbc, 0xd4, 0x81, 0x26,
	0xff, 0x31, 0xc2, 0xe2, 0x37, 0x00, 0x00, 0xff, 0xff, 0xce, 0xf6, 0xf3, 0xc1, 0xed, 0x01, 0x00,
	0x00,
}