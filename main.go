package main

import (
	"fmt"
	"log"
	"net"
	httpGateway "opentelemetry-example/gateway/http"
	"opentelemetry-example/gateway/postgres"
	"opentelemetry-example/gateway/redis"
	proto "opentelemetry-example/protogen/go/api/v1"
	"opentelemetry-example/service/api/interceptors"
	api "opentelemetry-example/service/api/v1"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcServerURL = "127.0.0.1:8080"
	postgresURL   = "postgresql://postgres:@localhost:5432/postgres?sslmode=disable"
	redisURL      = "redis://localhost:6379"
)

func run() error {
	listener, err := net.Listen("tcp", grpcServerURL)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", grpcServerURL, err)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
			interceptors.RequestID,
		),
	)
	reflection.Register(server)

	catService, err := getCatService()
	if err != nil {
		return fmt.Errorf("failed to build cat service: %w", err)
	}

	proto.RegisterCatServiceServer(server, catService)

	log.Printf("running gRPC server at %s", grpcServerURL)
	if err = server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

func getCatService() (*api.CatService, error) {
	storageGateway, err := postgres.NewGateway(postgresURL)
	if err != nil {
		return nil, fmt.Errorf("failed to build postgres gateway: %w", err)
	}

	httpClient := httpGateway.GetHTTPClient()

	cacheGateway, err := redis.NewGateway(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to build redis gateway: %w", err)
	}

	catService, err := api.NewCatService(httpClient, storageGateway, cacheGateway)
	if err != nil {
		return nil, fmt.Errorf("failed to create app: %w", err)
	}

	return catService, nil
}

func main() {
	closeTracer, err := ConfigureTracer()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = closeTracer()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = run()
	if err != nil {
		log.Fatal(err)
	}
}
