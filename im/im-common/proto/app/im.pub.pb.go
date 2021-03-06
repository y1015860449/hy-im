// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im.pub.proto

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

type ImCmd int32

const (
	ImCmd_cmd_unknown               ImCmd = 0
	ImCmd_cmd_heartbeat             ImCmd = 1
	ImCmd_cmd_heartbeat_ack         ImCmd = 2
	ImCmd_cmd_login                 ImCmd = 2049
	ImCmd_cmd_login_ack             ImCmd = 2050
	ImCmd_cmd_logout                ImCmd = 2051
	ImCmd_cmd_logout_ack            ImCmd = 2052
	ImCmd_cmd_kick_out              ImCmd = 2054
	ImCmd_cmd_group_msg             ImCmd = 4097
	ImCmd_cmd_group_msg_ack         ImCmd = 4098
	ImCmd_cmd_group_msg_deliver     ImCmd = 4100
	ImCmd_cmd_group_msg_deliver_ack ImCmd = 4099
	// 打开
	ImCmd_cmd_group_open             ImCmd = 4101
	ImCmd_cmd_group_open_ack         ImCmd = 4102
	ImCmd_cmd_group_open_deliver     ImCmd = 4104
	ImCmd_cmd_group_open_deliver_ack ImCmd = 4103
	// 加入
	ImCmd_cmd_group_join             ImCmd = 4105
	ImCmd_cmd_group_join_ack         ImCmd = 4106
	ImCmd_cmd_group_join_deliver     ImCmd = 4108
	ImCmd_cmd_group_join_deliver_ack ImCmd = 4107
	// 移除
	ImCmd_cmd_group_remove             ImCmd = 4109
	ImCmd_cmd_group_remove_ack         ImCmd = 4110
	ImCmd_cmd_group_remove_deliver     ImCmd = 4112
	ImCmd_cmd_group_remove_deliver_ack ImCmd = 4111
	// 退出
	ImCmd_cmd_group_quit             ImCmd = 4113
	ImCmd_cmd_group_quit_ack         ImCmd = 4114
	ImCmd_cmd_group_quit_deliver     ImCmd = 4116
	ImCmd_cmd_group_quit_deliver_ack ImCmd = 4115
	// 关闭
	ImCmd_cmd_group_close             ImCmd = 4117
	ImCmd_cmd_group_close_ack         ImCmd = 4118
	ImCmd_cmd_group_close_deliver     ImCmd = 4120
	ImCmd_cmd_group_close_deliver_ack ImCmd = 4119
	// 群组消息撤回
	ImCmd_cmd_group_msg_cancel             ImCmd = 4121
	ImCmd_cmd_group_msg_cancel_ack         ImCmd = 4122
	ImCmd_cmd_group_msg_cancel_deliver     ImCmd = 4124
	ImCmd_cmd_group_msg_cancel_deliver_ack ImCmd = 4123
	ImCmd_cmd_p2p_msg                      ImCmd = 8193
	ImCmd_cmd_p2p_msg_ack                  ImCmd = 8194
	ImCmd_cmd_p2p_msg_deliver              ImCmd = 8196
	ImCmd_cmd_p2p_msg_deliver_ack          ImCmd = 8195
	ImCmd_cmd_p2p_msg_notify               ImCmd = 8198
	ImCmd_cmd_p2p_msg_notify_ack           ImCmd = 8197
	ImCmd_cmd_p2p_msg_read                 ImCmd = 8199
	ImCmd_cmd_p2p_msg_read_ack             ImCmd = 8200
	ImCmd_cmd_p2p_msg_read_deliver         ImCmd = 8202
	ImCmd_cmd_p2p_msg_read_deliver_ack     ImCmd = 8201
	ImCmd_cmd_p2p_msg_cancel               ImCmd = 8203
	ImCmd_cmd_p2p_msg_cancel_ack           ImCmd = 8204
	ImCmd_cmd_p2p_msg_cancel_deliver       ImCmd = 8206
	ImCmd_cmd_p2p_msg_cancel_deliver_ack   ImCmd = 8205
)

