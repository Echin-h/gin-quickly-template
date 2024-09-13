package tracex

import (
	"context"
	"fmt"
	"gin-quickly-template/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

// https://opentelemetry.io/docs/instrumentation/go/exporters/

// OTLP Exporter
func newOTLPExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	// Change default HTTPS -> HTTP
	insecureOpt := otlptracehttp.WithInsecure()

	// Update default OTLP reciver endpoint
	endPoint := fmt.Sprintf("%s:%s", config.GetConfig().OTel.AgentHost, config.GetConfig().OTel.AgentPort)
	endpointOpt := otlptracehttp.WithEndpoint(endPoint)
	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)
}

func Init() {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.GetConfig().OTel.ServiceName),
		),
	)
	if err != nil {

	}

	provider := sdktrace.NewTracerProvider(sdktrace.WithResource(r))
	otel.SetTracerProvider(provider)
	exp, err := newOTLPExporter(context.Background())
	if err != nil {

	}

	bsp := sdktrace.NewBatchSpanProcessor(exp)
	provider.RegisterSpanProcessor(bsp)
}
