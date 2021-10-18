package cache

type UserTokenInfo struct {
	UserId      uint32 `json:"user_id"`
	UserName    string `json:"user_name"`
	ExpiredTime int64  `json:"expired_time"`
	KickOutTime int64  `json:"kick_out_time"`
}

type CacheDao interface {
	UserTokenDao

	PcLoginDao
}

type UserTokenDao interface {
	SaveUserToken(loginToken string, info *UserTokenInfo, expired int64) error
	GetUserToken(loginToken string) (*UserTokenInfo, error)
	DeleteUserToken(loginToken string) error
}

type PcUniqueSignInfo struct {
	UserId     uint32 `json:"user_id"`
	UserName   string `json:"user_name"`
	LoginToken string `json:"login_token"`
}

type PcLoginDao interface {
	SavePcUniqueSignInfo(uniqueSign string, info *PcUniqueSignInfo, expired int64) error
	GetPcUniqueSignInfo(uniqueSign string) (*PcUniqueSignInfo, error)
}
