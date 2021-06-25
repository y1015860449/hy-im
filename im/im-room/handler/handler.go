package handler

import (
	"context"
	innerPt "hy-im/im/im-common/proto/inner"
	"hy-im/im/im-room/dao/cache"
)

type Handler struct {
	CacheDao cache.CacheDao
}

func (h *Handler) Room(ctx context.Context, req *innerPt.RoomReq, rsp *innerPt.RoomRsp) error {

	panic("implement me")
}






