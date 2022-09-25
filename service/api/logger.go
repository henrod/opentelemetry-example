package api

import (
	"context"

	"go.uber.org/zap"
)

const RequestIDKey = "request.id"

func Logger(ctx context.Context) *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	requestID, ok := ctx.Value(RequestIDKey).(string)
	if ok {
		sugar = sugar.With("request_id", requestID)
	}

	return sugar
}
