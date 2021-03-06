// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mq.proto

package im_mq

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

type PushGroupMsg struct {
	Command              int32    `protobuf:"varint,1,opt,name=command,proto3" json:"command,omitempty"`
	GroupId              int64    `protobuf:"varint,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	UserId               int64    `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
	OtherId              []int64  `protobuf:"varint,5,rep,packed,name=otherId,proto3" json:"otherId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushGroupMsg) Reset()         { *m = PushGroupMsg{} }
func (m *PushGroupMsg) String() string { return proto.CompactTextString(m) }
func (*PushGroupMsg) ProtoMessage()    {}
func (*PushGroupMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_7caa8199c0db1966, []int{0}
}

func (m *PushGroupMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushGroupMsg.Unmarshal(m, b)
}
func (m *PushGroupMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushGroupMsg.Marshal(b, m, deterministic)
}
func (m *PushGroupMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushGroupMsg.Merge(m, src)
}
func (m *PushGroupMsg) XXX_Size() int {
	return xxx_messageInfo_PushGroupMsg.Size(m)
}
func (m *PushGroupMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PushGroupMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PushGroupMsg proto.InternalMessageInfo

func (m *PushGroupMsg) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *PushGroupMsg) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PushGroupMsg) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PushGroupMsg) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PushGroupMsg) GetOtherId() []int64 {
	if m != nil {
		return m.OtherId
	}
	return nil
}

type PushP2PMsg struct {
	Command              int32    `protobuf:"varint,1,opt,name=command,proto3" json:"command,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ToId                 int64    `protobuf:"varint,3,opt,name=toId,proto3" json:"toId,omitempty"`
	Content              []byte   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushP2PMsg) Reset()         { *m = PushP2PMsg{} }
func (m *PushP2PMsg) String() string { return proto.CompactTextString(m) }
func (*PushP2PMsg) ProtoMessage()    {}
func (*PushP2PMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_7caa8199c0db1966, []int{1}
}

func (m *PushP2PMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushP2PMsg.Unmarshal(m, b)
}
func (m *PushP2PMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushP2PMsg.Marshal(b, m, deterministic)
}
func (m *PushP2PMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushP2PMsg.Merge(m, src)
}
func (m *PushP2PMsg) XXX_Size() int {
	return xxx_messageInfo_PushP2PMsg.Size(m)
}
func (m *PushP2PMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PushP2PMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PushP2PMsg proto.InternalMessageInfo

func (m *PushP2PMsg) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *PushP2PMsg) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PushP2PMsg) GetToId() int64 {
	if m != nil {
		return m.ToId
	}
	return 0
}

func (m *PushP2PMsg) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*PushGroupMsg)(nil), "im.mq.pushGroupMsg")
	proto.RegisterType((*PushP2PMsg)(nil), "im.mq.pushP2pMsg")
}

func init() { proto.RegisterFile("mq.proto", fileDescriptor_7caa8199c0db1966) }

var fileDescriptor_7caa8199c0db1966 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2d, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcc, 0xd5, 0xcb, 0x2d, 0x54, 0xea, 0x63, 0xe4, 0xe2,
	0x29, 0x28, 0x2d, 0xce, 0x70, 0x2f, 0xca, 0x2f, 0x2d, 0xf0, 0x2d, 0x4e, 0x17, 0x92, 0xe0, 0x62,
	0x4f, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x71,
	0x41, 0x32, 0xe9, 0x20, 0x55, 0x9e, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0x2e,
	0x44, 0x4f, 0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x8c, 0x2b,
	0x24, 0xc6, 0xc5, 0x56, 0x5a, 0x9c, 0x5a, 0xe4, 0x99, 0x22, 0xc1, 0x02, 0xd6, 0x02, 0xe5, 0x81,
	0x74, 0xe4, 0x97, 0x64, 0x80, 0x25, 0x58, 0x15, 0x98, 0x41, 0x66, 0x41, 0xb9, 0x4a, 0x39, 0x5c,
	0x5c, 0x20, 0xf7, 0x04, 0x18, 0x11, 0x70, 0x0d, 0xc2, 0x64, 0x26, 0x14, 0x93, 0x85, 0xb8, 0x58,
	0x4a, 0xf2, 0x3d, 0x53, 0xc0, 0x0e, 0x61, 0x0e, 0x02, 0xb3, 0x91, 0xdd, 0xc7, 0x82, 0xe2, 0xbe,
	0x24, 0x36, 0x70, 0x60, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x42, 0xc0, 0xf8, 0x18,
	0x01, 0x00, 0x00,
}
