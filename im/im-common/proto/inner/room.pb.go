// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room.proto

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

type RoomReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LoginType            int32    `protobuf:"varint,2,opt,name=loginType,proto3" json:"loginType,omitempty"`
	RoleType             int32    `protobuf:"varint,3,opt,name=roleType,proto3" json:"roleType,omitempty"`
	RoomId               int64    `protobuf:"varint,4,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Command              int32    `protobuf:"varint,5,opt,name=command,proto3" json:"command,omitempty"`
	Content              []byte   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Retry                int32    `protobuf:"varint,7,opt,name=retry,proto3" json:"retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomReq) Reset()         { *m = RoomReq{} }
func (m *RoomReq) String() string { return proto.CompactTextString(m) }
func (*RoomReq) ProtoMessage()    {}
func (*RoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{0}
}

func (m *RoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomReq.Unmarshal(m, b)
}
func (m *RoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomReq.Marshal(b, m, deterministic)
}
func (m *RoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomReq.Merge(m, src)
}
func (m *RoomReq) XXX_Size() int {
	return xxx_messageInfo_RoomReq.Size(m)
}
func (m *RoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomReq proto.InternalMessageInfo

func (m *RoomReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RoomReq) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *RoomReq) GetRoleType() int32 {
	if m != nil {
		return m.RoleType
	}
	return 0
}

func (m *RoomReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RoomReq) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *RoomReq) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RoomReq) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

type RoomRsp struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LoginType            int32    `protobuf:"varint,2,opt,name=loginType,proto3" json:"loginType,omitempty"`
	RoleType             int32    `protobuf:"varint,3,opt,name=roleType,proto3" json:"roleType,omitempty"`
	RoomId               int64    `protobuf:"varint,4,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Command              int32    `protobuf:"varint,5,opt,name=command,proto3" json:"command,omitempty"`
	Content              []byte   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	SvcErr               int32    `protobuf:"varint,7,opt,name=svcErr,proto3" json:"svcErr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomRsp) Reset()         { *m = RoomRsp{} }
func (m *RoomRsp) String() string { return proto.CompactTextString(m) }
func (*RoomRsp) ProtoMessage()    {}
func (*RoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{1}
}

func (m *RoomRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomRsp.Unmarshal(m, b)
}
func (m *RoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomRsp.Marshal(b, m, deterministic)
}
func (m *RoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomRsp.Merge(m, src)
}
func (m *RoomRsp) XXX_Size() int {
	return xxx_messageInfo_RoomRsp.Size(m)
}
func (m *RoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RoomRsp proto.InternalMessageInfo

func (m *RoomRsp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RoomRsp) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *RoomRsp) GetRoleType() int32 {
	if m != nil {
		return m.RoleType
	}
	return 0
}

func (m *RoomRsp) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RoomRsp) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *RoomRsp) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RoomRsp) GetSvcErr() int32 {
	if m != nil {
		return m.SvcErr
	}
	return 0
}

func init() {
	proto.RegisterType((*RoomReq)(nil), "inner.roomReq")
	proto.RegisterType((*RoomRsp)(nil), "inner.roomRsp")
}

func init() { proto.RegisterFile("room.proto", fileDescriptor_c5fd27dd97284ef4) }

var fileDescriptor_c5fd27dd97284ef4 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x91, 0x41, 0x4a, 0xc6, 0x30,
	0x14, 0x84, 0x89, 0xff, 0x9f, 0x54, 0x1f, 0xe2, 0x22, 0xc8, 0x4f, 0x28, 0x2e, 0x4a, 0x17, 0xd2,
	0x55, 0x11, 0x3d, 0x83, 0x8b, 0x6e, 0x83, 0x17, 0xd0, 0x36, 0x48, 0xa1, 0xc9, 0x8b, 0x2f, 0x51,
	0xe8, 0xe1, 0xc4, 0xab, 0x49, 0x93, 0x54, 0xf1, 0x06, 0xae, 0xc2, 0xf7, 0x86, 0x19, 0x26, 0x0c,
	0x00, 0x21, 0xda, 0xde, 0x13, 0x46, 0x94, 0x7c, 0x76, 0xce, 0x50, 0xfb, 0xc9, 0xa0, 0xda, 0xae,
	0xda, 0xbc, 0xc9, 0x13, 0x88, 0xf7, 0x60, 0x68, 0x98, 0x14, 0x6b, 0x58, 0x77, 0xd0, 0x85, 0xe4,
	0x0d, 0x5c, 0x2c, 0xf8, 0x3a, 0xbb, 0xa7, 0xd5, 0x1b, 0x75, 0xd6, 0xb0, 0x8e, 0xeb, 0xdf, 0x83,
	0xac, 0xe1, 0x9c, 0x70, 0x31, 0x49, 0x3c, 0x24, 0xf1, 0x87, 0xb7, 0xc4, 0x2d, 0x7c, 0x98, 0xd4,
	0x31, 0x27, 0x66, 0x92, 0x0a, 0xaa, 0x11, 0xad, 0x7d, 0x76, 0x93, 0xe2, 0xc9, 0xb2, 0x63, 0x56,
	0x5c, 0x34, 0x2e, 0x2a, 0xd1, 0xb0, 0xee, 0x52, 0xef, 0x28, 0xaf, 0x81, 0x93, 0x89, 0xb4, 0xaa,
	0x2a, 0x39, 0x32, 0xb4, 0x5f, 0x7b, 0xff, 0xe0, 0xff, 0x79, 0xff, 0x13, 0x88, 0xf0, 0x31, 0x3e,
	0x12, 0x95, 0x0f, 0x14, 0xba, 0xbf, 0x03, 0x31, 0x58, 0x8d, 0x68, 0xe5, 0x2d, 0x1c, 0xd3, 0x7b,
	0xd5, 0xa7, 0x6d, 0xfa, 0xb2, 0x4b, 0xfd, 0x87, 0x83, 0x7f, 0x11, 0x69, 0xc1, 0x87, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xab, 0x32, 0x70, 0xdc, 0xcf, 0x01, 0x00, 0x00,
}