var ImCmd_name = map[int32]string{
	0:    "cmd_unknown",
	1:    "cmd_heartbeat",
	2:    "cmd_heartbeat_ack",
	2049: "cmd_login",
	2050: "cmd_login_ack",
	2051: "cmd_logout",
	2052: "cmd_logout_ack",
	2054: "cmd_kick_out",
	4097: "cmd_group_msg",
	4098: "cmd_group_msg_ack",
	4100: "cmd_group_msg_deliver",
	4099: "cmd_group_msg_deliver_ack",
	4101: "cmd_group_open",
	4102: "cmd_group_open_ack",
	4104: "cmd_group_open_deliver",
	4103: "cmd_group_open_deliver_ack",
	4105: "cmd_group_join",
	4106: "cmd_group_join_ack",
	4108: "cmd_group_join_deliver",
	4107: "cmd_group_join_deliver_ack",
	4109: "cmd_group_remove",
	4110: "cmd_group_remove_ack",
	4112: "cmd_group_remove_deliver",
	4111: "cmd_group_remove_deliver_ack",
	4113: "cmd_group_quit",
	4114: "cmd_group_quit_ack",
	4116: "cmd_group_quit_deliver",
	4115: "cmd_group_quit_deliver_ack",
	4117: "cmd_group_close",
	4118: "cmd_group_close_ack",
	4120: "cmd_group_close_deliver",
	4119: "cmd_group_close_deliver_ack",
	4121: "cmd_group_msg_cancel",
	4122: "cmd_group_msg_cancel_ack",
	4124: "cmd_group_msg_cancel_deliver",
	4123: "cmd_group_msg_cancel_deliver_ack",
	8193: "cmd_p2p_msg",
	8194: "cmd_p2p_msg_ack",
	8196: "cmd_p2p_msg_deliver",
	8195: "cmd_p2p_msg_deliver_ack",
	8198: "cmd_p2p_msg_notify",
	8197: "cmd_p2p_msg_notify_ack",
	8199: "cmd_p2p_msg_read",
	8200: "cmd_p2p_msg_read_ack",
	8202: "cmd_p2p_msg_read_deliver",
	8201: "cmd_p2p_msg_read_deliver_ack",
	8203: "cmd_p2p_msg_cancel",
	8204: "cmd_p2p_msg_cancel_ack",
	8206: "cmd_p2p_msg_cancel_deliver",
	8205: "cmd_p2p_msg_cancel_deliver_ack",
}

var ImCmd_value = map[string]int32{
	"cmd_unknown":                      0,
	"cmd_heartbeat":                    1,
	"cmd_heartbeat_ack":                2,
	"cmd_login":                        2049,
	"cmd_login_ack":                    2050,
	"cmd_logout":                       2051,
	"cmd_logout_ack":                   2052,
	"cmd_kick_out":                     2054,
	"cmd_group_msg":                    4097,
	"cmd_group_msg_ack":                4098,
	"cmd_group_msg_deliver":            4100,
	"cmd_group_msg_deliver_ack":        4099,
	"cmd_group_open":                   4101,
	"cmd_group_open_ack":               4102,
	"cmd_group_open_deliver":           4104,
	"cmd_group_open_deliver_ack":       4103,
	"cmd_group_join":                   4105,
	"cmd_group_join_ack":               4106,
	"cmd_group_join_deliver":           4108,
	"cmd_group_join_deliver_ack":       4107,
	"cmd_group_remove":                 4109,
	"cmd_group_remove_ack":             4110,
	"cmd_group_remove_deliver":         4112,
	"cmd_group_remove_deliver_ack":     4111,
	"cmd_group_quit":                   4113,
	"cmd_group_quit_ack":               4114,
	"cmd_group_quit_deliver":           4116,
	"cmd_group_quit_deliver_ack":       4115,
	"cmd_group_close":                  4117,
	"cmd_group_close_ack":              4118,
	"cmd_group_close_deliver":          4120,
	"cmd_group_close_deliver_ack":      4119,
	"cmd_group_msg_cancel":             4121,
	"cmd_group_msg_cancel_ack":         4122,
	"cmd_group_msg_cancel_deliver":     4124,
	"cmd_group_msg_cancel_deliver_ack": 4123,
	"cmd_p2p_msg":                      8193,
	"cmd_p2p_msg_ack":                  8194,
	"cmd_p2p_msg_deliver":              8196,
	"cmd_p2p_msg_deliver_ack":          8195,
	"cmd_p2p_msg_notify":               8198,
	"cmd_p2p_msg_notify_ack":           8197,
	"cmd_p2p_msg_read":                 8199,
	"cmd_p2p_msg_read_ack":             8200,
	"cmd_p2p_msg_read_deliver":         8202,
	"cmd_p2p_msg_read_deliver_ack":     8201,
	"cmd_p2p_msg_cancel":               8203,
	"cmd_p2p_msg_cancel_ack":           8204,
	"cmd_p2p_msg_cancel_deliver":       8206,
	"cmd_p2p_msg_cancel_deliver_ack":   8205,
}

