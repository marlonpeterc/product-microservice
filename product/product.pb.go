// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

/*
Package product is a generated protocol buffer package.

It is generated from these files:
	product.proto

It has these top-level messages:
	ProductRequest
	ProductResponse
	ProductFilter
*/
package product

import proto "github.com/golang/protobuf/proto"
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

type ProductRequest struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ShortDesc string `protobuf:"bytes,3,opt,name=shortDesc" json:"shortDesc,omitempty"`
	Desc      string `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
}

func (m *ProductRequest) Reset()                    { *m = ProductRequest{} }
func (m *ProductRequest) String() string            { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()               {}
func (*ProductRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductRequest) GetShortDesc() string {
	if m != nil {
		return m.ShortDesc
	}
	return ""
}

func (m *ProductRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ProductResponse struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *ProductResponse) Reset()                    { *m = ProductResponse{} }
func (m *ProductResponse) String() string            { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()               {}
func (*ProductResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProductResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProductResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ProductFilter struct {
	Keyword string `protobuf:"bytes,1,opt,name=keyword" json:"keyword,omitempty"`
}

func (m *ProductFilter) Reset()                    { *m = ProductFilter{} }
func (m *ProductFilter) String() string            { return proto.CompactTextString(m) }
func (*ProductFilter) ProtoMessage()               {}
func (*ProductFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProductFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "product.ProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "product.ProductResponse")
	proto.RegisterType((*ProductFilter)(nil), "product.ProductFilter")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xde, 0xd4, 0xc5, 0xba, 0x23, 0x5d, 0x61, 0x0e, 0x1a, 0x16, 0x0f, 0x92, 0x93, 0x5e, 0x16,
	0xd1, 0xa3, 0x37, 0x2d, 0x7a, 0x95, 0xbe, 0x41, 0x4d, 0x46, 0x2c, 0x6a, 0x53, 0x33, 0x29, 0xe2,
	0x73, 0xf8, 0xc2, 0x92, 0x34, 0x51, 0xb4, 0x78, 0x9b, 0xef, 0x2f, 0x1f, 0xf9, 0xa0, 0x1a, 0x9c,
	0x35, 0xa3, 0xf6, 0xdb, 0xc1, 0x59, 0x6f, 0xb1, 0x4c, 0x50, 0x3d, 0xc2, 0xfa, 0x7e, 0x3a, 0x1b,
	0x7a, 0x1b, 0x89, 0x3d, 0xae, 0xa1, 0xe8, 0x8c, 0x14, 0x27, 0xe2, 0x74, 0xd5, 0x14, 0x9d, 0x41,
	0x84, 0x65, 0xdf, 0xbe, 0x92, 0x2c, 0x22, 0x13, 0x6f, 0x3c, 0x86, 0x15, 0x3f, 0x59, 0xe7, 0x6b,
	0x62, 0x2d, 0x77, 0xa2, 0xf0, 0x43, 0x84, 0x84, 0x09, 0xc2, 0x72, 0x4a, 0x84, 0x5b, 0x5d, 0xc1,
	0xc1, 0x77, 0x0f, 0x0f, 0xb6, 0x67, 0x9a, 0x15, 0x49, 0x28, 0x79, 0xd4, 0x9a, 0x98, 0x63, 0xd7,
	0x5e, 0x93, 0xa1, 0x3a, 0x83, 0x2a, 0x85, 0x6f, 0xbb, 0x17, 0x4f, 0x2e, 0x58, 0x9f, 0xe9, 0xe3,
	0xdd, 0xba, 0x9c, 0xcf, 0xf0, 0xe2, 0x53, 0x40, 0x99, 0xbc, 0x78, 0x0d, 0xfb, 0x77, 0xe4, 0x13,
	0x62, 0x3c, 0xdc, 0xe6, 0x0d, 0x7e, 0x3d, 0xb6, 0x39, 0xfa, 0xcb, 0xa7, 0x25, 0xd4, 0xe2, 0x5c,
	0x60, 0x0d, 0xd5, 0x8d, 0xa3, 0xd6, 0x53, 0x7e, 0xf4, 0x3f, 0xf7, 0x46, 0xce, 0x85, 0xe9, 0xa3,
	0x6a, 0xf1, 0xb0, 0x1b, 0x57, 0xbf, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x01, 0x86, 0xda, 0x61,
	0x86, 0x01, 0x00, 0x00,
}