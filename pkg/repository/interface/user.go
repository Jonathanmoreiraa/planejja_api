package interfaces

import (
	"context"
	"github/jonathanmoreiraa/planejja/pkg/domain"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	FindByEmail(ctx context.Context, email string) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Update(ctx context.Context, user domain.Users) error
	Delete(ctx context.Context, user domain.Users) error
}
