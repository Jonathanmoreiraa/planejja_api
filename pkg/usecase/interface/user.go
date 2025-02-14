package interfaces

import (
	"context"
	domain "github/jonathanmoreiraa/planejja/pkg/domain"
)

type UserUseCase interface {
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	FindByEmail(ctx context.Context, email string) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Update(ctx context.Context, user domain.Users) error
	Delete(ctx context.Context, user domain.Users) error
	Login(ctx context.Context, email string, password string) (map[string]any, error)
}
