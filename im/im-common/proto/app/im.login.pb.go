// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im.login.proto

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

type Login struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LoginType            int32    `protobuf:"varint,2,opt,name=loginType,proto3" json:"loginType,omitempty"`
	RoleType             int32    `protobuf:"varint,3,opt,name=roleType,proto3" json:"roleType,omitempty"`
	LoginToken           string   `protobuf:"bytes,4,opt,name=loginToken,proto3" json:"loginToken,omitempty"`
	DeviceToken          string   `protobuf:"bytes,5,opt,name=deviceToken,proto3" json:"deviceToken,omitempty"`
	Extend               string   `protobuf:"bytes,6,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_03f939dffd6ea369, []int{0}
}

func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (m *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(m, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Login) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *Login) GetRoleType() int32 {
	if m != nil {
		return m.RoleType
	}
	return 0
}

func (m *Login) GetLoginToken() string {
	if m != nil {
		return m.LoginToken
	}
	return ""
}

func (m *Login) GetDeviceToken() string {
	if m != nil {
		return m.DeviceToken
	}
	return ""
}

func (m *Login) GetExtend() string {
	if m != nil {
		return m.Extend
	}
	return ""
}

type LoginAck struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ErrCode              int32    `protobuf:"varint,2,opt,name=errCode,proto3" json:"errCode,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginAck) Reset()         { *m = LoginAck{} }
func (m *LoginAck) String() string { return proto.CompactTextString(m) }
func (*LoginAck) ProtoMessage()    {}
func (*LoginAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_03f939dffd6ea369, []int{1}
}

func (m *LoginAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginAck.Unmarshal(m, b)
}
func (m *LoginAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginAck.Marshal(b, m, deterministic)
}
func (m *LoginAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAck.Merge(m, src)
}
func (m *LoginAck) XXX_Size() int {
	return xxx_messageInfo_LoginAck.Size(m)
}
func (m *LoginAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAck.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAck proto.InternalMessageInfo

func (m *LoginAck) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LoginAck) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *LoginAck) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type KickOut struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ReasonCode           int32    `protobuf:"varint,2,opt,name=reasonCode,proto3" json:"reasonCode,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KickOut) Reset()         { *m = KickOut{} }
func (m *KickOut) String() string { return proto.CompactTextString(m) }
func (*KickOut) ProtoMessage()    {}
func (*KickOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_03f939dffd6ea369, []int{2}
}

func (m *KickOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickOut.Unmarshal(m, b)
}
func (m *KickOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickOut.Marshal(b, m, deterministic)
}
func (m *KickOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickOut.Merge(m, src)
}
func (m *KickOut) XXX_Size() int {
	return xxx_messageInfo_KickOut.Size(m)
}
func (m *KickOut) XXX_DiscardUnknown() {
	xxx_messageInfo_KickOut.DiscardUnknown(m)
}

var xxx_messageInfo_KickOut proto.InternalMessageInfo

func (m *KickOut) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *KickOut) GetReasonCode() int32 {
	if m != nil {
		return m.ReasonCode
	}
	return 0
}

func (m *KickOut) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Login)(nil), "im.app.login")
	proto.RegisterType((*LoginAck)(nil), "im.app.loginAck")
	proto.RegisterType((*KickOut)(nil), "im.app.kickOut")
}

func init() { proto.RegisterFile("im.login.proto", fileDescriptor_03f939dffd6ea369) }

var fileDescriptor_03f939dffd6ea369 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x14, 0xc4, 0x89, 0xeb, 0xa6, 0xed, 0x13, 0x3c, 0xe4, 0x20, 0x41, 0xa4, 0x84, 0x3d, 0xed, 0x69,
	0x2f, 0x7e, 0x02, 0xf1, 0xe4, 0x49, 0x08, 0x9e, 0xbc, 0xc8, 0xba, 0x79, 0x48, 0x48, 0xf3, 0x87,
	0x6c, 0x2a, 0xfa, 0xb9, 0xfc, 0x82, 0xd2, 0xd7, 0xd6, 0xe6, 0x52, 0x8f, 0xf3, 0x9b, 0x61, 0x32,
	0xe1, 0xc1, 0xb5, 0xf5, 0xc3, 0x26, 0x7e, 0xd8, 0x30, 0xa4, 0x1c, 0x4b, 0x14, 0xdc, 0xfa, 0x61,
	0x4c, 0xa9, 0xfb, 0x61, 0xd0, 0x12, 0x17, 0x37, 0xc0, 0xb7, 0x33, 0xe6, 0x27, 0x23, 0x99, 0x62,
	0x7d, 0xa3, 0x0f, 0x4a, 0xdc, 0xc1, 0x8a, 0x02, 0x2f, 0xdf, 0x09, 0xe5, 0x85, 0x62, 0x7d, 0xab,
	0x4f, 0x40, 0xdc, 0xc2, 0x32, 0xc7, 0x0d, 0x92, 0xd9, 0x90, 0xf9, 0xa7, 0xc5, 0x1a, 0x60, 0x1f,
	0x8c, 0x0e, 0x83, 0xbc, 0x54, 0xac, 0x5f, 0xe9, 0x8a, 0x08, 0x05, 0x57, 0x06, 0x3f, 0xed, 0x84,
	0xfb, 0x40, 0x4b, 0x81, 0x1a, 0xed, 0x36, 0xe1, 0x57, 0xc1, 0x60, 0x24, 0x27, 0xf3, 0xa0, 0xba,
	0x57, 0x58, 0x52, 0xcf, 0xc3, 0xe4, 0xce, 0xee, 0x96, 0xb0, 0xc0, 0x9c, 0x1f, 0xa3, 0x39, 0xae,
	0x3e, 0xca, 0xdd, 0x8f, 0x8a, 0xf5, 0x38, 0x97, 0xd1, 0x27, 0x1a, 0xdd, 0xe8, 0x13, 0xe8, 0xde,
	0x60, 0xe1, 0xec, 0xe4, 0x9e, 0xb7, 0xe5, 0x6c, 0xf5, 0x1a, 0x20, 0xe3, 0x38, 0xc7, 0x50, 0xb5,
	0x57, 0xe4, 0xff, 0x07, 0xde, 0x39, 0x5d, 0xe0, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x8d,
	0x8d, 0x07, 0x93, 0x01, 0x00, 0x00,
}