package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	v1 "opentelemetry-example/protogen/go/api/v1"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type App struct {
	http *http.Client
}

func NewApp() (*App, error) {
	return &App{
		http: &http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		},
	}, nil
}

type CatFact struct {
	Fact string
}

func (app *App) GetFact(ctx context.Context, _ *v1.GetFactRequest) (*v1.GetFactResponse, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://catfact.ninja/fact", nil)
	if err != nil {
		log.Printf("failed to build request: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp, err := app.http.Do(request)
	if err != nil {
		log.Printf("failed to get catfact url: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	defer func() { _ = resp.Body.Close() }()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	catFact := new(CatFact)
	err = json.Unmarshal(respBytes, catFact)
	if err != nil {
		log.Printf("failed to unmarshal cat fact: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &v1.GetFactResponse{
		Fact: catFact.Fact,
	}, nil
}
