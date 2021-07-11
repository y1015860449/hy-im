module hy-im/api

go 1.16

require (
	github.com/common v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.4.3
	github.com/juju/ratelimit v1.0.2-0.20191002062651-f60b32039441
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.3
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/y1015860449/go-tools v0.0.6
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/common => ../../hy-im/common
