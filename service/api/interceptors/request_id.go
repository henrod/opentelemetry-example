package interceptors

import (
	"context"
	"opentelemetry-example/service/api"

	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/attribute"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func RequestID(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	requestID := uuid.New().String()
	ctx = context.WithValue(ctx, api.RequestIDKey, requestID)

	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return handler(ctx, req)
	}
	defer span.End()

	span.SetAttributes(attribute.String(api.RequestIDKey, requestID))

	return handler(ctx, req)
}
