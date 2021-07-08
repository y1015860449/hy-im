module hy-im/im

go 1.16

require (
	github.com/Allenxuxu/gev v0.2.3
	github.com/Allenxuxu/ringbuffer v0.0.9
	github.com/antonfisher/nested-logrus-formatter v1.3.0 // indirect
	github.com/gobwas/pool v0.2.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/juju/ratelimit v1.0.2-0.20191002062651-f60b32039441
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.3
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spaolacci/murmur3 v1.1.0
	github.com/y1015860449/go-tools v0.0.6
	go.mongodb.org/mongo-driver v1.5.3
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/common => ../../hy-im/common
