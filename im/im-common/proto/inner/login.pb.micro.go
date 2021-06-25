// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: login.proto

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

// Api Endpoints for ImLogin service

func NewImLoginEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ImLogin service

type ImLoginService interface {
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginRsp, error)
}

type imLoginService struct {
	c    client.Client
	name string
}

func NewImLoginService(name string, c client.Client) ImLoginService {
	return &imLoginService{
		c:    c,
		name: name,
	}
}

func (c *imLoginService) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginRsp, error) {
	req := c.c.NewRequest(c.name, "ImLogin.Login", in)
	out := new(LoginRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImLogin service

type ImLoginHandler interface {
	Login(context.Context, *LoginReq, *LoginRsp) error
}

func RegisterImLoginHandler(s server.Server, hdlr ImLoginHandler, opts ...server.HandlerOption) error {
	type imLogin interface {
		Login(ctx context.Context, in *LoginReq, out *LoginRsp) error
	}
	type ImLogin struct {
		imLogin
	}
	h := &imLoginHandler{hdlr}
	return s.Handle(s.NewHandler(&ImLogin{h}, opts...))
}

type imLoginHandler struct {
	ImLoginHandler
}

func (h *imLoginHandler) Login(ctx context.Context, in *LoginReq, out *LoginRsp) error {
	return h.ImLoginHandler.Login(ctx, in, out)
}