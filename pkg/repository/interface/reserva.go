package interfaces

import (
	"context"
	"github/jonathanmoreiraa/planejja/pkg/domain"
)

type ReservaRepository interface {
	FindAll(ctx context.Context) ([]domain.Reservas, error)
	FindByID(ctx context.Context, id uint) (domain.Reservas, error)
	Save(ctx context.Context, user domain.Reservas) (domain.Reservas, error)
	Delete(ctx context.Context, user domain.Reservas) error
}
