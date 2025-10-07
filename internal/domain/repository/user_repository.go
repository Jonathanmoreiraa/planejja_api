package repository

import (
	"context"

	entity "github.com/jonathanmoreiraa/2cents/internal/domain/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int) (entity.User, error)
	FindByColumn(ctx context.Context, column string, data any) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, user entity.User) error
}
