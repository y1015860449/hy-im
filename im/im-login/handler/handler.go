package handler

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/tal-tech/go-zero/core/logx"
	appPt "hy-im/im/im-common/proto/app"
	innerPt "hy-im/im/im-common/proto/inner"
	"hy-im/im/im-login/dao"
	"time"
)

type Handler struct {
	CacheDao dao.CacheDao
}

func (h *Handler) Login(ctx context.Context, req *innerPt.LoginReq, rsp *innerPt.LoginRsp) error {
	if req.Command == 0 || (req.Command != int32(appPt.ImCmd_cmd_logout) && len(req.Content) <= 0) {
		logx.Errorf("request param except (%+v)", req)
		rsp.UserId = req.UserId
		rsp.LoginType = req.LoginType
		rsp.SvcErr = int32(innerPt.SrvErr_srv_err_param)
		return nil
	}

	switch req.Command {
	case int32(appPt.ImCmd_cmd_login):
		var login appPt.Login
		if err := proto.Unmarshal(req.Content, &login); err != nil {
			return err
		}
		return h.LoginHandler(ctx, rsp, &login, req.LinkToken)
	case int32(appPt.ImCmd_cmd_logout):
		return h.LogoutHandler(ctx, rsp, req.UserId, req.LoginType, req.LinkToken)
	default:
		logx.Infof("unknown command %d (%+v)", req.Command, req)

	}
	return nil
}

func (h *Handler) LoginHandler(ctx context.Context, rsp *innerPt.LoginRsp, in *appPt.Login, linkToken string) error {
	packLoginRsp := func(appCode, svcCode int32) {
		loginAck := &appPt.LoginAck{
			UserId:    in.UserId,
			ErrCode:   appCode,
			Timestamp: time.Now().UnixNano() / 1e6,
		}
		data, _ := proto.Marshal(loginAck)
		rsp.UserId = in.UserId
		rsp.LoginType = in.LoginType
		rsp.Command = int32(appPt.ImCmd_cmd_login_ack)
		rsp.SvcErr = svcCode
		rsp.Content = data
	}

	if in.UserId <= 0 || len(in.LoginToken) <= 0 || len(in.DeviceToken) <= 0 {
		packLoginRsp(int32(appPt.ImErrCode_err_param_except), int32(innerPt.SrvErr_srv_err_param))
		return nil
	}
	// todo 验证token


	loginInfo, err := h.CacheDao.GetUserLoginInfo(in.UserId, in.LoginType)
	if err != nil {
		logx.Errorf("get user login info err(%+v) userId %d type %d in(%+v)", err, in.UserId, in.LoginType, in)
		packLoginRsp(int32(appPt.ImErrCode_err_server_except), int32(innerPt.SrvErr_srv_err_redis))
		return err
	}

	if loginInfo.Status == dao.ImUserOnline && loginInfo.LinkToken != linkToken {
		// todo 剔除其他登陆
	}

	loginInfo.Status = dao.ImUserOnline
	loginInfo.LoginToken = in.LoginToken
	loginInfo.DeviceToken = in.DeviceToken
	loginInfo.LinkToken = linkToken
	if err := h.CacheDao.SaveUserLoginInfo(in.UserId, in.LoginType, loginInfo); err != nil {
		logx.Errorf("save user login info err(%+v) userId %d type %d in(%+v)", err, in.UserId, in.LoginType, in)
		packLoginRsp(int32(appPt.ImErrCode_err_server_except), int32(innerPt.SrvErr_srv_err_redis))
		return err
	}
	packLoginRsp(int32(appPt.ImErrCode_err_success), int32(innerPt.SrvErr_srv_err_success))
	return nil
}

func (h *Handler) LogoutHandler(ctx context.Context, rsp *innerPt.LoginRsp, userId int64, loginType int32, linkToken string) error {
	packLoginRsp := func(svcCode int32) {
		rsp.SvcErr = svcCode
		rsp.UserId = userId
		rsp.LoginType = loginType
	}
	if len(linkToken) <= 0 || userId <= 0 {
		packLoginRsp(int32(innerPt.SrvErr_srv_err_param))
		return nil
	}
	loginInfo, err := h.CacheDao.GetUserLoginInfo(userId, loginType)
	if err != nil {
		logx.Errorf("get user login info err(%+v) userId %d type %d", err, userId, loginType)
		packLoginRsp(int32(innerPt.SrvErr_srv_err_redis))
		return err
	}
	if loginInfo.Status == dao.ImUserOnline && loginInfo.LinkToken != linkToken {
		if err := h.CacheDao.SetUserLoginStatus(userId, loginType, dao.ImUserOffline); err != nil {
			logx.Errorf("set user login status err(%+v) userId %d type %d", err, userId, loginType)
			packLoginRsp(int32(innerPt.SrvErr_srv_err_redis))
			return err
		}
	}
	packLoginRsp(int32(innerPt.SrvErr_srv_err_success))
	return nil
}


