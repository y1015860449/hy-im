// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv.err.proto

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

type SrvErr int32

const (
	SrvErr_srv_err_unknown           SrvErr = 0
	SrvErr_srv_err_success           SrvErr = 1
	SrvErr_srv_err_redis             SrvErr = 2
	SrvErr_srv_err_mongo             SrvErr = 3
	SrvErr_srv_err_mysql             SrvErr = 4
	SrvErr_srv_err_param             SrvErr = 5
	SrvErr_srv_err_account_already   SrvErr = 6
	SrvErr_srv_err_account_not_exist SrvErr = 7
	SrvErr_srv_err_user_auth         SrvErr = 8
)

var SrvErr_name = map[int32]string{
	0: "srv_err_unknown",
	1: "srv_err_success",
	2: "srv_err_redis",
	3: "srv_err_mongo",
	4: "srv_err_mysql",
	5: "srv_err_param",
	6: "srv_err_account_already",
	7: "srv_err_account_not_exist",
	8: "srv_err_user_auth",
}

var SrvErr_value = map[string]int32{
	"srv_err_unknown":           0,
	"srv_err_success":           1,
	"srv_err_redis":             2,
	"srv_err_mongo":             3,
	"srv_err_mysql":             4,
	"srv_err_param":             5,
	"srv_err_account_already":   6,
	"srv_err_account_not_exist": 7,
	"srv_err_user_auth":         8,
}

func (x SrvErr) String() string {
	return proto.EnumName(SrvErr_name, int32(x))
}

func (SrvErr) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_81745e75c9de74d0, []int{0}
}

func init() {
	proto.RegisterEnum("inner.SrvErr", SrvErr_name, SrvErr_value)
}

func init() { proto.RegisterFile("srv.err.proto", fileDescriptor_81745e75c9de74d0) }

var fileDescriptor_81745e75c9de74d0 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0x61, 0x8a, 0xc2, 0x30,
	0x10, 0x05, 0xe0, 0xed, 0xee, 0xb6, 0xbb, 0x04, 0xc4, 0x69, 0x44, 0x44, 0xc4, 0x0b, 0xf8, 0xa3,
	0x7f, 0x3c, 0x83, 0x27, 0xf0, 0x00, 0x21, 0xa6, 0x83, 0x16, 0xdb, 0x49, 0x9d, 0x49, 0xaa, 0x3d,
	0xa6, 0x37, 0x12, 0x0a, 0x4a, 0xf5, 0xef, 0xf7, 0xe0, 0xf1, 0x9e, 0x9a, 0x08, 0x77, 0x05, 0x32,
	0x17, 0x2d, 0xfb, 0xe0, 0x75, 0x5a, 0x11, 0x21, 0x6f, 0xee, 0x89, 0xca, 0xf6, 0xdc, 0xed, 0x98,
	0xf5, 0x4c, 0x4d, 0x85, 0x3b, 0x83, 0xcc, 0x26, 0xd2, 0x99, 0xfc, 0x95, 0xe0, 0x6b, 0x8c, 0x12,
	0x9d, 0x43, 0x11, 0x48, 0x74, 0x3e, 0x94, 0x0d, 0xc8, 0x58, 0x56, 0x02, 0xdf, 0x63, 0x6a, 0x3c,
	0x1d, 0x3d, 0xfc, 0xbc, 0x51, 0x2f, 0x97, 0x1a, 0x7e, 0xc7, 0xd4, 0x5a, 0xb6, 0x0d, 0xa4, 0x7a,
	0xa5, 0x16, 0x4f, 0xb2, 0xce, 0xf9, 0x48, 0xc1, 0xd8, 0x9a, 0xd1, 0x96, 0x3d, 0x64, 0x7a, 0xad,
	0x96, 0x9f, 0x21, 0xf9, 0x60, 0xf0, 0x56, 0x49, 0x80, 0x3f, 0x3d, 0x57, 0xf9, 0x6b, 0xb1, 0x20,
	0x1b, 0x1b, 0xc3, 0x09, 0xfe, 0x0f, 0xd9, 0xf0, 0x70, 0xfb, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x73, 0x68, 0xba, 0xf2, 0x00, 0x00, 0x00,
}
