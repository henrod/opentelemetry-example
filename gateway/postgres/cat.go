package postgres

import "opentelemetry-example/domain/entities"

type Cat struct {
	ID   string `pg:",pk"`
	Name string
}

func NewCat(cat *entities.Cat) *Cat {
	return &Cat{
		ID:   cat.ID,
		Name: cat.Name,
	}
}

func (cat *Cat) ToEntity() *entities.Cat {
	return &entities.Cat{
		ID:   cat.ID,
		Name: cat.Name,
	}
}
