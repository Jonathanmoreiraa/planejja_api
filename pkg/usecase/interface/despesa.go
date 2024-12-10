package interfaces

import (
	"context"
	domain "github/jonathanmoreiraa/planejja/pkg/domain"
)

type DespesaUseCase interface {
	FindAll(ctx context.Context) ([]domain.Despesas, error)
	FindByID(ctx context.Context, id uint) (domain.Despesas, error)
	Save(ctx context.Context, user domain.Despesas) (domain.Despesas, error)
	Delete(ctx context.Context, user domain.Despesas) error
}
