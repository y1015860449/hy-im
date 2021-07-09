// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im.websocket.proto

package im_app

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

type WsMsg struct {
	Cmd                  int32    `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Retry                int32    `protobuf:"varint,2,opt,name=Retry,proto3" json:"Retry,omitempty"`
	SeqNum               int32    `protobuf:"varint,3,opt,name=seqNum,proto3" json:"seqNum,omitempty"`
	Content              []byte   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WsMsg) Reset()         { *m = WsMsg{} }
func (m *WsMsg) String() string { return proto.CompactTextString(m) }
func (*WsMsg) ProtoMessage()    {}
func (*WsMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9406e7cd387c966, []int{0}
}

func (m *WsMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WsMsg.Unmarshal(m, b)
}
func (m *WsMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WsMsg.Marshal(b, m, deterministic)
}
func (m *WsMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WsMsg.Merge(m, src)
}
func (m *WsMsg) XXX_Size() int {
	return xxx_messageInfo_WsMsg.Size(m)
}
func (m *WsMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_WsMsg.DiscardUnknown(m)
}

var xxx_messageInfo_WsMsg proto.InternalMessageInfo

func (m *WsMsg) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *WsMsg) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

func (m *WsMsg) GetSeqNum() int32 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

func (m *WsMsg) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*WsMsg)(nil), "im.app.wsMsg")
}

func init() { proto.RegisterFile("im.websocket.proto", fileDescriptor_e9406e7cd387c966) }

var fileDescriptor_e9406e7cd387c966 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcc, 0xd5, 0x2b,
	0x4f, 0x4d, 0x2a, 0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0xcb, 0xcc, 0xd5, 0x4b, 0x2c, 0x28, 0x50, 0x4a, 0xe4, 0x62, 0x2d, 0x2f, 0xf6, 0x2d, 0x4e, 0x17,
	0x12, 0xe0, 0x62, 0x4e, 0xce, 0x4d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0x31, 0x85,
	0x44, 0xb8, 0x58, 0x83, 0x52, 0x4b, 0x8a, 0x2a, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0x90, 0x18,
	0x17, 0x5b, 0x71, 0x6a, 0xa1, 0x5f, 0x69, 0xae, 0x04, 0x33, 0x58, 0x18, 0xca, 0x13, 0x92, 0xe0,
	0x62, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x91, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x09, 0x82,
	0x71, 0x93, 0xd8, 0xc0, 0x36, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x87, 0xa9, 0x94, 0x9b,
	0x87, 0x00, 0x00, 0x00,
}
