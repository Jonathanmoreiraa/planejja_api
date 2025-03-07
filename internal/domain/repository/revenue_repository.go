package repository

import (
	"context"

	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
)

type RevenueRepository interface {
	FindByID(ctx context.Context, id int) (entity.Revenue, error)
	FindAll(ctx context.Context) (entity.Revenue, error)
	Create(ctx context.Context, user entity.Revenue) (entity.Revenue, error)
	// Update(ctx context.Context, user entity.Revenue) error
	Delete(ctx context.Context, user entity.Revenue) error
}
