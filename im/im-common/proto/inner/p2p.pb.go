// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p.proto

package inner

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

type P2PReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LoginType            int32    `protobuf:"varint,2,opt,name=loginType,proto3" json:"loginType,omitempty"`
	Command              int32    `protobuf:"varint,3,opt,name=command,proto3" json:"command,omitempty"`
	Content              []byte   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Retry                int32    `protobuf:"varint,5,opt,name=retry,proto3" json:"retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P2PReq) Reset()         { *m = P2PReq{} }
func (m *P2PReq) String() string { return proto.CompactTextString(m) }
func (*P2PReq) ProtoMessage()    {}
func (*P2PReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{0}
}

func (m *P2PReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PReq.Unmarshal(m, b)
}
func (m *P2PReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PReq.Marshal(b, m, deterministic)
}
func (m *P2PReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PReq.Merge(m, src)
}
func (m *P2PReq) XXX_Size() int {
	return xxx_messageInfo_P2PReq.Size(m)
}
func (m *P2PReq) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PReq.DiscardUnknown(m)
}

var xxx_messageInfo_P2PReq proto.InternalMessageInfo

func (m *P2PReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *P2PReq) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *P2PReq) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *P2PReq) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *P2PReq) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

type P2PRsp struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LoginType            int32    `protobuf:"varint,2,opt,name=loginType,proto3" json:"loginType,omitempty"`
	Command              int32    `protobuf:"varint,3,opt,name=command,proto3" json:"command,omitempty"`
	Content              []byte   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	SvcErr               int32    `protobuf:"varint,5,opt,name=svcErr,proto3" json:"svcErr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P2PRsp) Reset()         { *m = P2PRsp{} }
func (m *P2PRsp) String() string { return proto.CompactTextString(m) }
func (*P2PRsp) ProtoMessage()    {}
func (*P2PRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{1}
}

func (m *P2PRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PRsp.Unmarshal(m, b)
}
func (m *P2PRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PRsp.Marshal(b, m, deterministic)
}
func (m *P2PRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PRsp.Merge(m, src)
}
func (m *P2PRsp) XXX_Size() int {
	return xxx_messageInfo_P2PRsp.Size(m)
}
func (m *P2PRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PRsp.DiscardUnknown(m)
}

var xxx_messageInfo_P2PRsp proto.InternalMessageInfo

func (m *P2PRsp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *P2PRsp) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *P2PRsp) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *P2PRsp) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *P2PRsp) GetSvcErr() int32 {
	if m != nil {
		return m.SvcErr
	}
	return 0
}

func init() {
	proto.RegisterType((*P2PReq)(nil), "inner.p2pReq")
	proto.RegisterType((*P2PRsp)(nil), "inner.p2pRsp")
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptor_e7fdddb109e6467a) }

var fileDescriptor_e7fdddb109e6467a = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x30, 0x2a, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcc, 0xcb, 0x4b, 0x2d, 0x52, 0xea, 0x60, 0xe4,
	0x62, 0x2b, 0x30, 0x2a, 0x08, 0x4a, 0x2d, 0x14, 0x12, 0xe3, 0x62, 0x2b, 0x2d, 0x4e, 0x2d, 0xf2,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf2, 0x84, 0x64, 0xb8, 0x38, 0x73, 0xf2,
	0xd3, 0x33, 0xf3, 0x42, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x10, 0x02,
	0x42, 0x12, 0x5c, 0xec, 0xc9, 0xf9, 0xb9, 0xb9, 0x89, 0x79, 0x29, 0x12, 0xcc, 0x60, 0x39, 0x18,
	0x17, 0x22, 0x93, 0x57, 0x92, 0x9a, 0x57, 0x22, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe3,
	0x0a, 0x89, 0x70, 0xb1, 0x16, 0xa5, 0x96, 0x14, 0x55, 0x4a, 0xb0, 0x82, 0x75, 0x40, 0x38, 0x4a,
	0x5d, 0x50, 0xa7, 0x14, 0x17, 0xd0, 0xd1, 0x29, 0x62, 0x5c, 0x6c, 0xc5, 0x65, 0xc9, 0xae, 0x45,
	0x45, 0x50, 0xb7, 0x40, 0x79, 0x46, 0x3a, 0x5c, 0xac, 0x9e, 0xb9, 0x01, 0x46, 0x05, 0x42, 0xca,
	0x5c, 0xcc, 0x20, 0x8a, 0x57, 0x0f, 0x1c, 0x5e, 0x7a, 0x90, 0xb0, 0x92, 0x42, 0xe6, 0x16, 0x17,
	0x24, 0xb1, 0x81, 0xc3, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x02, 0x08, 0xb1, 0xf8, 0x60,
	0x01, 0x00, 0x00,
}
