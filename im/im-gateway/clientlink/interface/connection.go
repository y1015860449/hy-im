package _interface

type Connection interface {
	SendData(data []byte) error
	LinkToken() string
	ConnectionType() int
	GetUserId() int64
	GetLoginType() int32
	GetRoleType() int32
	SetContext(ctx interface{})
	GetContext() interface{}
	Close() error
}