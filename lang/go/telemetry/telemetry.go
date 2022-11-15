package telemetry

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type _Config struct {
	sname    string
	endpoint string
}

type Option func(*_Config) error

func WithServiceName(name string) Option {
	return func(opts *_Config) error {
		opts.sname = name
		return nil
	}
}
func WithEndpoint(addr string) Option {
	return func(opts *_Config) error {
		opts.endpoint = addr
		return nil
	}
}

func apply(opts ...Option) (*_Config, error) {
	cfg := &_Config{}
	for _, opt := range opts {
		err := opt(cfg)
		if err != nil {
			return nil, err
		}
	}
	if cfg.sname == "" {
		cfg.sname = os.Getenv("LAZYCAT_APP_ID")
		if cfg.sname == "" {
			cfg.sname = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
		}
	}
	if cfg.endpoint == "" {
		endpoint := "host.lzcapp:4317"
		_, err := net.LookupIP("host.lzcapp")
		if err != nil {
			endpoint = "127.0.0.1:4317"
		}
		cfg.endpoint = endpoint
	}
	return cfg, nil
}

// 等待InitTracer调用后替换实际的tracer
type lazyTracer struct {
	real trace.Tracer
	lock sync.Mutex

	name string
	opts []trace.TracerOption
}

func (l *lazyTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if tp == nil {
		panic("Use Tracer before InitTracer")
	}

	l.lock.Lock()
	if l.real == nil {
		l.real = tp.Tracer(l.name, l.opts...)
	}
	l.lock.Unlock()

	return l.real.Start(ctx, spanName, opts...)
}

var (
	tp   *sdktrace.TracerProvider
	lock sync.RWMutex
)

func Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	if tp != nil {
		return tp.Tracer(name, opts...)
	}
	return &lazyTracer{name: name, opts: opts}
}

func InitTracerProvider(ctx context.Context, opts ...Option) error {
	cfg, err := apply(opts...)
	if err != nil {
		return err
	}

	exporter, err := otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(cfg.endpoint),
		),
	)
	if err != nil {
		return err
	}

	b, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		fmt.Println(err)
		b = []byte{}
	}

	res, err := sdkresource.New(ctx,
		sdkresource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.sname),
		),
		sdkresource.WithAttributes(
			attribute.Key("resource_machine_id").String(string(b)),
		),
	)
	if err != nil {
		return err
	}

	tp = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tp)

	return nil
}

func Shutdown(ctx context.Context) {
	if tp != nil {
		tp.Shutdown(ctx)
	}
}
