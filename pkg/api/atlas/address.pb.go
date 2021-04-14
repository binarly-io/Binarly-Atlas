// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

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

// Address describes general address
type Address struct {
	// Types that are valid to be assigned to DomainOptional:
	//	*Address_AddressDomain
	DomainOptional isAddress_DomainOptional `protobuf_oneof:"domain_optional"`
	// Types that are valid to be assigned to AddressOptional:
	//	*Address_S3
	//	*Address_Kafka
	//	*Address_Digest
	//	*Address_Uuid
	//	*Address_UserId
	//	*Address_Filename
	//	*Address_Url
	//	*Address_Domain
	//	*Address_CustomString
	AddressOptional      isAddress_AddressOptional `protobuf_oneof:"address_optional"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Address) Reset()      { *m = Address{} }
func (*Address) ProtoMessage() {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

type isAddress_DomainOptional interface {
	isAddress_DomainOptional()
}

type Address_AddressDomain struct {
	AddressDomain *Domain `protobuf:"bytes,10,opt,name=address_domain,json=addressDomain,proto3,oneof"`
}

func (*Address_AddressDomain) isAddress_DomainOptional() {}

func (m *Address) GetDomainOptional() isAddress_DomainOptional {
	if m != nil {
		return m.DomainOptional
	}
	return nil
}

func (m *Address) GetAddressDomain() *Domain {
	if x, ok := m.GetDomainOptional().(*Address_AddressDomain); ok {
		return x.AddressDomain
	}
	return nil
}

type isAddress_AddressOptional interface {
	isAddress_AddressOptional()
}

type Address_S3 struct {
	S3 *S3Address `protobuf:"bytes,100,opt,name=s3,proto3,oneof"`
}

type Address_Kafka struct {
	Kafka *KafkaAddress `protobuf:"bytes,200,opt,name=kafka,proto3,oneof"`
}

type Address_Digest struct {
	Digest *Digest `protobuf:"bytes,300,opt,name=digest,proto3,oneof"`
}

type Address_Uuid struct {
	Uuid *UUID `protobuf:"bytes,400,opt,name=uuid,proto3,oneof"`
}

type Address_UserId struct {
	UserId *UserID `protobuf:"bytes,500,opt,name=user_id,json=userId,proto3,oneof"`
}

type Address_Filename struct {
	Filename *Filename `protobuf:"bytes,600,opt,name=filename,proto3,oneof"`
}

type Address_Url struct {
	Url *URL `protobuf:"bytes,700,opt,name=url,proto3,oneof"`
}

type Address_Domain struct {
	Domain *Domain `protobuf:"bytes,800,opt,name=domain,proto3,oneof"`
}

type Address_CustomString struct {
	CustomString string `protobuf:"bytes,900,opt,name=custom_string,json=customString,proto3,oneof"`
}

func (*Address_S3) isAddress_AddressOptional() {}

func (*Address_Kafka) isAddress_AddressOptional() {}

func (*Address_Digest) isAddress_AddressOptional() {}

func (*Address_Uuid) isAddress_AddressOptional() {}

func (*Address_UserId) isAddress_AddressOptional() {}

func (*Address_Filename) isAddress_AddressOptional() {}

func (*Address_Url) isAddress_AddressOptional() {}

func (*Address_Domain) isAddress_AddressOptional() {}

func (*Address_CustomString) isAddress_AddressOptional() {}

func (m *Address) GetAddressOptional() isAddress_AddressOptional {
	if m != nil {
		return m.AddressOptional
	}
	return nil
}

func (m *Address) GetS3() *S3Address {
	if x, ok := m.GetAddressOptional().(*Address_S3); ok {
		return x.S3
	}
	return nil
}

func (m *Address) GetKafka() *KafkaAddress {
	if x, ok := m.GetAddressOptional().(*Address_Kafka); ok {
		return x.Kafka
	}
	return nil
}

func (m *Address) GetDigest() *Digest {
	if x, ok := m.GetAddressOptional().(*Address_Digest); ok {
		return x.Digest
	}
	return nil
}

func (m *Address) GetUuid() *UUID {
	if x, ok := m.GetAddressOptional().(*Address_Uuid); ok {
		return x.Uuid
	}
	return nil
}

func (m *Address) GetUserId() *UserID {
	if x, ok := m.GetAddressOptional().(*Address_UserId); ok {
		return x.UserId
	}
	return nil
}

func (m *Address) GetFilename() *Filename {
	if x, ok := m.GetAddressOptional().(*Address_Filename); ok {
		return x.Filename
	}
	return nil
}

func (m *Address) GetUrl() *URL {
	if x, ok := m.GetAddressOptional().(*Address_Url); ok {
		return x.Url
	}
	return nil
}

func (m *Address) GetDomain() *Domain {
	if x, ok := m.GetAddressOptional().(*Address_Domain); ok {
		return x.Domain
	}
	return nil
}

func (m *Address) GetCustomString() string {
	if x, ok := m.GetAddressOptional().(*Address_CustomString); ok {
		return x.CustomString
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Address) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Address_AddressDomain)(nil),
		(*Address_S3)(nil),
		(*Address_Kafka)(nil),
		(*Address_Digest)(nil),
		(*Address_Uuid)(nil),
		(*Address_UserId)(nil),
		(*Address_Filename)(nil),
		(*Address_Url)(nil),
		(*Address_Domain)(nil),
		(*Address_CustomString)(nil),
	}
}

func init() {
	proto.RegisterType((*Address)(nil), "atlas.Address")
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e) }

var fileDescriptor_982c640dad8fe78e = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4e, 0xf2, 0x40,
	0x14, 0xc7, 0xbf, 0xa1, 0x50, 0x3e, 0x1e, 0x14, 0x70, 0xd8, 0x4c, 0xd8, 0x48, 0x58, 0x18, 0x4c,
	0x4c, 0x17, 0x36, 0x71, 0x2f, 0x21, 0xa6, 0x44, 0x57, 0x43, 0x58, 0x37, 0xa3, 0x53, 0x48, 0x43,
	0x69, 0xc9, 0x4c, 0x7b, 0x03, 0x0f, 0xe0, 0x11, 0x3c, 0x80, 0x47, 0xf0, 0x00, 0x9e, 0xc0, 0x78,
	0x10, 0x0f, 0x60, 0xfa, 0x66, 0xaa, 0xc1, 0xb8, 0x9b, 0xf9, 0xbd, 0xdf, 0xbc, 0xbc, 0xf7, 0x1f,
	0xf0, 0x84, 0x94, 0x2a, 0xd6, 0xda, 0x3f, 0xa8, 0xbc, 0xc8, 0x69, 0x4b, 0x14, 0xa9, 0xd0, 0xe3,
	0x9e, 0x4c, 0xb6, 0xb1, 0x2e, 0x0c, 0x1c, 0xf7, 0x64, 0xbe, 0x17, 0x49, 0x66, 0x6f, 0xfd, 0x4d,
	0x92, 0xc6, 0x99, 0xd8, 0xc7, 0xf6, 0x3e, 0xda, 0x89, 0xcd, 0x4e, 0x44, 0x47, 0x7d, 0xc6, 0x43,
	0x1d, 0xfc, 0x22, 0x9d, 0x52, 0xa5, 0xf6, 0xe8, 0x95, 0x3a, 0x56, 0x51, 0x22, 0xed, 0x15, 0xca,
	0xb2, 0x3e, 0x4f, 0xdf, 0x1d, 0x68, 0x5f, 0x9b, 0x77, 0xf4, 0x0a, 0xfa, 0xb6, 0x45, 0x64, 0x06,
	0x60, 0x30, 0x21, 0xb3, 0xee, 0xa5, 0xe7, 0xe3, 0x90, 0xfe, 0x02, 0x61, 0xf8, 0x8f, 0xd7, 0x3b,
	0x18, 0x40, 0xa7, 0xd0, 0xd0, 0x01, 0x93, 0xe8, 0x0e, 0xad, 0xbb, 0x0a, 0x6c, 0xd7, 0x90, 0xf0,
	0x86, 0x0e, 0xe8, 0x05, 0xb4, 0x70, 0x6c, 0xf6, 0x46, 0xd0, 0x1b, 0x59, 0xef, 0xb6, 0x82, 0x3f,
	0xaa, 0x91, 0xe8, 0x0c, 0x5c, 0x13, 0x08, 0x7b, 0x69, 0x1c, 0x8f, 0x80, 0x34, 0x24, 0xdc, 0xd6,
	0xe9, 0x14, 0x9a, 0xd5, 0x36, 0xec, 0xc9, 0x41, 0xaf, 0x6b, 0xbd, 0xf5, 0x7a, 0xb9, 0x08, 0x09,
	0xc7, 0x1a, 0x3d, 0x87, 0xb6, 0x0d, 0x80, 0x7d, 0x3a, 0x47, 0xed, 0xd6, 0x3a, 0x56, 0x28, 0xba,
	0x95, 0xb0, 0x94, 0xd4, 0x87, 0xff, 0x75, 0xda, 0xec, 0xa3, 0x89, 0xee, 0xc0, 0xba, 0x37, 0x96,
	0x87, 0x84, 0x7f, 0x3b, 0xf4, 0x14, 0x9c, 0x52, 0xa5, 0xec, 0xb5, 0x85, 0x2a, 0xd4, 0x6d, 0xf9,
	0x5d, 0x48, 0x78, 0x55, 0xc1, 0x4d, 0x4c, 0x96, 0xcf, 0xee, 0x5f, 0x61, 0x56, 0x9b, 0x98, 0x14,
	0xcf, 0xc0, 0x7b, 0x28, 0x75, 0x91, 0xef, 0x23, 0x5d, 0xa8, 0x24, 0xdb, 0xb2, 0xc7, 0xf6, 0x84,
	0xcc, 0x3a, 0x21, 0xe1, 0x3d, 0xc3, 0x57, 0x88, 0xe7, 0x27, 0x30, 0x30, 0x2f, 0xa2, 0xfc, 0x50,
	0x24, 0x79, 0x26, 0xd2, 0x39, 0x85, 0x61, 0xfd, 0x71, 0x35, 0xbb, 0x77, 0xf1, 0x7f, 0x83, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x97, 0x17, 0x21, 0x70, 0x02, 0x00, 0x00,
}
