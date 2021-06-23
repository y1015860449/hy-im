package trace

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

// InitTracer 初始化 jaeger Tracer
func InitTracer(srvName string, addr string) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: srvName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Reporter(reporter),
		//jaegercfg.Logger(jaeger.StdLogger),
	)

	opentracing.SetGlobalTracer(tracer) // 将jaeger tracer注册到全局
	return closer, err
}
