package tcp

import (
	"context"
	"fmt"
	"github.com/Allenxuxu/gev/connection"
	log "github.com/sirupsen/logrus"
	"hy-im/im/im-common/base"
	appPt "hy-im/im/im-common/proto/app"
	innerPt "hy-im/im/im-common/proto/inner"
	"hy-im/im/im-gateway/clientlink/connectionmanger"
	"hy-im/im/im-gateway/common"
	imHandler "hy-im/im/im-gateway/imhandler"
)

type handler struct {
	connManager connectionmanger.ConnectionManager
	imHandler   imHandler.ImHandler
	roomManager connectionmanger.RoomConnectionManager
}

func (h *handler) OnMessage(c *connection.Connection, ctx interface{}, data []byte) []byte {
	if len(data) < headerLen {
		_ = c.Close()
		return nil
	}
	var rcvMsg Message
	if err := rcvMsg.Deserialize(data); err != nil {
		// parse fail
		log.Errorf("except data %s", c.PeerAddr())
		_ = c.Close()
		return nil
	}
	// login msg
	if rcvMsg.Header.Command == uint16(appPt.ImCmd_cmd_login) {
		if command, content, loginIfo, svcCode, err := h.imHandler.HandlerLogin(context.Background(), int32(rcvMsg.Header.Command), rcvMsg.Content); err != nil {
			log.Errorf("handler login err(%v)", err)
		} else {
			if command == int32(appPt.ImCmd_cmd_login_ack) && svcCode == int32(innerPt.SrvErr_srv_err_success) {
				c.SetContext(loginIfo)
				h.connManager.AddConnection(fmt.Sprintf("%d:%d", loginIfo.UserId, loginIfo.LoginType), NewTcpConnection(c, base.TcpConnection, loginIfo))
				sendMessage(c, &rcvMsg.Header, uint16(command), content)
			} else {
				sendMessage(c, &rcvMsg.Header, uint16(command), content)
				_ = c.Close()
			}
		}
	}

	loginInfo, ok := c.Context().(common.LoginInfo)
	if ok {
		// user no login
		_ = c.Close()
		return nil
	}
	cmd := rcvMsg.Header.Command
	switch {
	case cmd == uint16(appPt.ImCmd_cmd_heartbeat):
		sendMessage(c, &rcvMsg.Header, uint16(appPt.ImCmd_cmd_heartbeat_ack), nil)
	case cmd >= uint16(appPt.ImCmd_cmd_room_msg) && cmd <= 0x2000:
		if command, content, svcCode, err := h.imHandler.HandlerRoom(context.Background(), int32(rcvMsg.Header.Command), int32(rcvMsg.Header.Retry), rcvMsg.Content, loginInfo); err != nil {
			log.Errorf("handler chat err(%v)", err)
		} else {
			if svcCode != int32(innerPt.SrvErr_srv_err_success) {
				log.Errorf("server error code %d", svcCode)
			}
			if command > 0 {
				sendMessage(c, &rcvMsg.Header, uint16(command), content)
			}
		}
	default:
		log.Infof("did not handle command(0x%04x)", cmd)
	}

	return nil
}

func (h *handler) OnClose(c *connection.Connection) {
	if info, ok := c.Context().(common.LoginInfo); ok {
		linkToken := c.PeerAddr()
		err := h.imHandler.HandlerLogout(context.Background(), info, linkToken)
		h.connManager.DelConnection(fmt.Sprintf("%d:%d", info.UserId, info.LoginType), linkToken)
		if err != nil {
			log.Errorf("handler logout err(%v)", err)
		}
	}
	return
}

func (h *handler) OnConnect(c *connection.Connection) {
	log.Infof("connection remote address %s", c.PeerAddr())
}

func sendMessage(c *connection.Connection, header *MessageHeader, command uint16, content []byte) {
	msg := Message{
		Header: MessageHeader{
			Flags:   header.Flags,
			AppID:   header.AppID,
			Command: command,
			BodyLen: uint32(len(content)),
			MsgSeq:  header.MsgSeq,
		},
		Content: content,
	}
	body, err := msg.Serialize()
	if err != nil {
		log.Errorf("response  serialize err (%v)", err)
		return
	}
	if err := c.Send(body); err != nil {
		log.Errorf("response  send err (%v)", err)
	}
}
