package initial

import (
	"fmt"
	innerPt "hy-im/im/im-common/proto/inner"
	"hy-im/im/im-login/conf"
	"hy-im/im/im-login/dao"
	"hy-im/im/im-login/handler"
	"hy-im/im/im-login/service"
	"os"
	"os/signal"
	"syscall"

	dmLog "github.com/common/log"
	dmServer "github.com/common/server"
	"github.com/common/trace"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hyredis"
	imName "hy-im/im/im-common/name"
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
	if err := dmLog.InitLog(func(options *dmLog.Options) {
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
		if err := dmServer.StartHystrix(c.Hystrix.Addr); err != nil {
			log.Fatalf("start hystrix err (%v)", err)
		}
	}

	if c.Prometheus.Enable {
		if err := dmServer.StartPrometheus(c.Prometheus.Addr); err != nil {
			log.Fatalf("start prometheus err (%v)", err)
		}
	}

	// 链路追踪
	io, err := trace.InitTracer(imName.RpcImLogin, c.Trace.Addr)
	if err != nil {
		log.Fatalf("init tracer err (%v)", err)
	}
	defer func() {
		_ = io.Close()
	}()

	// 创建 micro service
	srv := service.NewService(c)

	hyRedis, err := hyredis.InitRedis(&hyredis.RedisConfig{
		Addrs:        c.Redis.Addrs,
		Password:     c.Redis.Password,
		MaxIdleConns: c.Redis.MaxIdleConns,
		MaxOpenConns: c.Redis.MaxOpenConns,
		MaxLifeTime:  c.Redis.MaxLifeTime,
	})
	if err != nil {
		log.Fatalf("init redis err (%v)", err)
	}
	cacheDao := dao.NewCacheOperator(hyRedis)
	loginHandler := handler.Handler{CacheDao: cacheDao}
	if err := innerPt.RegisterImLoginHandler(srv.Server(), &loginHandler); err != nil {
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
