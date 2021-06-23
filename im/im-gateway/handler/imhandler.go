package handler

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	appPt "hy-im/im/im-common/proto/app"
	inner "hy-im/im/im-common/proto/inner"
	"hy-im/im/im-gateway/common"
)

type ImHandler interface {
	HandlerLogin(ctx context.Context, command int32, content []byte) (int32, []byte, common.LoginInfo, int32, error)
	HandlerLogout(ctx context.Context, loginInfo common.LoginInfo, linkToken string) error
	HandlerRoom(ctx context.Context, command int32, content []byte, loginInfo common.LoginInfo)(int32, []byte, common.LoginInfo, int32, error)
}


type imHandler struct {
	loginCli inner.ImLoginService
}

type Options struct {
	LoginCli inner.ImLoginService
}

type Option func(*Options)

func (i *imHandler) HandlerLogin(ctx context.Context, command int32, content []byte) (int32, []byte, common.LoginInfo, int32, error) {
	loginReq := &inner.LoginReq{
		Command: command,
		Content: content,
	}
	loginRsp, err := i.loginCli.Login(ctx, loginReq)
	if err != nil {
		return 0, nil, common.LoginInfo{}, 0, err
	}
	return loginRsp.Command, loginRsp.Content, common.LoginInfo{UserId: loginRsp.UserId, LoginType: loginRsp.LoginType}, loginRsp.SvcErr, nil
}

func (i *imHandler) HandlerLogout(ctx context.Context, loginInfo common.LoginInfo, linkToken string) error {
	loginReq := &inner.LoginReq{
		UserId: loginInfo.UserId,
		LoginType: loginInfo.LoginType,
		RoleType: loginInfo.RoleType,
		LinkToken: linkToken,
		Command: int32(appPt.ImCmd_cmd_logout),
	}
	loginRsp, err := i.loginCli.Login(ctx, loginReq)
	if err != nil {
		logx.Errorf("rpc login err(%+v) rsp(%+v)", err, loginRsp)
		return err
	}
	return nil
}

func (i *imHandler) HandlerRoom(ctx context.Context, command int32, content []byte, loginInfo common.LoginInfo) (int32, []byte, common.LoginInfo, int32, error) {
	panic("implement me")
}

func NewImHandler(opt ...Option) ImHandler {
	opts := Options{}
	for _, value := range opt {
		value(&opts)
	}
	return &imHandler{loginCli: opts.LoginCli}
}
