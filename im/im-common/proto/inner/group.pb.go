// Code generated by protoc-gen-go. DO NOT EDIT.
// source: group.proto

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

type GroupReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LoginType            int32    `protobuf:"varint,2,opt,name=loginType,proto3" json:"loginType,omitempty"`
	RoleType             int32    `protobuf:"varint,3,opt,name=roleType,proto3" json:"roleType,omitempty"`
	GroupId              int64    `protobuf:"varint,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Command              int32    `protobuf:"varint,5,opt,name=command,proto3" json:"command,omitempty"`
	Content              []byte   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Retry                int32    `protobuf:"varint,7,opt,name=retry,proto3" json:"retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupReq) Reset()         { *m = GroupReq{} }
func (m *GroupReq) String() string { return proto.CompactTextString(m) }
func (*GroupReq) ProtoMessage()    {}
func (*GroupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e10f4c9b19ad8eee, []int{0}
}

func (m *GroupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupReq.Unmarshal(m, b)
}
func (m *GroupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupReq.Marshal(b, m, deterministic)
}
func (m *GroupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupReq.Merge(m, src)
}
func (m *GroupReq) XXX_Size() int {
	return xxx_messageInfo_GroupReq.Size(m)
}
func (m *GroupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupReq.DiscardUnknown(m)
}

var xxx_messageInfo_GroupReq proto.InternalMessageInfo

func (m *GroupReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GroupReq) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *GroupReq) GetRoleType() int32 {
	if m != nil {
		return m.RoleType
	}
	return 0
}

func (m *GroupReq) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupReq) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *GroupReq) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *GroupReq) GetRetry() int32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

type GroupRsp struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LoginType            int32    `protobuf:"varint,2,opt,name=loginType,proto3" json:"loginType,omitempty"`
	RoleType             int32    `protobuf:"varint,3,opt,name=roleType,proto3" json:"roleType,omitempty"`
	GroupId              int64    `protobuf:"varint,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Command              int32    `protobuf:"varint,5,opt,name=command,proto3" json:"command,omitempty"`
	Content              []byte   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	SvcErr               int32    `protobuf:"varint,7,opt,name=svcErr,proto3" json:"svcErr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupRsp) Reset()         { *m = GroupRsp{} }
func (m *GroupRsp) String() string { return proto.CompactTextString(m) }
func (*GroupRsp) ProtoMessage()    {}
func (*GroupRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e10f4c9b19ad8eee, []int{1}
}

func (m *GroupRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupRsp.Unmarshal(m, b)
}
func (m *GroupRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupRsp.Marshal(b, m, deterministic)
}
func (m *GroupRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupRsp.Merge(m, src)
}
func (m *GroupRsp) XXX_Size() int {
	return xxx_messageInfo_GroupRsp.Size(m)
}
func (m *GroupRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GroupRsp proto.InternalMessageInfo

func (m *GroupRsp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GroupRsp) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *GroupRsp) GetRoleType() int32 {
	if m != nil {
		return m.RoleType
	}
	return 0
}

func (m *GroupRsp) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupRsp) GetCommand() int32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *GroupRsp) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *GroupRsp) GetSvcErr() int32 {
	if m != nil {
		return m.SvcErr
	}
	return 0
}

func init() {
	proto.RegisterType((*GroupReq)(nil), "inner.groupReq")
	proto.RegisterType((*GroupRsp)(nil), "inner.groupRsp")
}

func init() { proto.RegisterFile("group.proto", fileDescriptor_e10f4c9b19ad8eee) }

var fileDescriptor_e10f4c9b19ad8eee = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x89, 0x63, 0xd2, 0xf1, 0x29, 0x08, 0x41, 0x86, 0x30, 0xb8, 0x28, 0xb3, 0xaa, 0x9b,
	0x2e, 0xd4, 0x5f, 0x10, 0xe9, 0x36, 0xf8, 0x03, 0xda, 0x86, 0x52, 0x68, 0x93, 0xf8, 0x92, 0x0a,
	0xfd, 0x3b, 0xf1, 0xcb, 0xa4, 0xaf, 0x4d, 0xc5, 0x3f, 0x98, 0xdd, 0x3b, 0xf7, 0xf2, 0x2e, 0x5c,
	0x2e, 0x5c, 0xb7, 0xe8, 0x46, 0x5f, 0x7a, 0x74, 0xd1, 0x49, 0xde, 0x59, 0x6b, 0xf0, 0xf4, 0xcd,
	0x60, 0x4f, 0xb2, 0x36, 0x9f, 0xf2, 0x00, 0x62, 0x0c, 0x06, 0xab, 0x46, 0xb1, 0x9c, 0x15, 0x3b,
	0xbd, 0x92, 0xbc, 0x87, 0xab, 0xde, 0xb5, 0x9d, 0x7d, 0x9b, 0xbc, 0x51, 0x17, 0x39, 0x2b, 0xb8,
	0xfe, 0x13, 0xe4, 0x11, 0xf6, 0xe8, 0x7a, 0x43, 0xe6, 0x8e, 0xcc, 0x8d, 0xa5, 0x82, 0x8c, 0xd2,
	0xab, 0x46, 0x5d, 0x52, 0x64, 0xc2, 0xd9, 0xa9, 0xdd, 0x30, 0xbc, 0xdb, 0x46, 0x71, 0x7a, 0x4a,
	0xb8, 0x38, 0x36, 0x1a, 0x1b, 0x95, 0xc8, 0x59, 0x71, 0xa3, 0x13, 0xca, 0x3b, 0xe0, 0x68, 0x22,
	0x4e, 0x2a, 0xa3, 0x8f, 0x05, 0x4e, 0x3f, 0x5b, 0x85, 0xe0, 0xcf, 0xbe, 0xc2, 0x01, 0x44, 0xf8,
	0xaa, 0x5f, 0x10, 0xd7, 0x0e, 0x2b, 0x3d, 0x3e, 0x43, 0x56, 0x0d, 0xaf, 0x73, 0xb0, 0x7c, 0x00,
	0xbe, 0x1c, 0xb7, 0x25, 0x6d, 0x54, 0xa6, 0x7d, 0x8e, 0xff, 0x85, 0xe0, 0x3f, 0x04, 0x6d, 0xf9,
	0xf4, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb3, 0x58, 0x69, 0xda, 0x01, 0x00, 0x00,
}
