package app

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// AppTracer struct
type AppTracer struct {
	TracerProvider *tracesdk.TracerProvider
}

// CreateAppTracer function
func CreateAppTracer(cfg *AppConfig) (at *AppTracer, err error) {
	at = new(AppTracer)

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.JaegerURL)))
	if err != nil {
		return
	}

	at.TracerProvider = tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.ServiceName),
			attribute.String("environment", cfg.Environment),
			attribute.Int64("ID", cfg.InstanceID),
		)),
	)

	otel.SetTracerProvider(at.TracerProvider)

	return
}
