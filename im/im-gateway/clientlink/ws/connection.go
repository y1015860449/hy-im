package ws

import (
	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/gev/plugins/websocket/ws"
	"github.com/Allenxuxu/gev/plugins/websocket/ws/util"
	"github.com/golang/protobuf/proto"
	appPt "hy-im/im/im-common/proto/app"
	_interface "hy-im/im/im-gateway/clientlink/interface"
	"hy-im/im/im-gateway/common"
)

type Message struct {
	Command     uint16
	Retry       uint16
	SeqNum      int32
	Content     []byte
	MsgType     ws.MessageType
	CloseReason string
}

func (m *Message) Serialize() ([]byte, error) {
	msg := &appPt.WsMsg{
		Cmd:     int32(m.Command),
		Retry:   int32(m.Retry),
		SeqNum:  m.SeqNum,
		Content: m.Content,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	bytes, err := util.PackData(m.MsgType, data)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (m *Message) SerializeClose() ([]byte, error) {
	bytes, err := util.PackCloseData(m.CloseReason)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (m *Message) Deserialize(data []byte) error {
	var msg appPt.WsMsg
	if err := proto.Unmarshal(data, &msg); err != nil {
		return err
	}
	m.MsgType = ws.MessageBinary
	m.Command = uint16(msg.Cmd)
	m.Retry = uint16(msg.Retry)
	m.SeqNum = msg.SeqNum
	m.Content = msg.Content
	return nil
}

type wsConnection struct {
	conn      *connection.Connection
	connType  int
	loginInfo common.LoginInfo
	ctx       interface{}
}

func (w *wsConnection) SetContext(ctx interface{}) {
	w.ctx = ctx
}

func (w *wsConnection) GetContext() interface{} {
	return w.ctx
}

func (w *wsConnection) GetUserId() int64 {
	return w.loginInfo.UserId
}

func (w *wsConnection) GetLoginType() int32 {
	return w.loginInfo.LoginType
}

func (w *wsConnection) ConnectionType() int {
	return w.connType
}

func (w *wsConnection) SendData(data []byte) error {
	if len(data) <= 0 {
		return nil
	}
	return w.conn.Send(data)
}

func (w *wsConnection) LinkToken() string {
	return w.conn.PeerAddr()
}

func (w *wsConnection) Close() error {
	return w.conn.Close()
}

func NewWSConnection(conn *connection.Connection, connType int, info common.LoginInfo) _interface.Connection {
	return &wsConnection{
		conn:      conn,
		connType:  connType,
		loginInfo: info,
	}
}
