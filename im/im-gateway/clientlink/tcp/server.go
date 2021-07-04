package tcp

import (
	"github.com/Allenxuxu/gev"
	"hy-im/im/im-gateway/clientlink/connectionmanger"
	_interface "hy-im/im/im-gateway/clientlink/interface"
	imHandler "hy-im/im/im-gateway/imhandler"
	"runtime"
	"time"
)

type tcpServer struct {
	srv         *gev.Server
	connManager connectionmanger.ConnectionManager
	roomManager connectionmanger.RoomConnectionManager
}

func (t *tcpServer) Start() error {
	if t.srv != nil {
		t.srv.Start()
	}
	return nil
}

func (t *tcpServer) Stop() error {
	if t.srv != nil {
		t.srv.Stop()
	}
	return nil
}

type Options struct {
	Addr     string
	IdleTime int

	ConnManager connectionmanger.ConnectionManager
	ImHandler   imHandler.ImHandler
	RoomManager connectionmanger.RoomConnectionManager
}

type Option func(*Options)

func NewTcpServer(opt ...Option) (_interface.Server, error) {
	opts := Options{}
	for _, v := range opt {
		v(&opts)
	}

	srv, err := gev.NewServer(
		&handler{
			connManager: opts.ConnManager,
			imHandler:   opts.ImHandler,
			roomManager: opts.RoomManager,
		},
		gev.Network("tcp"),
		gev.Address(opts.Addr),
		gev.IdleTime(time.Duration(opts.IdleTime)*time.Second),
		gev.NumLoops(runtime.GOMAXPROCS(-1)),
		gev.Protocol(&protocol{}),
	)
	if err != nil {
		return nil, err
	}

	return &tcpServer{
		srv:         srv,
		connManager: opts.ConnManager,
		roomManager: opts.RoomManager,
	}, nil
}
