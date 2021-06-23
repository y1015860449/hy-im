package service

import (
	jujuRL "github.com/juju/ratelimit"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	goTracing "github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	imName "hy-im/im/im-common/name"
	"hy-im/im/im-gateway/conf"
	"time"
)

func NewService(conf *conf.Config) micro.Service {
	// 使用etcd作为默认注册中心
	etcdRegistry := etcdv3.NewRegistry(func(opt *registry.Options) {
		opt.Addrs = conf.Etcd.Addrs
	})

	limit := conf.Micro.LimitRate // 限制每秒请求数
	rlBucket := jujuRL.NewBucketWithRate(float64(limit), int64(limit))

	// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	srv := micro.NewService(
		micro.Name(imName.RpcImGateway),
		micro.Version(conf.Micro.Version),
		micro.RegisterTTL(time.Duration(conf.Micro.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(conf.Micro.RegisterInterval)*time.Second),
		micro.Registry(etcdRegistry),         // 使用etcd注册中心
		micro.Transport(grpc.NewTransport()), // 默认传输方式模式是http,改为grpc
		micro.WrapHandler(
			opentracing.NewHandlerWrapper(goTracing.GlobalTracer()), // 链路追踪
			prometheus.NewHandlerWrapper(),                          // 监控
			ratelimit.NewHandlerWrapper(rlBucket, false),            // 限流
		),
	)

	log.Infoln("init service success")
	return srv
}
