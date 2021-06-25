// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: room.proto

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

// Api Endpoints for ImRoom service

func NewImRoomEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ImRoom service

type ImRoomService interface {
	Room(ctx context.Context, in *RoomReq, opts ...client.CallOption) (*RoomRsp, error)
}

type imRoomService struct {
	c    client.Client
	name string
}

func NewImRoomService(name string, c client.Client) ImRoomService {
	return &imRoomService{
		c:    c,
		name: name,
	}
}

func (c *imRoomService) Room(ctx context.Context, in *RoomReq, opts ...client.CallOption) (*RoomRsp, error) {
	req := c.c.NewRequest(c.name, "ImRoom.Room", in)
	out := new(RoomRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImRoom service

type ImRoomHandler interface {
	Room(context.Context, *RoomReq, *RoomRsp) error
}

func RegisterImRoomHandler(s server.Server, hdlr ImRoomHandler, opts ...server.HandlerOption) error {
	type imRoom interface {
		Room(ctx context.Context, in *RoomReq, out *RoomRsp) error
	}
	type ImRoom struct {
		imRoom
	}
	h := &imRoomHandler{hdlr}
	return s.Handle(s.NewHandler(&ImRoom{h}, opts...))
}

type imRoomHandler struct {
	ImRoomHandler
}

func (h *imRoomHandler) Room(ctx context.Context, in *RoomReq, out *RoomRsp) error {
	return h.ImRoomHandler.Room(ctx, in, out)
}
