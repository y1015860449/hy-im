package imhandler

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	appPt "hy-im/im/im-common/proto/app"
	innerPt "hy-im/im/im-common/proto/inner"
	"hy-im/im/im-gateway/clientlink/connectionmanger"
	"hy-im/im/im-gateway/common"
)

type ImHandler interface {
	HandlerLogin(ctx context.Context, command int32, content []byte) (int32, []byte, common.LoginInfo, int32, error)
	HandlerLogout(ctx context.Context, loginInfo common.LoginInfo, linkToken string) error
	HandlerGroup(ctx context.Context, command, retry int32, content []byte, loginInfo common.LoginInfo) (int32, []byte, int32, error)
}

type imHandler struct {
	loginCli innerPt.ImLoginService
	groupCli innerPt.ImGroupService

	connManager  connectionmanger.ConnectionManager
	groupManager connectionmanger.GroupConnectionManager
}

type Options struct {
	LoginCli innerPt.ImLoginService
	GroupCli innerPt.ImGroupService

	ConnManager  connectionmanger.ConnectionManager
	GroupManager connectionmanger.GroupConnectionManager
}

type Option func(*Options)

func (i *imHandler) HandlerLogin(ctx context.Context, command int32, content []byte) (int32, []byte, common.LoginInfo, int32, error) {
	loginReq := &innerPt.LoginReq{
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
	loginReq := &innerPt.LoginReq{
		UserId:    loginInfo.UserId,
		LoginType: loginInfo.LoginType,
		LinkToken: linkToken,
		Command:   int32(appPt.ImCmd_cmd_logout),
	}
	loginRsp, err := i.loginCli.Login(ctx, loginReq)
	if err != nil {
		log.Errorf("rpc login err(%+v) rsp(%+v)", err, loginRsp)
		return err
	}
	return nil
}

func (i *imHandler) HandlerGroup(ctx context.Context, command, retry int32, content []byte, loginInfo common.LoginInfo) (int32, []byte, int32, error) {
	conn := i.connManager.GetConnection(fmt.Sprintf("%d:%d", loginInfo.UserId, loginInfo.LoginType))
	connCtx := conn.GetContext().(*common.ConnectionCtx)
	groupReq := &innerPt.GroupReq{
		UserId:    loginInfo.UserId,
		LoginType: loginInfo.LoginType,
		RoleType:  connCtx.RoleType,
		GroupId:    connCtx.GroupId,
		Command:   command,
		Retry:     retry,
		Content:   content,
	}
	groupRsp, err := i.groupCli.Group(ctx, groupReq)
	if err != nil {
		log.Errorf("rpc login err(%+v) rsp(%+x)", err, groupRsp)
		if groupRsp == nil {
			return 0, nil, 0, err
		}
	}
	i.groupManagerHandler(groupRsp)

	return groupRsp.Command, groupRsp.Content, groupRsp.SvcErr, nil
}

func NewImHandler(opt ...Option) ImHandler {
	opts := Options{}
	for _, value := range opt {
		value(&opts)
	}
	return &imHandler{
		loginCli:     opts.LoginCli,
		groupCli:     opts.GroupCli,
		connManager:  opts.ConnManager,
		groupManager: opts.GroupManager,
	}
}

func (i *imHandler) groupManagerHandler(rsp *innerPt.GroupRsp) {
	if rsp.SvcErr != int32(innerPt.SrvErr_srv_err_success) {
		return
	}
	key := fmt.Sprintf("%d:%d", rsp.UserId, rsp.LoginType)
	conn := i.connManager.GetConnection(key)
	connCtx := conn.GetContext().(common.ConnectionCtx)
	switch rsp.Command {
	case int32(appPt.ImCmd_cmd_group_open_ack),
		int32(appPt.ImCmd_cmd_group_join_ack):
		connCtx.GroupId = rsp.GroupId
		conn.SetContext(connCtx)
		i.groupManager.AddConnection(rsp.GroupId, key, conn)
	case int32(appPt.ImCmd_cmd_group_quit_ack):
		connCtx.GroupId = 0
		conn.SetContext(connCtx)
		i.groupManager.DelConnection(rsp.GroupId, key)
	case int32(appPt.ImCmd_cmd_group_close_ack):
		i.groupManager.DelGroup(rsp.GroupId)
	case int32(appPt.ImCmd_cmd_group_remove_ack):
		log.Infof("remove group %d", rsp.Command)
	default:
		log.Warnf("unknown command %d", rsp.Command)
	}
}
