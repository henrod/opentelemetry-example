package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

const (
	ExporterTypeFile = iota
	ExporterTypeJaeger
)

func ConfigureTracer() (func() error, error) {
	exporter, closeExporter, err := getExporter(ExporterTypeJaeger)
	if err != nil {
		return nil, fmt.Errorf("failed to get exporter: %w", err)
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

		err = closeExporter()
		if err != nil {
			return fmt.Errorf("failed to close exporter: %w", err)
		}

		return nil
	}

	return closeFunction, nil
}

func getExporter(exporterType int) (trace.SpanExporter, func() error, error) {
	switch exporterType {
	case ExporterTypeFile:
		tracesFile, err := os.Create("traces.txt")
		if err != nil {
			return nil, nil, fmt.Errorf("faield to open traces file: %w", err)
		}

		exporter, err := stdouttrace.New(
			stdouttrace.WithWriter(tracesFile),
			stdouttrace.WithPrettyPrint(),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create trace exporter: %w", err)
		}

		return exporter, func() error {
			err = tracesFile.Close()
			if err != nil {
				return fmt.Errorf("failed to close traces file: %w", err)
			}

			return nil
		}, nil
	case ExporterTypeJaeger:
		exporter, err := jaeger.New(
			jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create jaeger exporter: %w", err)
		}

		return exporter, func() error { return nil }, nil
	}

	return nil, nil, errors.New("invalid exporter type")
}
