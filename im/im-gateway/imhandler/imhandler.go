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
	HandlerRoom(ctx context.Context, command int32, content []byte, loginInfo common.LoginInfo) (int32, []byte, int32, error)
}

type imHandler struct {
	loginCli innerPt.ImLoginService
	roomCli  innerPt.ImRoomService

	connManager connectionmanger.ConnectionManager
	roomManager connectionmanger.RoomConnectionManager
}

type Options struct {
	LoginCli innerPt.ImLoginService
	RoomCli  innerPt.ImRoomService

	ConnManager connectionmanger.ConnectionManager
	RoomManager connectionmanger.RoomConnectionManager
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

func (i *imHandler) HandlerRoom(ctx context.Context, command int32, content []byte, loginInfo common.LoginInfo) (int32, []byte, int32, error) {
	conn := i.connManager.GetConnection(fmt.Sprintf("%d:%d", loginInfo.UserId, loginInfo.LoginType))
	connCtx := conn.GetContext().(*common.ConnectionCtx)
	roomReq := &innerPt.RoomReq{
		UserId:    loginInfo.UserId,
		LoginType: loginInfo.LoginType,
		RoleType:  connCtx.RoleType,
		RoomId:    connCtx.RoomId,
		Command:   command,
		Content:   content,
	}
	roomRsp, err := i.roomCli.Room(ctx, roomReq)
	if err != nil {
		log.Errorf("rpc login err(%+v) rsp(%+x)", err, roomRsp)
		if roomRsp == nil {
			return 0, nil, 0, err
		}
	}
	i.roomManagerHandler(roomRsp)

	return roomRsp.Command, roomRsp.Content, roomRsp.SvcErr, nil
}

func NewImHandler(opt ...Option) ImHandler {
	opts := Options{}
	for _, value := range opt {
		value(&opts)
	}
	return &imHandler{
		loginCli:    opts.LoginCli,
		roomCli:     opts.RoomCli,
		connManager: opts.ConnManager,
		roomManager: opts.RoomManager,
	}
}

func (i *imHandler) roomManagerHandler(rsp *innerPt.RoomRsp) {
	if rsp.SvcErr != int32(innerPt.SrvErr_srv_err_success) {
		return
	}
	key := fmt.Sprintf("%d:%d", rsp.UserId, rsp.LoginType)
	conn := i.connManager.GetConnection(key)
	connCtx := conn.GetContext().(common.ConnectionCtx)
	switch rsp.Command {
	case int32(appPt.ImCmd_cmd_room_open_ack),
		int32(appPt.ImCmd_cmd_room_join_ack):
		connCtx.RoomId = rsp.RoomId
		conn.SetContext(connCtx)
		i.roomManager.AddConnection(rsp.RoomId, key, conn)
	case int32(appPt.ImCmd_cmd_room_quit_ack):
		connCtx.RoomId = 0
		conn.SetContext(connCtx)
		i.roomManager.DelConnection(rsp.RoomId, key)
	case int32(appPt.ImCmd_cmd_room_close_ack):
		i.roomManager.DelRoom(rsp.RoomId)
	case int32(appPt.ImCmd_cmd_room_remove_ack):
		log.Infof("remove room %d", rsp.Command)
	default:
		log.Warnf("unknown command %d", rsp.Command)
	}
}
