package service

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hy_log/hy_logrus"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_tracer"
	"hy-im/api/api-common/name"
	"hy-im/api/api-gateway/conf"
	"hy-im/api/api-gateway/router"
	"io"
	"os"
	"time"
)

// 使用gin接管micro web路由
func StartWeb(conf *conf.Config) error {
	// create new web service
	service := web.NewService(
		web.Name(name.RpcApiGateway),
		web.Version(conf.Micro.Version),
		web.Address(conf.Http.Addr),
		web.Registry(getEtcdRegistry(conf)),
		web.RegisterTTL(time.Duration(conf.Micro.RegisterTTL)*time.Second),
		web.RegisterInterval(time.Duration(conf.Micro.RegisterInterval)*time.Second),
		web.Flags(&cli.StringFlag{Name: "conf", Usage: "this is conf"}), // fix "flag provided but not defined: -conf"
	)

	// 初始化服务, micro web service 需要调用其他 micro service 时需要init
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	writer := hy_logrus.GetWriter(func(options *hy_logrus.Options) {
		options.FileName = conf.Log.File
		options.Release = conf.Release
		options.Level = conf.Log.Level
		options.MaxAge = conf.Log.MaxAge
		options.RotationTime = conf.Log.RotationTime
	})

	gin.DefaultWriter = io.MultiWriter(writer, os.Stdout)

	if conf.Release {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
	engine := gin.Default()

	// gin 加入 tracer
	engine.Use(hy_tracer.TracerWrapper)

	cli := client.DefaultClient
	_ = cli
	// 管理路由
	router.NewRouter(func(opt *router.Options) {

	}).Route(engine)

	// 默认使用gin的路由器
	service.Handle("/", engine)

	return service.Run()
}

// 使用etcd作为默认注册中心
func getEtcdRegistry(conf *conf.Config) registry.Registry {
	return etcdv3.NewRegistry(func(opt *registry.Options) {
		opt.Addrs = conf.Etcd.Addrs
		//etcdv3.Auth("username", "password")(opt)
	})
}
