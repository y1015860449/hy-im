package cache

type TokenInfo struct {
	UserId int64
	ExpiredTime int64
	KickedTime int64
}

type UserCacheDao interface {
	GetToken(token string) (*TokenInfo, error)
	SetToken(token string, info *TokenInfo, keyExpiredSec int64) error
}
