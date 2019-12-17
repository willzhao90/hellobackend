// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

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

type GetHelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHelloRequest) Reset()         { *m = GetHelloRequest{} }
func (m *GetHelloRequest) String() string { return proto.CompactTextString(m) }
func (*GetHelloRequest) ProtoMessage()    {}
func (*GetHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *GetHelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHelloRequest.Unmarshal(m, b)
}
func (m *GetHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHelloRequest.Marshal(b, m, deterministic)
}
func (m *GetHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHelloRequest.Merge(m, src)
}
func (m *GetHelloRequest) XXX_Size() int {
	return xxx_messageInfo_GetHelloRequest.Size(m)
}
func (m *GetHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHelloRequest proto.InternalMessageInfo

func (m *GetHelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetHelloResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHelloResponse) Reset()         { *m = GetHelloResponse{} }
func (m *GetHelloResponse) String() string { return proto.CompactTextString(m) }
func (*GetHelloResponse) ProtoMessage()    {}
func (*GetHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *GetHelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHelloResponse.Unmarshal(m, b)
}
func (m *GetHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHelloResponse.Marshal(b, m, deterministic)
}
func (m *GetHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHelloResponse.Merge(m, src)
}
func (m *GetHelloResponse) XXX_Size() int {
	return xxx_messageInfo_GetHelloResponse.Size(m)
}
func (m *GetHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHelloResponse proto.InternalMessageInfo

func (m *GetHelloResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHelloRequest)(nil), "hello.GetHelloRequest")
	proto.RegisterType((*GetHelloResponse)(nil), "hello.GetHelloResponse")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x54, 0xb9, 0xf8, 0xdd,
	0x53, 0x4b, 0x3c, 0x40, 0xec, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x8d, 0x4b,
	0x00, 0xa1, 0xac, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0x9b, 0x3a, 0x23, 0x5f, 0x2e, 0x1e, 0xb0,
	0xa2, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x5b, 0x2e, 0x2e, 0x97, 0x7c, 0x98, 0x4e,
	0x21, 0x31, 0x3d, 0x88, 0x0b, 0xd0, 0x6c, 0x94, 0x12, 0xc7, 0x10, 0x87, 0x58, 0x91, 0xc4, 0x06,
	0x76, 0xab, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x13, 0xed, 0x62, 0xfb, 0xba, 0x00, 0x00, 0x00,
}
