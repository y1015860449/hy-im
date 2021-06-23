package tcp

import (
	"hy-im/im/im-gateway/clientlink/interface"
	"encoding/binary"
	"errors"
	"github.com/Allenxuxu/gev/connection"
	"hy-im/im/im-gateway/common"
)

const headerLen = 12

type MessageHeader struct {
	Flags   uint8
	AppID   uint8
	Command uint16
	BodyLen uint32
	MsgSeq  uint32
}

type Message struct {
	Header  MessageHeader
	Content []byte
}

func (m *Message) Serialize() ([]byte, error) {
	if len(m.Content) != int(m.Header.BodyLen) {
		return nil, errors.New("data exception")
	}

	data := make([]byte, headerLen)
	data[0] = m.Header.Flags
	data[1] = m.Header.AppID
	binary.BigEndian.PutUint16(data[2:4], m.Header.Command)
	binary.BigEndian.PutUint32(data[4:8], m.Header.BodyLen)
	binary.BigEndian.PutUint32(data[8:headerLen], m.Header.MsgSeq)

	if m.Header.BodyLen > 0 {
		data = append(data, m.Content...)
	}
	return data, nil
}

func (m *Message) Deserialize(data []byte) error {
	if len(data) < headerLen {
		return errors.New("data not enough")
	}
	m.Header.Flags = data[0]
	m.Header.AppID = data[1]
	m.Header.Command = binary.BigEndian.Uint16(data[2:4])
	m.Header.BodyLen = binary.BigEndian.Uint32(data[4:8])
	m.Header.MsgSeq = binary.BigEndian.Uint32(data[8:headerLen])

	if len(data[headerLen:]) != int(m.Header.BodyLen) {
		return errors.New("content not enough")
	}
	m.Content = append(m.Content, data[headerLen:]...)
	return nil
}

type tcpConnection struct {
	conn *connection.Connection
	connType int
	loginInfo common.LoginInfo
	ctx interface{}
}

func (t *tcpConnection) SetContext(ctx interface{}) {
	t.ctx = ctx
}

func (t *tcpConnection) GetContext() interface{} {
	return t.ctx
}

func (t *tcpConnection) GetUserId() int64 {
	return t.loginInfo.UserId
}

func (t *tcpConnection) GetLoginType() int32 {
	return t.loginInfo.LoginType
}

func (t *tcpConnection) GetRoleType() int32 {
	return t.loginInfo.RoleType
}

func (t *tcpConnection) ConnectionType() int {
	return t.connType
}

func (t *tcpConnection) LinkToken() string {
	return t.conn.PeerAddr()
}

func (t *tcpConnection) Close() error {
	return t.conn.Close()
}

func (t *tcpConnection) SendData(data []byte) error {
	if len(data) <= 0 {
		return nil
	}
	return t.conn.Send(data)
}

func NewTcpConnection(conn *connection.Connection, connType int, info common.LoginInfo) _interface.Connection {
	return &tcpConnection{
		conn: conn,
		connType: connType,
		loginInfo: info,
	}
}