func (x ImCmd) String() string {
	return proto.EnumName(ImCmd_name, int32(x))
}

func (ImCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_419a8c3136d42e2a, []int{0}
}

type ImErrCode int32

const (
	ImErrCode_err_unknown          ImErrCode = 0
	ImErrCode_err_success          ImErrCode = 61440
	ImErrCode_err_server_except    ImErrCode = 61441
	ImErrCode_err_param_except     ImErrCode = 61442
	ImErrCode_err_msg_empty        ImErrCode = 61443
	ImErrCode_err_msg_except       ImErrCode = 61444
	ImErrCode_err_user_written_off ImErrCode = 61445
	ImErrCode_err_user_freeze      ImErrCode = 61446
	ImErrCode_err_user_except      ImErrCode = 61447
	ImErrCode_err_user_visitor     ImErrCode = 61448
	ImErrCode_err_login_auth       ImErrCode = 61697
	ImErrCode_err_login_kick_out   ImErrCode = 61698
)

var ImErrCode_name = map[int32]string{
	0:     "err_unknown",
	61440: "err_success",
	61441: "err_server_except",
	61442: "err_param_except",
	61443: "err_msg_empty",
	61444: "err_msg_except",
	61445: "err_user_written_off",
	61446: "err_user_freeze",
	61447: "err_user_except",
	61448: "err_user_visitor",
	61697: "err_login_auth",
	61698: "err_login_kick_out",
}

var ImErrCode_value = map[string]int32{
	"err_unknown":          0,
	"err_success":          61440,
	"err_server_except":    61441,
	"err_param_except":     61442,
	"err_msg_empty":        61443,
	"err_msg_except":       61444,
	"err_user_written_off": 61445,
	"err_user_freeze":      61446,
	"err_user_except":      61447,
	"err_user_visitor":     61448,
	"err_login_auth":       61697,
	"err_login_kick_out":   61698,
}

func (x ImErrCode) String() string {
	return proto.EnumName(ImErrCode_name, int32(x))
}

func (ImErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_419a8c3136d42e2a, []int{1}
}

func init() {
	proto.RegisterEnum("im.app.ImCmd", ImCmd_name, ImCmd_value)
	proto.RegisterEnum("im.app.ImErrCode", ImErrCode_name, ImErrCode_value)
}

func init() { proto.RegisterFile("im.pub.proto", fileDescriptor_419a8c3136d42e2a) }

