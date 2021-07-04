package dao

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hyredis"
	"strconv"
)

const (
	loginStatus = "status"
	loginToken = "loginToken"
	loginDeviceToken = "deviceToken"
	loginLinkToken = "linkToken"
)

type cacheOperator struct {
	rdCli *hyredis.HyRedis
}

func (c *cacheOperator) GetUserLoginInfo(userId int64, loginType int32) (*UserLoginInfo, error) {
	result, err := c.rdCli.HGetAll(getUserLoginInfoKey(userId, loginType))
	if err != nil {
		return nil, err
	}
	loginInfo := &UserLoginInfo{}
	for key, value := range result {
		switch key {
		case loginStatus:
			status, _ := strconv.Atoi(value)
			loginInfo.Status = int8(status)
		case loginToken:
			loginInfo.LoginToken = value
		case loginDeviceToken:
			loginInfo.DeviceToken = value
		case loginLinkToken:
			loginInfo.LinkToken = value

		default:
			log.Warnf("unknown key %s", key)
		}
	}
	return loginInfo, nil
}

func (c *cacheOperator) SaveUserLoginInfo(userId int64, loginType int32, info *UserLoginInfo) error {
	fields := make(map[string]string, 0)
	fields[loginStatus] = strconv.FormatInt(int64(info.Status), 10)
	fields[loginToken] = info.LoginToken
	fields[loginDeviceToken] = info.DeviceToken
	fields[loginLinkToken] = info.LinkToken
	return c.rdCli.HMSet(getUserLoginInfoKey(userId, loginType), fields)
}

func (c *cacheOperator) SetUserLoginStatus(userId int64, loginType int32, status int8) error {
	return c.rdCli.HSet(getUserLoginInfoKey(userId, loginType), loginStatus, strconv.FormatInt(int64(status), 10))
}

func NewCacheOperator(rdCli *hyredis.HyRedis) CacheDao {
	return &cacheOperator{rdCli: rdCli}
}


func getUserLoginInfoKey(userId int64, loginType int32) string {
	return fmt.Sprintf("login:%d:%d", userId, loginType)
}