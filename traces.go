package main

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func ConfigureTracer() (func() error, error) {
	tracesFile, err := os.Create("traces.txt")
	if err != nil {
		return nil, fmt.Errorf("faield to open traces file: %w", err)
	}

	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(tracesFile),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	traceResource, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("opentelemetry-example"),
			semconv.ServiceVersionKey.String("v0.1.0"),
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

		err = tracesFile.Close()
		if err != nil {
			return fmt.Errorf("failed to close traces file: %w", err)
		}

		return nil
	}

	return closeFunction, nil
}
