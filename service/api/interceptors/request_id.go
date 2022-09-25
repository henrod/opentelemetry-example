package interceptors

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/attribute"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const requestIDKey = "request.id"

func RequestID(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	requestID := uuid.New().String()
	ctx = context.WithValue(ctx, requestIDKey, requestID)

	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return handler(ctx, req)
	}
	defer span.End()

	span.SetAttributes(attribute.String(requestIDKey, requestID))

	return handler(ctx, req)
}