var fileDescriptor_419a8c3136d42e2a = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0xcb, 0x52, 0xdb, 0x4c,
	0x10, 0x85, 0xff, 0x1f, 0x55, 0xa8, 0xa2, 0x03, 0xb8, 0xdd, 0xd8, 0xdc, 0x43, 0x26, 0x95, 0xca,
	0x86, 0x05, 0x8b, 0xe4, 0x05, 0x54, 0x45, 0x65, 0xc1, 0x53, 0xb8, 0x84, 0x3c, 0x06, 0xc5, 0x96,
	0x46, 0x19, 0x4b, 0x10, 0xb2, 0x8a, 0xb9, 0x98, 0x3b, 0xb9, 0xdf, 0x93, 0x07, 0xc9, 0xdb, 0xb8,
	0x78, 0x92, 0x94, 0x5a, 0x1e, 0xcb, 0x1e, 0x20, 0x4b, 0x7f, 0xe7, 0xf4, 0x9c, 0x9e, 0x1e, 0xcd,
	0x18, 0x26, 0x83, 0x70, 0x2d, 0x4e, 0x37, 0xd7, 0x62, 0xad, 0x12, 0x45, 0xe3, 0x41, 0xb8, 0xe6,
	0xc5, 0xf1, 0xea, 0x1f, 0x80, 0x7b, 0x1b, 0xe1, 0x7a, 0x58, 0xa7, 0x12, 0xdc, 0xf7, 0xc3, 0x7a,
	0x2d, 0x8d, 0x9a, 0x91, 0xda, 0x8d, 0xf0, 0x3f, 0x2a, 0xc3, 0x54, 0x06, 0xb6, 0xa5, 0xa7, 0x93,
	0x4d, 0xe9, 0x25, 0xf8, 0x3f, 0x55, 0xa1, 0x3c, 0x82, 0x6a, 0x9e, 0xdf, 0xc4, 0x31, 0x9a, 0x86,
	0x89, 0x0c, 0xb7, 0xd4, 0x56, 0x10, 0x61, 0x07, 0x89, 0xf2, 0x4a, 0xfe, 0xcd, 0x96, 0x7d, 0xa4,
	0x12, 0x40, 0x9f, 0xa9, 0x34, 0xc1, 0x03, 0xa4, 0x19, 0x98, 0x2e, 0x00, 0xbb, 0x0e, 0x91, 0xca,
	0x30, 0x99, 0xc1, 0x66, 0xe0, 0x37, 0x6b, 0x99, 0xaf, 0x3b, 0x58, 0x6c, 0x4b, 0xab, 0x34, 0xae,
	0x85, 0xed, 0x2d, 0xec, 0x08, 0x9a, 0xcd, 0xfb, 0x18, 0xb0, 0x3c, 0x44, 0xd0, 0x22, 0x54, 0x47,
	0x79, 0x5d, 0xb6, 0x82, 0x1d, 0xa9, 0xf1, 0x50, 0xd0, 0x0a, 0x2c, 0xdc, 0xaa, 0x71, 0xed, 0x81,
	0x30, 0xfd, 0xe4, 0xba, 0x8a, 0x65, 0x84, 0x47, 0x82, 0xe6, 0x80, 0x46, 0x21, 0xbb, 0xbb, 0x82,
	0x96, 0x60, 0xd6, 0x12, 0x4c, 0xd4, 0x89, 0xa0, 0x87, 0xb0, 0x78, 0xbb, 0xc8, 0xd5, 0xc7, 0x56,
	0xd6, 0x0b, 0x15, 0x44, 0x78, 0x6a, 0x65, 0x65, 0x90, 0xdd, 0x67, 0x56, 0x16, 0x0b, 0x26, 0xeb,
	0xc2, 0xca, 0x1a, 0x16, 0xb9, 0xfa, 0x5c, 0x50, 0x15, 0xb0, 0x30, 0x68, 0x19, 0xaa, 0x1d, 0x89,
	0x97, 0x82, 0x16, 0xa0, 0x62, 0x63, 0xae, 0xb8, 0x12, 0xf4, 0x00, 0xe6, 0x6f, 0x48, 0x26, 0xf1,
	0x9d, 0xa0, 0x47, 0xb0, 0x7c, 0x97, 0xcc, 0x2b, 0xbc, 0xb5, 0xf6, 0xf7, 0x32, 0x0d, 0x12, 0x7c,
	0x6f, 0xed, 0x2f, 0x83, 0xec, 0xfe, 0x60, 0xed, 0x8f, 0x05, 0x93, 0xf6, 0xc9, 0xda, 0xdf, 0xb0,
	0xc8, 0xd5, 0x1f, 0x05, 0x55, 0xa0, 0x54, 0x18, 0xfc, 0x96, 0x6a, 0x4b, 0xfc, 0x2c, 0x68, 0x1e,
	0x66, 0x2c, 0xca, 0xfe, 0x2f, 0x82, 0x96, 0x61, 0xce, 0x56, 0x4c, 0xdc, 0x37, 0x41, 0x02, 0x96,
	0xee, 0x50, 0xb9, 0xfe, 0xab, 0x35, 0xb8, 0xec, 0x3b, 0xf2, 0xbd, 0xc8, 0x97, 0x2d, 0xfc, 0x6e,
	0x0d, 0xae, 0x90, 0xb8, 0xf2, 0x87, 0x35, 0xb8, 0x21, 0xd9, 0xc4, 0xff, 0x12, 0xf4, 0x04, 0xc4,
	0xbf, 0x2c, 0xbc, 0xd2, 0x4f, 0x41, 0x98, 0xdf, 0xd5, 0xf8, 0x69, 0xff, 0x46, 0xb8, 0x66, 0x0a,
	0x7d, 0x92, 0xdf, 0x07, 0xd7, 0x4c, 0xc1, 0xd0, 0xc1, 0x6d, 0x70, 0xcd, 0x14, 0x2c, 0x25, 0xbf,
	0x0b, 0xae, 0x39, 0x2a, 0xa3, 0x46, 0x2a, 0x09, 0x1a, 0x7b, 0xd8, 0x75, 0xcd, 0x51, 0x8d, 0x0a,
	0x5c, 0x75, 0xe4, 0x9a, 0x2f, 0xcd, 0x88, 0x5a, 0x7a, 0x75, 0x3c, 0x76, 0xcd, 0xc0, 0x86, 0x31,
	0x57, 0x9c, 0xb8, 0x66, 0x60, 0x23, 0x92, 0x69, 0xf2, 0xcc, 0x35, 0x03, 0xbb, 0x4d, 0xe6, 0x15,
	0x4e, 0x6f, 0x74, 0xda, 0x3f, 0x8b, 0xf3, 0x1b, 0x9d, 0x0e, 0x9d, 0xc4, 0x85, 0x6b, 0x3e, 0x2a,
	0x4b, 0x34, 0xc9, 0x57, 0x2e, 0x3d, 0x86, 0x95, 0xbb, 0x0d, 0xbc, 0xca, 0xa5, 0xbb, 0xfa, 0x7b,
	0x0c, 0x26, 0x36, 0xc2, 0xe7, 0x5a, 0xaf, 0xab, 0xba, 0xcc, 0xde, 0x4f, 0xa9, 0xf5, 0xc8, 0xfb,
	0xc9, 0xa0, 0x9d, 0xfa, 0xbe, 0x6c, 0xb7, 0xf1, 0x4d, 0xcf, 0xa1, 0x39, 0x28, 0x33, 0x92, 0x3a,
	0x5b, 0x46, 0xbe, 0xf2, 0x65, 0x9c, 0x60, 0xa7, 0xe7, 0xd0, 0x2c, 0x60, 0x26, 0xc4, 0x9e, 0xf6,
	0x42, 0xc3, 0xf7, 0x7b, 0x0e, 0xcd, 0xc0, 0x54, 0xc6, 0xb3, 0x1e, 0x64, 0x18, 0x27, 0x7b, 0x78,
	0xd0, 0x73, 0xa8, 0x02, 0xd3, 0x03, 0x98, 0x5b, 0x0f, 0x7b, 0x0e, 0x2d, 0x42, 0x85, 0xf3, 0xdb,
	0x52, 0xd7, 0x76, 0x75, 0x90, 0x24, 0x32, 0xaa, 0xa9, 0x46, 0x03, 0x8f, 0x7a, 0x0e, 0x55, 0xa1,
	0x34, 0xd0, 0x1a, 0x5a, 0xca, 0xd7, 0x12, 0xbb, 0x16, 0xee, 0xaf, 0x74, 0x5c, 0x34, 0xc3, 0x78,
	0x27, 0x68, 0x07, 0x89, 0xd2, 0x78, 0x52, 0xe4, 0xf6, 0x9f, 0xf5, 0x34, 0xd9, 0xc6, 0xce, 0xb5,
	0x43, 0xf3, 0x40, 0x05, 0x1d, 0x3c, 0xdc, 0xfb, 0xd7, 0xce, 0xe6, 0x38, 0xff, 0xd5, 0x3c, 0xfb,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x65, 0x13, 0x98, 0x7a, 0x06, 0x00, 0x00,
}
