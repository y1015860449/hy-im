// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im.p2p.proto

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

////////////////////////////
// cmd_p2p_msg = 0x2001;
// cmd_p2p_msg_deliver = 0x2004;
///////////////////////////
type P2PMsg struct {
	ToId                 int64    `protobuf:"varint,1,opt,name=toId,proto3" json:"toId,omitempty"`
	FromId               int64    `protobuf:"varint,2,opt,name=fromId,proto3" json:"fromId,omitempty"`
	Priority             int32    `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	ClientMsgId          string   `protobuf:"bytes,4,opt,name=clientMsgId,proto3" json:"clientMsgId,omitempty"`
	MsgId                string   `protobuf:"bytes,5,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Encrypt              int32    `protobuf:"varint,6,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	Content              []byte   `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	MsgTime              int64    `protobuf:"varint,8,opt,name=msgTime,proto3" json:"msgTime,omitempty"`
	Extend               []byte   `protobuf:"bytes,9,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P2PMsg) Reset()         { *m = P2PMsg{} }
func (m *P2PMsg) String() string { return proto.CompactTextString(m) }
func (*P2PMsg) ProtoMessage()    {}
func (*P2PMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_88f2e254ad685a22, []int{0}
}

func (m *P2PMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PMsg.Unmarshal(m, b)
}
func (m *P2PMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PMsg.Marshal(b, m, deterministic)
}
func (m *P2PMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PMsg.Merge(m, src)
}
func (m *P2PMsg) XXX_Size() int {
	return xxx_messageInfo_P2PMsg.Size(m)
}
func (m *P2PMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PMsg.DiscardUnknown(m)
}

var xxx_messageInfo_P2PMsg proto.InternalMessageInfo

func (m *P2PMsg) GetToId() int64 {
	if m != nil {
		return m.ToId
	}
	return 0
}

func (m *P2PMsg) GetFromId() int64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *P2PMsg) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *P2PMsg) GetClientMsgId() string {
	if m != nil {
		return m.ClientMsgId
	}
	return ""
}

func (m *P2PMsg) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *P2PMsg) GetEncrypt() int32 {
	if m != nil {
		return m.Encrypt
	}
	return 0
}

func (m *P2PMsg) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *P2PMsg) GetMsgTime() int64 {
	if m != nil {
		return m.MsgTime
	}
	return 0
}

func (m *P2PMsg) GetExtend() []byte {
	if m != nil {
		return m.Extend
	}
	return nil
}

///////////////////////////////
// cmd_p2p_msg_ack = 0x2002;
// cmd_p2p_msg_notify_ack = 0x2005;
// cmd_p2p_msg_read_ack = 0x2008;
///////////////////////////////
type P2PAck struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ClientMsgId          string   `protobuf:"bytes,2,opt,name=clientMsgId,proto3" json:"clientMsgId,omitempty"`
	MsgId                string   `protobuf:"bytes,3,opt,name=msgId,proto3" json:"msgId,omitempty"`
	MsgTime              int64    `protobuf:"varint,4,opt,name=msgTime,proto3" json:"msgTime,omitempty"`
	ErrCode              int32    `protobuf:"varint,5,opt,name=errCode,proto3" json:"errCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P2PAck) Reset()         { *m = P2PAck{} }
func (m *P2PAck) String() string { return proto.CompactTextString(m) }
func (*P2PAck) ProtoMessage()    {}
func (*P2PAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_88f2e254ad685a22, []int{1}
}

func (m *P2PAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PAck.Unmarshal(m, b)
}
func (m *P2PAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PAck.Marshal(b, m, deterministic)
}
func (m *P2PAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PAck.Merge(m, src)
}
func (m *P2PAck) XXX_Size() int {
	return xxx_messageInfo_P2PAck.Size(m)
}
func (m *P2PAck) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PAck.DiscardUnknown(m)
}

var xxx_messageInfo_P2PAck proto.InternalMessageInfo

func (m *P2PAck) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *P2PAck) GetClientMsgId() string {
	if m != nil {
		return m.ClientMsgId
	}
	return ""
}

func (m *P2PAck) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *P2PAck) GetMsgTime() int64 {
	if m != nil {
		return m.MsgTime
	}
	return 0
}

func (m *P2PAck) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

/////////////////////////////////////
// cmd_p2p_msg_deliver_ack = 0x2003;
// cmd_p2p_msg_notify_ack = 0x2005;
// cmd_p2p_msg_read_deliver_ack = 0x2009;
/////////////////////////////////////
type P2PDeliverAck struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FromId               int64    `protobuf:"varint,2,opt,name=fromId,proto3" json:"fromId,omitempty"`
	ClientMsgId          string   `protobuf:"bytes,3,opt,name=clientMsgId,proto3" json:"clientMsgId,omitempty"`
	MsgId                string   `protobuf:"bytes,4,opt,name=msgId,proto3" json:"msgId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P2PDeliverAck) Reset()         { *m = P2PDeliverAck{} }
