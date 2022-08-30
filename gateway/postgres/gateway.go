package postgres

import (
	"context"
	"fmt"
	"opentelemetry-example/domain/entities"

	"github.com/go-pg/pg/extra/pgotel/v10"
	"github.com/go-pg/pg/v10"
)

type Gateway struct {
	postgres *pg.DB
}

func NewGateway(postgresURL string) (*Gateway, error) {
	options, err := pg.ParseURL(postgresURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres URL: %w", err)
	}

	db := pg.Connect(options)
	db.AddQueryHook(pgotel.NewTracingHook())

	return &Gateway{
		postgres: db,
	}, nil
}

func getPostgresClient(options *pg.Options) *pg.DB {
	db := pg.Connect(options)

	// Add tracer
	db.AddQueryHook(pgotel.NewTracingHook())

	return db
}

func (gateway *Gateway) CreateCat(ctx context.Context, cat *entities.Cat) error {
	_, err := gateway.postgres.ModelContext(ctx, NewCat(cat)).Insert()
	if err != nil {
		return fmt.Errorf("failed to insert cat into postgres: %w", err)
	}

	return nil
}

func (gateway *Gateway) GetCat(ctx context.Context, id string) (*entities.Cat, error) {
	cat := &Cat{ID: id}
	err := gateway.postgres.ModelContext(ctx, cat).WherePK().Select()
	if err != nil {
		return nil, fmt.Errorf("failed to select cat from postgres: %w", err)
	}

	return cat.ToEntity(), nil
}

func (gateway *Gateway) ListCats(ctx context.Context) ([]*entities.Cat, error) {
	var cats []*Cat
	err := gateway.postgres.ModelContext(ctx, &cats).Select()
	if err != nil {
		return nil, fmt.Errorf("failed to select many cats from postgres: %w", err)
	}

	eCats := make([]*entities.Cat, 0, len(cats))
	for _, cat := range cats {
		eCats = append(eCats, cat.ToEntity())
	}

	return eCats, nil
}
