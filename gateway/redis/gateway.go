package redis

import (
	"context"
	"errors"
	"fmt"
	v1 "opentelemetry-example/service/api/v1"
	"strings"
	"time"

	"github.com/go-redis/redis/extra/redisotel/v8"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"

	"github.com/go-redis/redis/v8"
)

type Gateway struct {
	redis      *redis.Client
	expiration time.Duration
}

func NewGateway(redisURL string) (*Gateway, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}

	redisClient := getRedisClient(options)

	return &Gateway{
		redis:      redisClient,
		expiration: 10 * time.Second,
	}, nil
}

func getRedisClient(options *redis.Options) *redis.Client {
	redisClient := redis.NewClient(options)

	hostPort := strings.Split(options.Addr, ":")

	// Add tracer
	redisClient.AddHook(redisotel.NewTracingHook(redisotel.WithAttributes(
		semconv.NetPeerNameKey.String(hostPort[0]),
		semconv.NetPeerPortKey.String(hostPort[1])),
	))

	return redisClient
}

func (gateway *Gateway) GetFact(ctx context.Context) (string, error) {
	fact, err := gateway.redis.Get(ctx, gateway.factKey()).Result()
	if errors.Is(err, redis.Nil) {
		return "", v1.ErrFactNotFound
	}

	if err != nil {
		return "", fmt.Errorf("failed to get fact from redis: %w", err)
	}

	return fact, nil
}

func (gateway *Gateway) SetFact(ctx context.Context, fact string) error {
	err := gateway.redis.SetEX(ctx, gateway.factKey(), fact, gateway.expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set fact on redis: %w", err)
	}

	return nil
}

func (gateway *Gateway) factKey() string {
	return "fact:last"
}
