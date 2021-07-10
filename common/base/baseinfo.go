package base

const (
	UserNormal     = 1 // 正常用户
	UserWrittenOff = 2 // 账号注销
	UserFreeze     = 3 // 账号冻结

	FriendNormal = 1 // 好友关系正常
	FriendDelete = 2 // 好友关系删除
	FriendBlack  = 3 // 好友关系拉黑

	GroupNormal  = 1 // 群组正常
	GroupDismiss = 1 // 群组解散
	GroupFreeze  = 1 // 群组冻结

	GroupTypeNormal     = 1 // 普通群组
	GroupTypeDiscussion = 2 // 讨论组
	GroupTypeChatRoom   = 3 // 聊天室
	GroupTypeLiveRoom   = 4 // 普通群组

	GroupMemberNormal = 1 // 正常群成员
	GroupMemberDelete = 2 // 删除群成员

	GroupMemberRoleNormal = 1 // 普通成员
	GroupMemberRoleAdmin  = 2 // 管理员
	GroupMemberRoleMaster = 3 // 群主

	GroupMuteOn  = 1 // 群组禁言开启
	GroupMuteOff = 0 // 群组禁言关闭
)
