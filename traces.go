package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

const jaegerURL = "http://localhost:14268/api/traces"

func ConfigureTracer() (func() error, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL)))
	if err != nil {
		return nil, fmt.Errorf("failed to create jaeger exporter: %w", err)
	}

	traceResource, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(applicationName),
			semconv.ServiceVersionKey.String(version),
			attribute.String("environment", "example"),
		),
	)

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(traceResource),
	)

	otel.SetTracerProvider(tp)

	closeFunction := func() error {
		err = tp.Shutdown(context.Background())
		if err != nil {
			return fmt.Errorf("failed to shutdown tracer provider: %w", err)
		}

		return nil
	}

	return closeFunction, nil
}
