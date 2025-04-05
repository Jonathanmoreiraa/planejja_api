package contract

import (
	"context"

	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
)

type CategoryUseCase interface {
	Create(ctx context.Context, revenue entity.Category) (entity.Category, error)
	GetAllCategories(ctx context.Context, userId int) ([]entity.Category, error)
	GetCategory(ctx context.Context, name string, userId int) ([]entity.Category, error)
	GetCategoryById(ctx context.Context, id int, userId int) (entity.Category, error)

	//TODO: ver uma forma de lidar com o update e delete
	// Update(ctx context.Context, revenue entity.Category) error
	// Delete(ctx context.Context, revenue entity.Category) error
}
