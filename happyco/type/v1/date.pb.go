// Code generated by protoc-gen-go.
// source: happyco/type/v1/date.proto
// DO NOT EDIT!

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	happyco/type/v1/date.proto
	happyco/type/v1/event.proto
	happyco/type/v1/integration_id.proto
	happyco/type/v1/paging.proto
	happyco/type/v1/varia.proto

It has these top-level messages:
	Date
	DateTime
	EventHandlerOptions
	Event
	EventAck
	IntegrationID
	Paging
	PagingResponse
	StringArray
	FloatArray
*/
package v1

import proto "github.com/happy-co/happyapis-golang/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Date struct {
	Year  int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type DateTime struct {
	Year   int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	Month  int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	Day    int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
	Hour24 int32 `protobuf:"varint,4,opt,name=hour24" json:"hour24,omitempty"`
	Minute int32 `protobuf:"varint,5,opt,name=minute" json:"minute,omitempty"`
	Second int32 `protobuf:"varint,6,opt,name=second" json:"second,omitempty"`
}

func (m *DateTime) Reset()                    { *m = DateTime{} }
func (m *DateTime) String() string            { return proto.CompactTextString(m) }
func (*DateTime) ProtoMessage()               {}
func (*DateTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DateTime) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *DateTime) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *DateTime) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *DateTime) GetHour24() int32 {
	if m != nil {
		return m.Hour24
	}
	return 0
}

func (m *DateTime) GetMinute() int32 {
	if m != nil {
		return m.Minute
	}
	return 0
}

func (m *DateTime) GetSecond() int32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func init() {
	proto.RegisterType((*Date)(nil), "happyco.type.v1.Date")
	proto.RegisterType((*DateTime)(nil), "happyco.type.v1.DateTime")
}

func init() { proto.RegisterFile("happyco/type/v1/date.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x48, 0x2c, 0x28,
	0xa8, 0x4c, 0xce, 0xd7, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0x49, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0xca, 0xe9, 0x81, 0xe4, 0xf4, 0xca, 0x0c,
	0x95, 0x9c, 0xb8, 0x58, 0x5c, 0x12, 0x4b, 0x52, 0x85, 0x84, 0xb8, 0x58, 0x2a, 0x53, 0x13, 0x8b,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xdc, 0xfc, 0xbc,
	0x92, 0x0c, 0x09, 0x26, 0xb0, 0x20, 0x84, 0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x92, 0x58, 0x29, 0xc1,
	0x0c, 0x16, 0x03, 0x31, 0x95, 0xfa, 0x18, 0xb9, 0x38, 0x40, 0x86, 0x84, 0x64, 0xe6, 0x52, 0x64,
	0x90, 0x90, 0x18, 0x17, 0x5b, 0x46, 0x7e, 0x69, 0x91, 0x91, 0x89, 0x04, 0x0b, 0x58, 0x10, 0xca,
	0x03, 0x89, 0xe7, 0x66, 0xe6, 0x95, 0x96, 0xa4, 0x4a, 0xb0, 0x42, 0xc4, 0x21, 0x3c, 0x90, 0x78,
	0x71, 0x6a, 0x72, 0x7e, 0x5e, 0x8a, 0x04, 0x1b, 0x44, 0x1c, 0xc2, 0x73, 0x12, 0x8c, 0xe2, 0x47,
	0x0b, 0x83, 0x24, 0x36, 0xb0, 0xff, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x18, 0x53,
	0x44, 0x1d, 0x01, 0x00, 0x00,
}
