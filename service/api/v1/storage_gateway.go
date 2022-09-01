package v1

import (
	"context"
	"errors"
	"opentelemetry-example/domain/entities"
)

var ErrCatAlreadyExists = errors.New("cat already exists")

type StorageGateway interface {
	CreateCat(context.Context, *entities.Cat) error
	GetCat(context.Context, string) (*entities.Cat, error)
	ListCats(context.Context) ([]*entities.Cat, error)
}