func (m *P2PDeliverAck) String() string { return proto.CompactTextString(m) }
func (*P2PDeliverAck) ProtoMessage()    {}
func (*P2PDeliverAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_88f2e254ad685a22, []int{2}
}

func (m *P2PDeliverAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PDeliverAck.Unmarshal(m, b)
}
func (m *P2PDeliverAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PDeliverAck.Marshal(b, m, deterministic)
}
func (m *P2PDeliverAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PDeliverAck.Merge(m, src)
}
func (m *P2PDeliverAck) XXX_Size() int {
	return xxx_messageInfo_P2PDeliverAck.Size(m)
}
func (m *P2PDeliverAck) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PDeliverAck.DiscardUnknown(m)
}

var xxx_messageInfo_P2PDeliverAck proto.InternalMessageInfo

func (m *P2PDeliverAck) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *P2PDeliverAck) GetFromId() int64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *P2PDeliverAck) GetClientMsgId() string {
	if m != nil {
		return m.ClientMsgId
	}
	return ""
}

func (m *P2PDeliverAck) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

////////////////////////////////
// cmd_p2p_msg_notify = 0x2006;     // 消息已接受通知（服务器发送）
////////////////////////////////
type P2PMsgNotify struct {
	ToId                 int64    `protobuf:"varint,1,opt,name=toId,proto3" json:"toId,omitempty"`
	FromId               int64    `protobuf:"varint,2,opt,name=fromId,proto3" json:"fromId,omitempty"`
	ClientMsgId          string   `protobuf:"bytes,3,opt,name=clientMsgId,proto3" json:"clientMsgId,omitempty"`
	MsgId                string   `protobuf:"bytes,4,opt,name=msgId,proto3" json:"msgId,omitempty"`
	MsgTime              int64    `protobuf:"varint,5,opt,name=msgTime,proto3" json:"msgTime,omitempty"`
	NotifyId             string   `protobuf:"bytes,6,opt,name=notifyId,proto3" json:"notifyId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P2PMsgNotify) Reset()         { *m = P2PMsgNotify{} }
func (m *P2PMsgNotify) String() string { return proto.CompactTextString(m) }
func (*P2PMsgNotify) ProtoMessage()    {}
func (*P2PMsgNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_88f2e254ad685a22, []int{3}
}

func (m *P2PMsgNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PMsgNotify.Unmarshal(m, b)
}
func (m *P2PMsgNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PMsgNotify.Marshal(b, m, deterministic)
}
func (m *P2PMsgNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PMsgNotify.Merge(m, src)
}
func (m *P2PMsgNotify) XXX_Size() int {
	return xxx_messageInfo_P2PMsgNotify.Size(m)
}
func (m *P2PMsgNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PMsgNotify.DiscardUnknown(m)
}

var xxx_messageInfo_P2PMsgNotify proto.InternalMessageInfo

func (m *P2PMsgNotify) GetToId() int64 {
	if m != nil {
		return m.ToId
	}
	return 0
}

func (m *P2PMsgNotify) GetFromId() int64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *P2PMsgNotify) GetClientMsgId() string {
	if m != nil {
		return m.ClientMsgId
	}
	return ""
}

func (m *P2PMsgNotify) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *P2PMsgNotify) GetMsgTime() int64 {
	if m != nil {
		return m.MsgTime
	}
	return 0
}

func (m *P2PMsgNotify) GetNotifyId() string {
	if m != nil {
		return m.NotifyId
	}
	return ""
}

//////////////////////////////
// cmd_p2p_msg_read = 0x2007;   // 消息已读
// cmd_p2p_msg_read_deliver = 0x200a;
//////////////////////////////
type P2PMsgRead struct {
	ToId                 int64    `protobuf:"varint,1,opt,name=toId,proto3" json:"toId,omitempty"`
	FromId               int64    `protobuf:"varint,2,opt,name=fromId,proto3" json:"fromId,omitempty"`
	ClientMsgId          string   `protobuf:"bytes,3,opt,name=clientMsgId,proto3" json:"clientMsgId,omitempty"`
	MsgId                string   `protobuf:"bytes,4,opt,name=msgId,proto3" json:"msgId,omitempty"`
	MsgIds               []string `protobuf:"bytes,5,rep,name=msgIds,proto3" json:"msgIds,omitempty"`
	MsgTime              int64    `protobuf:"varint,6,opt,name=msgTime,proto3" json:"msgTime,omitempty"`
	Extend               []byte   `protobuf:"bytes,7,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P2PMsgRead) Reset()         { *m = P2PMsgRead{} }
func (m *P2PMsgRead) String() string { return proto.CompactTextString(m) }
func (*P2PMsgRead) ProtoMessage()    {}
func (*P2PMsgRead) Descriptor() ([]byte, []int) {
	return fileDescriptor_88f2e254ad685a22, []int{4}
}

func (m *P2PMsgRead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PMsgRead.Unmarshal(m, b)
}
func (m *P2PMsgRead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PMsgRead.Marshal(b, m, deterministic)
}
func (m *P2PMsgRead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PMsgRead.Merge(m, src)
}
func (m *P2PMsgRead) XXX_Size() int {
	return xxx_messageInfo_P2PMsgRead.Size(m)
}
func (m *P2PMsgRead) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PMsgRead.DiscardUnknown(m)
}

