package telemetry

import (
	"context"
	"errors"
	"net"
	"os"
	"sync"

	"gitee.com/linakesi/lzc-sdk/lang/go/sys"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	DevoptSocketPath    = "/run/lzc-sys/devopt.sock"
	AltDevoptSocketPath = "/lzcapp/run/sys/devopt.sock"
)

type _Config struct {
	sname    string
	endpoint string
	batcher  bool
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
func WithBatcher(batcher bool) Option {
	return func(opts *_Config) error {
		opts.batcher = batcher
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
	var socketPath string
	if _, err := os.Stat(DevoptSocketPath); err == nil {
		socketPath = DevoptSocketPath
	} else if _, err := os.Stat(AltDevoptSocketPath); err == nil {
		socketPath = AltDevoptSocketPath
	} else {
		return errors.New("devopt socket not found")
	}
	conn, err := grpc.Dial("unix://" + socketPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	devopt := sys.NewDevOptServiceClient(conn)
	resp, err := devopt.GetDeveloperOptions(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}
	if resp.OpenTelemetry {
		return initTracerProvider(ctx, opts...)
	} else {
		return nil
	}
}

func initTracerProvider(ctx context.Context, opts ...Option) error {
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

	res, err := sdkresource.New(ctx,
		sdkresource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.sname),
			semconv.DeviceIDKey.String(os.Getenv("LAZYCAT_BOX_DOMAIN")),
		),
	)
	if err != nil {
		return err
	}

	tracerOpts := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
	}

	if cfg.batcher {
		tracerOpts = append(tracerOpts, sdktrace.WithBatcher(exporter))
	} else {
		tracerOpts = append(tracerOpts, sdktrace.WithSyncer(exporter))
	}

	tp = sdktrace.NewTracerProvider(tracerOpts...)

	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tp)

	return nil
}

func Shutdown(ctx context.Context) {
	if tp != nil {
		tp.Shutdown(ctx)
	}
}
