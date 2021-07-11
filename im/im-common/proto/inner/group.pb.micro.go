// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: group.proto

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

// Api Endpoints for ImGroup service

func NewImGroupEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ImGroup service

type ImGroupService interface {
	Group(ctx context.Context, in *GroupReq, opts ...client.CallOption) (*GroupRsp, error)
}

type imGroupService struct {
	c    client.Client
	name string
}

func NewImGroupService(name string, c client.Client) ImGroupService {
	return &imGroupService{
		c:    c,
		name: name,
	}
}

func (c *imGroupService) Group(ctx context.Context, in *GroupReq, opts ...client.CallOption) (*GroupRsp, error) {
	req := c.c.NewRequest(c.name, "ImGroup.Group", in)
	out := new(GroupRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImGroup service

type ImGroupHandler interface {
	Group(context.Context, *GroupReq, *GroupRsp) error
}

func RegisterImGroupHandler(s server.Server, hdlr ImGroupHandler, opts ...server.HandlerOption) error {
	type imGroup interface {
		Group(ctx context.Context, in *GroupReq, out *GroupRsp) error
	}
	type ImGroup struct {
		imGroup
	}
	h := &imGroupHandler{hdlr}
	return s.Handle(s.NewHandler(&ImGroup{h}, opts...))
}

type imGroupHandler struct {
	ImGroupHandler
}

func (h *imGroupHandler) Group(ctx context.Context, in *GroupReq, out *GroupRsp) error {
	return h.ImGroupHandler.Group(ctx, in, out)
}