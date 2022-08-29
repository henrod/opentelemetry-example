package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type App struct {
	http *http.Client
}

func NewApp() (*App, error) {
	return &App{
		http: newHTTPClient(),
	}, nil
}

type CatFact struct {
	Fact string
}

func (app *App) GetCatFact(ctx context.Context) (string, error) {
	resp, err := app.http.Get("https://catfact.ninja/fact")
	if err != nil {
		return "", fmt.Errorf("failed to get catfact url: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	catFact := new(CatFact)
	err = json.Unmarshal(respBytes, catFact)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal cat fact: %w", err)
	}

	return catFact.Fact, nil
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
}
