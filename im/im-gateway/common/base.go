package common

type LoginInfo struct {
	UserId    int64
	LoginType int32 // 登录类型	0：APP 1：PC
}

type ConnectionCtx struct {
	RoleType int32 // 角色类型  0；游客 1：用户
	RoomId   int64
}
