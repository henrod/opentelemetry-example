package main

import (
	"fmt"
	"log"
	"net"
	v1 "opentelemetry-example/protogen/go/api/v1"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcServerURL = "127.0.0.1:8080"
)

func run() error {
	listener, err := net.Listen("tcp", grpcServerURL)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", grpcServerURL, err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)
	reflection.Register(server)

	app, err := NewApp()
	if err != nil {
		return fmt.Errorf("failed to create app: %w", err)
	}

	v1.RegisterCatServiceServer(server, app)

	log.Printf("running gRPC server at %s", grpcServerURL)
	if err = server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
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
