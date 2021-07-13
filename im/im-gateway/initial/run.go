package initial

import (
	"fmt"
	"github.com/micro/go-micro/v2/client"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hy_log/hy_logrus"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_hystrix"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_prometheus"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_tracer"
	imName "hy-im/im/im-common/name"
	"hy-im/im/im-common/proto/inner"
	"hy-im/im/im-gateway/clientlink/connectionmanger"
	"hy-im/im/im-gateway/clientlink/tcp"
	"hy-im/im/im-gateway/clientlink/ws"
	"hy-im/im/im-gateway/conf"
	"hy-im/im/im-gateway/imhandler"
	"hy-im/im/im-gateway/service"
	"os"
	"os/signal"
	"syscall"
)

func Run(f string) {

	c, err := conf.NewConfig(f)
	if err != nil {
		fmt.Printf("init config err (%v)\n", err)
		panic(err)
	}
	if !c.Release {
		fmt.Printf("load config success (%v)\n", c)
	}
	if err := hy_logrus.InitLog(func(options *hy_logrus.Options) {
		options.FileName = c.Log.File
		options.Release = c.Release
		options.Level = c.Log.Level
		options.MaxAge = c.Log.MaxAge
		options.RotationTime = c.Log.RotationTime
	}); err != nil {
		log.Fatalf("init log err (%v)", err)
	}
	log.Infoln("init log success")

	if c.Hystrix.Enable {
		if err := hy_hystrix.StartHystrix(c.Hystrix.Addr); err != nil {
			log.Fatalf("start hystrix err (%v)", err)
		}
	}

	if c.Prometheus.Enable {
		if err := hy_prometheus.StartPrometheus(c.Prometheus.Addr); err != nil {
			log.Fatalf("start prometheus err (%v)", err)
		}
	}

	// 链路追踪
	io, err := hy_tracer.InitTracer(imName.RpcImGateway, c.Trace.Addr)
	if err != nil {
		log.Fatalf("init tracer err (%v)", err)
	}
	defer io.Close()

	// 创建 micro service
	srv := service.NewService(c)

	connManger := connectionmanger.NewConnManager()
	roomManager := connectionmanger.NewRoomConnectionManager()
	cl := client.DefaultClient
	loginRpc := inner.NewImLoginService(imName.RpcImLogin, cl)
	roomRpc := inner.NewImRoomService(imName.RpcImGroup, cl)
	opts := &imhandler.Options{
		LoginCli:     loginRpc,
		GroupCli:     roomRpc,
		ConnManager:  connManger,
		GroupManager: roomManager,
	}
	handler := imhandler.NewImHandler(func(options *imhandler.Options) {
		options = opts
	})

	tcpOpts := &tcp.Options{
		Addr:        c.Tcp.Addr,
		IdleTime:    c.Tcp.IdleTime,
		ConnManager: connManger,
		ImHandler:   handler,
	}
	tcpSrv, err := tcp.NewTcpServer(func(options *tcp.Options) {
		options = tcpOpts
	})
	if err != nil {
		log.Fatalf("start tcp server err(%+v)", err)
	}
	go func() {
		_ = tcpSrv.Start()
	}()

	wsOpts := &ws.Options{
		Addr:        c.Tcp.Addr,
		IdleTime:    c.Tcp.IdleTime,
		ConnManager: connManger,
		ImHandler:   handler,
	}
	wsSrv, err := ws.NewWebsocketServer(func(options *ws.Options) {
		options = wsOpts
	})
	if err != nil {
		log.Fatalf("start websocket server err(%+v)", err)
	}
	go func() {
		_ = wsSrv.Start()
	}()

	// 运行
	if err := srv.Run(); err != nil {
		log.Fatalf("run server err (%v)", err)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		log.Infof("recv system shutdown signal. signal is %d", sig)
	}()
}
