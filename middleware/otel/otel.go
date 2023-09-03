package otel

import (
	"context"
	"douyin-server/config"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var tp *trace.TracerProvider

func Init(ctx context.Context, serviceName string) {
	// Write telemetry data to jaeger.
	client := otlptracehttp.NewClient(
		otlptracehttp.WithInsecure(), //UI view in http://localhost:16686 not https
		otlptracehttp.WithEndpoint(config.JaegerAddr))
	exp, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	tp = trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)

}

func Close() {
	if err := tp.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
