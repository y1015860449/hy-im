package handler

import (
	"InstantCommunication/api/apiuser/api.user"
	"InstantCommunication/api/apiuser/dao/cache"
	"errors"
	"github.com/common/define"
	"github.com/common/log"
	"github.com/y1015860449/go-tools/hy_utils"
	"time"
)

func (p *Handler) AuthTokenHandler(in *api_user.AuthTokenRequest) (*api_user.AuthTokenResponse, error) {
	// todo 确认规则，后处理

	return &api_user.AuthTokenResponse{Code: define.SERVICE_SUCCESS}, nil
}

func (p *Handler) MakeToken(userId uint32, userName string) (string, error) {
	if userId <= 0 || len(userName) <= 0 {
		return "", errors.New("param is exception")
	}
	now := time.Now().Unix()
	expiredTime := now + int64(p.conf.Token.ExpiredSec)
	token := hy_utils.GetUUID()
	keyExpiredTime := int64(p.conf.Token.ExpiredSec) + int64(p.conf.Token.SafetySec)
	info := &cache.UserTokenInfo{
		UserId:      userId,
		UserName:    userName,
		ExpiredTime: expiredTime,
		KickOutTime: 0,
	}
	if err := p.cacheDao.SaveUserToken(token, info, keyExpiredTime); err != nil {
		log.SugarLog.Errorf("cache save user token err(%#v)", err)
		return "", err
	}
	return token, nil
}

func (p *Handler) CheckToken(token string, userId uint32, userName string) (define.SvcCode, int64, error) {
	now := time.Now().Unix()
	if len(token) <= 0 || (userId <= 0 && len(userName) <= 0) {
		return define.SvcCode(define.COMMON_PARAM_EXCEPTION), now, errors.New("param is exception")
	}
	info, err := p.cacheDao.GetUserToken(token)
	if err != nil {
		log.SugarLog.Error("get user token err(%#v)", err)
		return define.SvcCode(define.COMMON_REDIS_ERROR), now, err
	}

	if info.UserId <= 0 && info.UserName == "" && info.ExpiredTime <= 0 {
		return define.SvcCode(define.USER_LOGIN_TOKEN_AUTH_ERROR), now, nil
	}
	if (userId != 0 && userId != info.UserId) || (len(userName) > 0 && userName != info.UserName) {
		return define.SvcCode(define.USER_LOGIN_TOKEN_AUTH_ERROR), now, nil
	}

	if info.KickOutTime != 0 {
		return define.SvcCode(define.USER_LOGIN_TOKEN_KICKOUT_ERROR), info.KickOutTime, nil
	}

	if time.Now().Unix() > info.ExpiredTime {
		return define.SvcCode(define.USER_LOGIN_TOKEN_EXPIRED_ERROR), now, nil
	}

	// 更新redis信息
	info.ExpiredTime = now + int64(p.conf.Token.ExpiredSec)
	keyExpiredTime := int64(p.conf.Token.ExpiredSec) + int64(p.conf.Token.SafetySec)
	if err := p.cacheDao.SaveUserToken(token, info, keyExpiredTime); err != nil {
		log.SugarLog.Errorf("save user token err(%#v)", err)
		return define.SvcCode(define.COMMON_REDIS_ERROR), now, err
	}
	return define.SvcCode(define.SERVICE_SUCCESS), now, err
}

func (p *Handler) KickOutToken(token string, kickTime int64) error {
	if len(token) <= 0 || kickTime <= 0 {
		return errors.New("param is exception")
	}
	info, err := p.cacheDao.GetUserToken(token)
	if err != nil {
		log.SugarLog.Error("get user token err(%#v)", err)
		return err
	}
	if info.UserId <= 0 && info.UserName == "" && info.ExpiredTime <= 0 {
		return nil
	}
	if info.KickOutTime > 0 {
		return nil
	}
	info.KickOutTime = kickTime
	if err := p.cacheDao.SaveUserToken(token, info, int64(p.conf.Token.SafetySec)); err != nil {
		log.SugarLog.Errorf("save user token err(%#v)", err)
		return err
	}
	return nil
}
