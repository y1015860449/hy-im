module hy-im/im

go 1.16

replace google.golang.org/grpc v1.28.0 => google.golang.org/grpc v1.26.0

require (
	github.com/Allenxuxu/gev v0.2.3
	github.com/Allenxuxu/ringbuffer v0.0.9
	github.com/common v0.0.0-00010101000000-000000000000
	github.com/gobwas/pool v0.2.0
	github.com/golang/protobuf v1.4.3
	github.com/juju/ratelimit v1.0.2-0.20191002062651-f60b32039441
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.3
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.6.0
	github.com/tal-tech/go-zero v1.1.7
	golang.org/x/tools v0.0.0-20200811215021-48a8ffc5b207 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/common => ../../hy-im/common
