package v1

import (
	"context"
	"opentelemetry-example/domain/entities"
)

type StorageGateway interface {
	CreateCat(context.Context, *entities.Cat) error
	GetCat(context.Context, string) (*entities.Cat, error)
	ListCats(context.Context) ([]*entities.Cat, error)
}
