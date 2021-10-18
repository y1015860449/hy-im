package define

type SvcCode uint32

const (
	SERVICE_SUCCESS uint32 = 0
	//common
	COMMON_PARAM_EXCEPTION     uint32 = 101 //	参数错误
	ERR_UNMARSHAL_FAIL         uint32 = 102 //	解析失败
	COMMON_MYSQL_ERROR         uint32 = 103 //	mysql错误
	COMMON_REDIS_ERROR         uint32 = 104 //  redis错误
	COMMON_MONGODB_ERROR       uint32 = 105 //	mongodb错误
	COMMON_RPC_SERVICE_ERROR   uint32 = 106 //	RPC服务错误
	COMMON_ILLEGAL_USER_ERROR  uint32 = 107 //	非法用户
	COMMON_THIRD_SERVICE_ERROR uint32 = 108 // 第三方错误
	COMMON_MINIO_ERROR         uint32 = 109 //	minio错误

	//用户服务
	USER_SESSION_TYPE_ERROR        uint32 = 11001 //session类型错误
	USER_SESSION_EXIST             uint32 = 11002 //session已经存在
	USER_WX_LOGIN_ERROR            uint32 = 11003 // 微信登陆错误
	USER_LOGIN_TOKEN_AUTH_ERROR    uint32 = 11004 // token验证错误
	USER_LOGIN_TOKEN_EXPIRED_ERROR uint32 = 11005 // token过期
	USER_LOGIN_TOKEN_KICKOUT_ERROR uint32 = 11006 // 在其他设备登陆
	USER_NOT_EXIST_IN_DB_ERROR     uint32 = 11007 // 用户不在数据库中

	//GROUP
	ERR_GROUP_ID_NOT_EXIST         uint32 = 13001 //不允许删除自己
	ERR_GROUP_NOT_ALLOW_DEL_SELF   uint32 = 13002 //不允许删除自己
	ERR_GROUP_NOT_ALLOW_DEL_MASTER uint32 = 13003 //不允许删除群主
	GROUP_DISMISS_ERROR            uint32 = 13004 // 群组已解散
	GROUP_NOT_EXIST_ERROR          uint32 = 13005 // 群组不存在
	GROUP_MEMBER_NOT_EXIT_ERROR    uint32 = 13006 // 不再群内

	//Minio
	ERROR_MINIO_INIT uint32 = 14006 // 不再群内
)
