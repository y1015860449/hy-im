package initial

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hy_log/hy_logrus"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_hystrix"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_prometheus"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_tracer"
	imName "hy-im/api/api-common/name"
	innerPt "hy-im/api/api-common/proto/inner"
	"hy-im/api/api-group/conf"
	"hy-im/api/api-group/handler"
	"hy-im/api/api-group/service"
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
	io, err := hy_tracer.InitTracer(imName.RpcApiGroup, c.Trace.Addr)
	if err != nil {
		log.Fatalf("init tracer err (%v)", err)
	}
	defer func() {
		_ = io.Close()
	}()

	// 创建 micro service
	srv := service.NewService(c)

	//hyRedis, err := hyredis.InitRedis(&hyredis.RedisConfig{
	//	Addrs:        c.Redis.Addrs,
	//	Password:     c.Redis.Password,
	//	MaxIdleConns: c.Redis.MaxIdleConns,
	//	MaxOpenConns: c.Redis.MaxOpenConns,
	//	MaxLifeTime:  c.Redis.MaxLifeTime,
	//})
	//if err != nil {
	//	log.Fatalf("init redis err (%v)", err)
	//}

	// todo 连接mysql

	groupHandler := handler.Handler{}
	if err := innerPt.RegisterApiUserHandler(srv.Server(), &groupHandler); err != nil {
		log.Fatalf("register login srv handler err (%v)", err)
	}

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
