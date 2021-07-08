package handler

import (
	"context"
	innerPt "hy-im/api/api-common/proto/inner"
)

type Handler struct {
}

func (h *Handler) Register(ctx context.Context, req *innerPt.RegisterReq, rsp *innerPt.RegisterRsp) error {
	panic("implement me")
}

func (h *Handler) Login(ctx context.Context, req *innerPt.LoginReq, rsp *innerPt.LoginRsp) error {
	panic("implement me")
}

func (h *Handler) Logout(ctx context.Context, req *innerPt.LogoutReq, rsp *innerPt.LogoutRsp) error {
	panic("implement me")
}


