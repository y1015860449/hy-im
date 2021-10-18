package handler

import (
	"context"
	innerPt "hy-im/api/api-common/proto/inner"
	"hy-im/api/api-friend/conf"
	"hy-im/api/api-friend/dao/db"
)

type Handler struct {
	friendDao db.ImFriendDao
	tokenConf *conf.Token
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

func (h *Handler) CheckToken(ctx context.Context, req *innerPt.CheckTokenReq, rsp *innerPt.CheckTokenRsp) error {
	panic("implement me")
}




