package ws

import (
	"context"
	"fmt"
	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/gev/plugins/websocket/ws"
	"github.com/tal-tech/go-zero/core/logx"
	"hy-im/im/im-common/base"
	appPt "hy-im/im/im-common/proto/app"
	innerPt "hy-im/im/im-common/proto/inner"
	"hy-im/im/im-gateway/clientlink/connectionmanger"
	"hy-im/im/im-gateway/common"
	imHandler "hy-im/im/im-gateway/handler"
)

type handler struct {
	connManager connectionmanger.ConnectionManager
	imHandler imHandler.ImHandler
	roomManager connectionmanger.RoomConnectionManager
}

func (h *handler) OnConnect(c *connection.Connection) {
	logx.Infof("connection remote address %s", c.PeerAddr())
}

func (h *handler) OnMessage(c *connection.Connection, msg []byte) (ws.MessageType, []byte) {
	logx.Infof("message %s   \n", string(msg))

	//msg, err := util.PackData(ws.MessageText, msg)
	//if err != nil {
	//	panic(err)
	//}
	//if err := c.Send(msg); err != nil {
	//	msg, err := util.PackCloseData(err.Error())
	//	if err != nil {
	//		panic(err)
	//	}
	//	if e := c.Send(msg); e != nil {
	//		panic(e)
	//	}
	//}

	var rcvMsg Message
	if err := rcvMsg.Deserialize(msg); err != nil {
		// parse fail
		logx.Errorf("except data %s", c.PeerAddr())
		_ = c.Close()
		return ws.MessageBinary, nil
	}
	// login msg
	if rcvMsg.Command == uint16(appPt.ImCmd_cmd_login) {
		if command, content, loginIfo, svcCode, err := h.imHandler.HandlerLogin(context.Background(), int32(rcvMsg.Command), rcvMsg.Content); err != nil {
			logx.Errorf("handler login err(%v)", err)
		} else {
			if command == int32(appPt.ImCmd_cmd_login_ack) && svcCode == int32(innerPt.SrvErr_srv_err_success) {
				c.SetContext(loginIfo)
				h.connManager.AddConnection(fmt.Sprintf("%d:%d", loginIfo.UserId, loginIfo.LoginType), NewWSConnection(c, base.WsConnection, loginIfo))
				sendMessage(c, uint16(command), content, ws.MessageBinary)
			} else {
				sendMessage(c, uint16(command), content, ws.MessageBinary)
				_ = c.Close()
			}
		}
	}

	loginInfo, ok := c.Context().(common.LoginInfo)
	if ok {
		// user no login
		_ = c.Close()
		return ws.MessageBinary, nil
	}

	cmd := rcvMsg.Command
	switch {
	case cmd == uint16(appPt.ImCmd_cmd_heartbeat):
		sendMessage(c, uint16(appPt.ImCmd_cmd_heartbeat_ack), nil, ws.MessageBinary)
	case cmd >= uint16(appPt.ImCmd_cmd_room_msg) && cmd <= 0x3000:
		if command, content, _, svcCode, err := h.imHandler.HandlerRoom(context.Background(), int32(rcvMsg.Command), rcvMsg.Content, loginInfo); err != nil {
			logx.Errorf("handler chat err(%v)", err)
		} else {
			if svcCode != int32(innerPt.SrvErr_srv_err_success) {
				logx.Errorf("server error code %d", svcCode)
			}
			if command > 0 {
				sendMessage(c, uint16(command), content, ws.MessageBinary)
			}
		}
	default:
		logx.Infof("did not handle command(0x%04x)", cmd)
	}


	return ws.MessageBinary, nil
}

func (h *handler) OnClose(c *connection.Connection) {
	if info, ok := c.Context().(common.LoginInfo); ok {
		linkToken := c.PeerAddr()
		err := h.imHandler.HandlerLogout(context.Background(), info, linkToken)
		h.connManager.DelConnection(fmt.Sprintf("%d:%d", info.UserId, info.LoginType), linkToken)
		if err != nil {
			logx.Errorf("handler logout err(%v)", err)
		}
	}
	return
}


func sendMessage(c *connection.Connection, command uint16, content []byte, msgType ws.MessageType) {
	msg := Message{
		Command: command,
		Content: content,
		MsgType: msgType,
	}
	body, err := msg.Serialize()
	if err != nil {
		logx.Errorf("response  serialize err (%v)", err)
		return
	}
	if err := c.Send(body); err != nil {
		logx.Errorf("response  send err (%v)", err)
	}
}
