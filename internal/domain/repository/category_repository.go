package repository

import (
	"context"

	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
)

type CategoryRepository interface {
	Create(ctx context.Context, revenue entity.Category) (entity.Category, error)
	FindAll(ctx context.Context, userId int) ([]entity.Category, error)
	FindByName(ctx context.Context, name string, userId int) ([]entity.Category, error)
	FindById(ctx context.Context, id int, userId int) (entity.Category, error)
	Update(ctx context.Context, revenue entity.Category) error
	Delete(ctx context.Context, revenue entity.Category) error
}
