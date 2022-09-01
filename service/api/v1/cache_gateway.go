package v1

import (
	"context"
	"errors"
)

var ErrFactNotFound = errors.New("fact not found")

type CacheGateway interface {
	GetFact(ctx context.Context) (string, error)
	SetFact(ctx context.Context, fact string) error
}
