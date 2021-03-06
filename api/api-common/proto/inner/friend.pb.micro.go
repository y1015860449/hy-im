// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: friend.proto

package inner

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ApiFriend service

func NewApiFriendEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ApiFriend service

type ApiFriendService interface {
	BecomeFriend(ctx context.Context, in *BecomeFriendReq, opts ...client.CallOption) (*BecomeFriendRsp, error)
	DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...client.CallOption) (*DeleteFriendRsp, error)
	BlackFriend(ctx context.Context, in *BlackFriendReq, opts ...client.CallOption) (*BlackFriendRsp, error)
	GetFriendList(ctx context.Context, in *GetFriendListReq, opts ...client.CallOption) (*GetFriendListRsp, error)
	GetFriendInfo(ctx context.Context, in *GetFriendInfoReq, opts ...client.CallOption) (*GetFriendInfoRsp, error)
}

type apiFriendService struct {
	c    client.Client
	name string
}

func NewApiFriendService(name string, c client.Client) ApiFriendService {
	return &apiFriendService{
		c:    c,
		name: name,
	}
}

func (c *apiFriendService) BecomeFriend(ctx context.Context, in *BecomeFriendReq, opts ...client.CallOption) (*BecomeFriendRsp, error) {
	req := c.c.NewRequest(c.name, "ApiFriend.BecomeFriend", in)
	out := new(BecomeFriendRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiFriendService) DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...client.CallOption) (*DeleteFriendRsp, error) {
	req := c.c.NewRequest(c.name, "ApiFriend.DeleteFriend", in)
	out := new(DeleteFriendRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiFriendService) BlackFriend(ctx context.Context, in *BlackFriendReq, opts ...client.CallOption) (*BlackFriendRsp, error) {
	req := c.c.NewRequest(c.name, "ApiFriend.BlackFriend", in)
	out := new(BlackFriendRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiFriendService) GetFriendList(ctx context.Context, in *GetFriendListReq, opts ...client.CallOption) (*GetFriendListRsp, error) {
	req := c.c.NewRequest(c.name, "ApiFriend.GetFriendList", in)
	out := new(GetFriendListRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiFriendService) GetFriendInfo(ctx context.Context, in *GetFriendInfoReq, opts ...client.CallOption) (*GetFriendInfoRsp, error) {
	req := c.c.NewRequest(c.name, "ApiFriend.GetFriendInfo", in)
	out := new(GetFriendInfoRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApiFriend service

type ApiFriendHandler interface {
	BecomeFriend(context.Context, *BecomeFriendReq, *BecomeFriendRsp) error
	DeleteFriend(context.Context, *DeleteFriendReq, *DeleteFriendRsp) error
	BlackFriend(context.Context, *BlackFriendReq, *BlackFriendRsp) error
	GetFriendList(context.Context, *GetFriendListReq, *GetFriendListRsp) error
	GetFriendInfo(context.Context, *GetFriendInfoReq, *GetFriendInfoRsp) error
}

func RegisterApiFriendHandler(s server.Server, hdlr ApiFriendHandler, opts ...server.HandlerOption) error {
	type apiFriend interface {
		BecomeFriend(ctx context.Context, in *BecomeFriendReq, out *BecomeFriendRsp) error
		DeleteFriend(ctx context.Context, in *DeleteFriendReq, out *DeleteFriendRsp) error
		BlackFriend(ctx context.Context, in *BlackFriendReq, out *BlackFriendRsp) error
		GetFriendList(ctx context.Context, in *GetFriendListReq, out *GetFriendListRsp) error
		GetFriendInfo(ctx context.Context, in *GetFriendInfoReq, out *GetFriendInfoRsp) error
	}
	type ApiFriend struct {
		apiFriend
	}
	h := &apiFriendHandler{hdlr}
	return s.Handle(s.NewHandler(&ApiFriend{h}, opts...))
}

type apiFriendHandler struct {
	ApiFriendHandler
}

func (h *apiFriendHandler) BecomeFriend(ctx context.Context, in *BecomeFriendReq, out *BecomeFriendRsp) error {
	return h.ApiFriendHandler.BecomeFriend(ctx, in, out)
}

func (h *apiFriendHandler) DeleteFriend(ctx context.Context, in *DeleteFriendReq, out *DeleteFriendRsp) error {
	return h.ApiFriendHandler.DeleteFriend(ctx, in, out)
}

func (h *apiFriendHandler) BlackFriend(ctx context.Context, in *BlackFriendReq, out *BlackFriendRsp) error {
	return h.ApiFriendHandler.BlackFriend(ctx, in, out)
}

func (h *apiFriendHandler) GetFriendList(ctx context.Context, in *GetFriendListReq, out *GetFriendListRsp) error {
	return h.ApiFriendHandler.GetFriendList(ctx, in, out)
}

func (h *apiFriendHandler) GetFriendInfo(ctx context.Context, in *GetFriendInfoReq, out *GetFriendInfoRsp) error {
	return h.ApiFriendHandler.GetFriendInfo(ctx, in, out)
}
