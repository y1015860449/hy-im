package base

type SvcCode int32

const (
	SERVICE_SUCCESS SvcCode = 0
	//common
	COMMON_PARAM_EXCEPTION     SvcCode = 101 //	参数错误
	COMMON_MYSQL_ERROR         SvcCode = 102 //	mysql错误
	COMMON_REDIS_ERROR         SvcCode = 103 //  redis错误
	COMMON_MONGODB_ERROR       SvcCode = 104 //	mongodb错误
	COMMON_RPC_SERVICE_ERROR   SvcCode = 105 //	RPC服务错误
	COMMON_ILLEGAL_USER_ERROR  SvcCode = 106 //	非法用户
	COMMON_THIRD_SERVICE_ERROR SvcCode = 107 // 第三方错误

	//用户服务
	USER_ALREADY_REGISTER_ERROR    SvcCode = 1001 // 用户已注册
	USER_ACCOUNT_FROZEN_ERROR      SvcCode = 1002 // 用户冻结
	USER_ACCOUNT_WRITE_OFF_ERROR   SvcCode = 1003 // 用户注销
	USER_MOBLIE_EXCEPT_ERROR       SvcCode = 1004 // 手机号异常
	USER_EMAIL_EXCEPT_ERROR        SvcCode = 1005 // 邮箱异常
	USER_ACCOUNT_EXCEPT_ERROR      SvcCode = 1006 // 账号异常
	USER_PASSWORD_ERROR            SvcCode = 1007 // 密码异常
	USER_LOGIN_FREQUENTLY_ERROR    SvcCode = 1008 // 频繁登录
	USER_LOGIN_TOKEN_AUTH_ERROR    SvcCode = 1009 // token验证错误
	USER_LOGIN_TOKEN_EXPIRED_ERROR SvcCode = 1010 // token过期
	USER_LOGIN_TOKEN_KICKOUT_ERROR SvcCode = 1011 // 在其他设备登陆
	USER_NOT_EXIST_IN_DB_ERROR     SvcCode = 1012 // 用户不在数据库中

	// 系统服务
	SYSTEM_AUTH_CODE_ERROR            SvcCode = 2001 // 验证码错误
	SYSTEM_AUTH_CODE_FREQUENTLY_ERROR SvcCode = 2002 // 发送验证码频繁
	SYSTEM_AUTH_CODE_SEND_ERROR       SvcCode = 2003 // 发送验证码失败

	//GROUP
	GROUP_ID_NOT_EXIST_ERROR         SvcCode = 3001 // 群不存在
	GROUP_NOT_ALLOW_DEL_SELF_ERROR   SvcCode = 3002 //不允许删除自己
	GROUP_NOT_ALLOW_DEL_MASTER_ERROR SvcCode = 3003 //不允许删除群主
	GROUP_DISMISS_ERROR              SvcCode = 3004 // 群组已解散
	GROUP_NOT_EXIST_ERROR            SvcCode = 3005 // 群组不存在
	GROUP_MEMBER_NOT_EXIT_ERROR      SvcCode = 3006 // 不再群内
)
