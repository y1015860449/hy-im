package cache

const (
	ImUserOffline = 0
	ImUserOnline = 1
)

type UserLoginInfo struct {
	Status int8
	LoginToken string
	DeviceToken string
	LinkToken string
}


type CacheDao interface {
	GetUserLoginInfo(userId int64, loginType int32) (*UserLoginInfo, error)
	SaveUserLoginInfo(userId int64, loginType int32, info *UserLoginInfo) error
	SetUserLoginStatus(userId int64, loginType int32, status int8) error
}

