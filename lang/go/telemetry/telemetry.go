package telemetry

import (
	"context"
	"net"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
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

func InitTracer(ctx context.Context, opts ...Option) error {
	cfg, err := apply(opts...)
	if err != nil {
		return err
	}

	exporter, err := otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithDialOption(
				grpc.WithBlock(),
			),
			otlptracegrpc.WithEndpoint(cfg.endpoint),
		),
	)
	if err != nil {
		return err
	}

	res, err := sdkresource.New(ctx,
		sdkresource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.sname),
		),
	)
	if err != nil {
		return err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(exporter),
		//		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tp)

	return nil
}