var xxx_messageInfo_P2PMsgRead proto.InternalMessageInfo

func (m *P2PMsgRead) GetToId() int64 {
	if m != nil {
		return m.ToId
	}
	return 0
}

func (m *P2PMsgRead) GetFromId() int64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *P2PMsgRead) GetClientMsgId() string {
	if m != nil {
		return m.ClientMsgId
	}
	return ""
}

func (m *P2PMsgRead) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *P2PMsgRead) GetMsgIds() []string {
	if m != nil {
		return m.MsgIds
	}
	return nil
}

func (m *P2PMsgRead) GetMsgTime() int64 {
	if m != nil {
		return m.MsgTime
	}
	return 0
}

func (m *P2PMsgRead) GetExtend() []byte {
	if m != nil {
		return m.Extend
	}
	return nil
}

func init() {
	proto.RegisterType((*P2PMsg)(nil), "im.app.p2pMsg")
	proto.RegisterType((*P2PAck)(nil), "im.app.p2pAck")
	proto.RegisterType((*P2PDeliverAck)(nil), "im.app.p2pDeliverAck")
	proto.RegisterType((*P2PMsgNotify)(nil), "im.app.p2pMsgNotify")
	proto.RegisterType((*P2PMsgRead)(nil), "im.app.p2pMsgRead")
}

func init() { proto.RegisterFile("im.p2p.proto", fileDescriptor_88f2e254ad685a22) }

