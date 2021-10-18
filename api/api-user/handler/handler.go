package handler

import (
	"context"
	"github.com/common/base"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hy_utils"
	innerPt "hy-im/api/api-common/proto/inner"
	"hy-im/api/api-user/conf"
	"hy-im/api/api-user/dao/cache"
	"hy-im/api/api-user/dao/db/model"
	"hy-im/api/api-user/dao/db/user"
	"time"
)

type Handler struct {
	userDao   user.ImUserDao
	cacheDao  cache.CacheDao
	tokenConf *conf.Token
}

func (p *Handler) Register(ctx context.Context, req *innerPt.RegisterReq, rsp *innerPt.RegisterRsp) error {
	if len(req.Account) == 0 {
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_param)
		return nil
	}
	userInfo, err := p.userDao.GetUserInfoByAccount(req.Account)
	if err != nil {
		log.Errorf("get user info err(%v)", err)
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_mysql)
		return err
	}
	if userInfo.Account == req.Account {
		rsp.UserId = userInfo.Id
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_account_already)
		return nil
	}
	userInfo.Account = req.Account
	userInfo.Status = base.UserNormal
	if err := p.userDao.SaveUserInfo(userInfo); err != nil {
		log.Errorf("save user info err(%v)", err)
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_mysql)
		return nil
	}
	loginToken, code := p.getToken(userInfo.Id)
	if code != int32(innerPt.SrvErr_srv_err_success) {
		rsp.SvcErr = code
		return nil
	}
	loginInfo := &model.UserLogins{
		UserId:     userInfo.Id,
		Account:    req.Account,
	}
	if err := p.userDao.SaveLoginInfo(loginInfo); err != nil {
		log.Errorf("save login info err(%v)", err)
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_mysql)
		return nil
	}
	rsp.Account = req.Account
	rsp.UserId = userInfo.Id
	rsp.LoginToken = loginToken
	rsp.SvcErr = int32(innerPt.SrvErr_srv_err_success)
	return nil
}

func (h *Handler) Login(ctx context.Context, req *innerPt.LoginReq, rsp *innerPt.LoginRsp) error {
	if len(req.Account) == 0 {
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_param)
		return nil
	}
	userInfo, err := h.userDao.GetUserInfoByAccount(req.Account)
	if err != nil {
		log.Errorf("get user info err(%v)", err)
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_mysql)
		return err
	}
	if userInfo.Account != req.Account {
		rsp.UserId = userInfo.Id
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_account_not_exist)
		return nil
	}
	// todo 获取最近的登录信息，剔除

	//
	loginToken, code := h.getToken(userInfo.Id)
	if code != int32(innerPt.SrvErr_srv_err_success) {
		rsp.SvcErr = code
		return nil
	}
	loginInfo := &model.UserLogins{
		UserId:     userInfo.Id,
		Account:    req.Account,
	}
	if err := h.userDao.SaveLoginInfo(loginInfo); err != nil {
		log.Errorf("save login info err(%v)", err)
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_mysql)
		return nil
	}
	rsp.Account = req.Account
	rsp.UserId = userInfo.Id
	rsp.LoginToken = loginToken
	rsp.SvcErr = int32(innerPt.SrvErr_srv_err_success)
	return nil
}

func (p *Handler) Logout(ctx context.Context, req *innerPt.LogoutReq, rsp *innerPt.LogoutRsp) error {
	if req.UserId == 0 || len(req.LoginToken) <= 0 {
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_param)
		return nil
	}
	code, _, _ := p.CheckToken(req.LoginToken, req.UserId, "")
	if code != int32(innerPt.SrvErr_srv_err_success) {
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_user_auth)
		return nil
	}
	// todo 删除token

	rsp.UserId = req.UserId
	rsp.SvcErr = int32(innerPt.SrvErr_srv_err_success)
	return nil
}
