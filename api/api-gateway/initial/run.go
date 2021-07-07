package initial

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/y1015860449/go-tools/hy_log/hy_logrus"
	"github.com/y1015860449/go-tools/hy_servicectrl/hy_tracer"
	"hy-im/api/api-common/name"
	"hy-im/api/api-gateway/conf"
	"hy-im/api/api-gateway/service"
)

// 执行
func Run(f string) {

	config, err := conf.NewConfig(f)
	if err != nil {
		fmt.Printf("init config err (%v)\n", err)
		panic(err)
	}
	if !config.Release {
		fmt.Printf("load config success (%v)\n", config)
	}
	if err = hy_logrus.InitLog(func(options *hy_logrus.Options) {
		options.FileName = config.Log.File
		options.Release = config.Release
		options.Level = config.Log.Level
		options.MaxAge = config.Log.MaxAge
		options.RotationTime = config.Log.RotationTime
	}); err != nil {
		log.Fatalf("init log err (%v)", err)
	}
	log.Infoln("init log success")

	// 链路追踪
	io, err := hy_tracer.InitTracer(name.RpcApiGateway, config.Trace.Addr)
	if err != nil {
		log.Errorf("init tracer err (%v)", err)
	} else {
		defer io.Close()
		log.Infoln("init tracer success")
	} // 程序退出前不能关闭,所以放在这里

	// 创建 micro web , 接入gin的路由到 handler.router.go 里
	if err := service.StartWeb(config); err != nil {
		log.Fatalf("start web error (%v)", err)
	}
}