var fileDescriptor_88f2e254ad685a22 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4d, 0x6e, 0xf2, 0x30,
	0x10, 0x86, 0x65, 0x92, 0x18, 0x32, 0x1f, 0xdf, 0xc6, 0xaa, 0x90, 0xc5, 0x2a, 0xca, 0x2a, 0x2b,
	0x16, 0xf4, 0x04, 0x55, 0xbb, 0xc9, 0x82, 0x2e, 0xac, 0x5e, 0x80, 0xc6, 0x06, 0x59, 0x25, 0xb1,
	0xe5, 0xb8, 0x3f, 0xdc, 0xa1, 0x07, 0xe9, 0x19, 0x7a, 0xa8, 0x9e, 0xa1, 0xca, 0x24, 0x41, 0x01,
	0x09, 0xa4, 0x76, 0xd1, 0x9d, 0x1f, 0x4f, 0x7e, 0x9e, 0xd7, 0x33, 0x86, 0xa9, 0x2e, 0x17, 0x76,
	0x69, 0x17, 0xd6, 0x19, 0x6f, 0x18, 0xd5, 0xe5, 0x62, 0x6d, 0x6d, 0xfa, 0x45, 0x80, 0xda, 0xa5,
	0x5d, 0xd5, 0x5b, 0xc6, 0x20, 0xf4, 0x26, 0x97, 0x9c, 0x24, 0x24, 0x0b, 0x04, 0xae, 0xd9, 0x0c,
	0xe8, 0xc6, 0x99, 0x32, 0x97, 0x7c, 0x84, 0xbb, 0x1d, 0xb1, 0x39, 0x4c, 0xac, 0xd3, 0xc6, 0x69,
	0xbf, 0xe7, 0x41, 0x42, 0xb2, 0x48, 0x1c, 0x98, 0x25, 0xf0, 0xaf, 0xd8, 0x69, 0x55, 0xf9, 0x55,
	0xbd, 0xcd, 0x25, 0x0f, 0x13, 0x92, 0xc5, 0x62, 0xb8, 0xc5, 0xae, 0x20, 0x2a, 0xb1, 0x16, 0x61,
	0xad, 0x05, 0xc6, 0x61, 0xac, 0xaa, 0xc2, 0xed, 0xad, 0xe7, 0x14, 0x3f, 0xd9, 0x63, 0x53, 0x29,
	0x4c, 0xe5, 0x55, 0xe5, 0xf9, 0x38, 0x21, 0xd9, 0x54, 0xf4, 0xd8, 0x54, 0xca, 0x7a, 0xfb, 0xa0,
	0x4b, 0xc5, 0x27, 0x28, 0xd8, 0x63, 0x63, 0xae, 0xde, 0xbc, 0xaa, 0x24, 0x8f, 0xf1, 0x95, 0x8e,
	0xd2, 0xf7, 0x36, 0xf0, 0x4d, 0xf1, 0xd4, 0x3c, 0xf2, 0x5c, 0x2b, 0x77, 0x88, 0xdc, 0xd1, 0x69,
	0x80, 0xd1, 0x85, 0x00, 0xc1, 0x49, 0x80, 0x5e, 0x26, 0x3c, 0x96, 0x69, 0xa2, 0x39, 0x77, 0x6b,
	0xa4, 0xc2, 0xc8, 0x4d, 0xb4, 0x16, 0xd3, 0x57, 0xf8, 0x6f, 0x97, 0xf6, 0x4e, 0xed, 0xf4, 0x8b,
	0x72, 0x97, 0xa4, 0xce, 0x75, 0xe2, 0x44, 0x36, 0xb8, 0x20, 0x1b, 0x0e, 0x64, 0xd3, 0x0f, 0x02,
	0xd3, 0xb6, 0xf1, 0xf7, 0xc6, 0xeb, 0xcd, 0xfe, 0x47, 0xed, 0xff, 0xe5, 0x4f, 0x87, 0x27, 0x14,
	0x1d, 0x9f, 0xd0, 0x1c, 0x26, 0x15, 0x7a, 0xe4, 0x12, 0xbb, 0x1f, 0x8b, 0x03, 0xa7, 0x9f, 0x04,
	0xa0, 0x55, 0x15, 0x6a, 0x2d, 0xff, 0x44, 0x74, 0x06, 0x14, 0x17, 0x35, 0x8f, 0x92, 0x20, 0x8b,
	0x45, 0x47, 0xc3, 0x00, 0xf4, 0xdc, 0xbc, 0x8d, 0x87, 0xf3, 0xf6, 0x48, 0xf1, 0xbe, 0x5d, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x4c, 0xa9, 0xd3, 0x7f, 0x03, 0x00, 0x00,
}
