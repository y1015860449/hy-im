package handler

import (
	"context"
	innerPt "hy-im/api/api-common/proto/inner"
	"hy-im/api/api-group/dao/db/group"
)

type Handler struct {
	groupDao group.ImGroupDao
}

func (h *Handler) CreateGroup(ctx context.Context, req *innerPt.CreateGroupReq, rsp *innerPt.CreateGroupRsp) error {
	panic("implement me")
}

func (h *Handler) DismissGroup(ctx context.Context, req *innerPt.DismissGroupReq, rsp *innerPt.DismissGroupRsp) error {
	panic("implement me")
}

func (h *Handler) ChangeGroupMaster(ctx context.Context, req *innerPt.ChangeGroupMasterReq, rsp *innerPt.ChangeGroupMasterRsp) error {
	panic("implement me")
}

func (h *Handler) AddGroupMember(ctx context.Context, req *innerPt.AddGroupMemberReq, rsp *innerPt.AddGroupMemberRsp) error {
	panic("implement me")
}

func (h *Handler) RemoveGroupMember(ctx context.Context, req *innerPt.RemoveGroupMemberReq, rsp *innerPt.RemoveGroupMemberRsp) error {
	panic("implement me")
}

func (h *Handler) QuitGroupMember(ctx context.Context, req *innerPt.QuitGroupMemberReq, rsp *innerPt.QuitGroupMemberRsp) error {
	panic("implement me")
}

func (h *Handler) GetGroupInfo(ctx context.Context, req *innerPt.GetGroupInfoReq, rsp *innerPt.GetGroupInfoRsp) error {
	panic("implement me")
}

func (h *Handler) GetGroupList(ctx context.Context, req *innerPt.GetGroupListReq, rsp *innerPt.GetGroupListRsp) error {
	panic("implement me")
}

