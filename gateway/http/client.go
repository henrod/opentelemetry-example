package http

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func GetHTTPClient() *http.Client {
	// Add tracer
	transport := otelhttp.NewTransport(http.DefaultTransport)

	return &http.Client{
		Transport: transport,
	}
}
