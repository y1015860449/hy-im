package ws

import (
	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/plugins/websocket"
	"github.com/Allenxuxu/gev/plugins/websocket/ws"
	"hy-im/im/im-gateway/clientlink/connectionmanger"
	_interface "hy-im/im/im-gateway/clientlink/interface"
	imHandler "hy-im/im/im-gateway/imhandler"
	"log"
	"runtime"
	"time"
)

func (w *wsServer) Start() error {
	if w.srv != nil {
		w.srv.Start()
	}
	return nil
}

func (w *wsServer) Stop() error {
	if w.srv != nil {
		w.srv.Stop()
	}
	return nil
}

type wsServer struct {
	srv         *gev.Server
	connManager connectionmanger.ConnectionManager
	roomManager connectionmanger.GroupConnectionManager
}

type Options struct {
	Addr     string
	IdleTime int

	ConnManager connectionmanger.ConnectionManager
	ImHandler   imHandler.ImHandler
	RoomManager connectionmanger.GroupConnectionManager
}

type Option func(*Options)

func NewWebsocketServer(opt ...Option) (_interface.Server, error) {
	opts := Options{}
	for _, value := range opt {
		value(&opts)
	}
	log.Printf("%#v", opts)
	handler := &handler{
		connManager: opts.ConnManager,
		imHandler:   opts.ImHandler,
		roomManager: opts.RoomManager,
	}
	wsUpgrader := &ws.Upgrader{}
	srv, err := gev.NewServer(
		websocket.NewHandlerWrap(wsUpgrader, handler),
		gev.Protocol(websocket.New(wsUpgrader)), gev.Network("tcp"),
		gev.Address(opts.Addr),
		gev.NumLoops(runtime.GOMAXPROCS(-1)),
		gev.IdleTime(time.Duration(opts.IdleTime)*time.Second))
	if err != nil {
		return nil, err
	}

	return &wsServer{
		srv:         srv,
		connManager: opts.ConnManager,
		roomManager: opts.RoomManager,
	}, nil
}
