package common

type LoginInfo struct {
	UserId int64
	LoginType int32			// 登录类型	0：APP 1：PC
	RoleType int32			// 角色类型  0；游客 1：用户
}

type ConnectionCtx struct {
	RoomId int64
}
