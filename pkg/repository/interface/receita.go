package interfaces

import (
	"context"
	"github/jonathanmoreiraa/planejja/pkg/domain"
)

type ReceitaRepository interface {
	FindAll(ctx context.Context) ([]domain.Receitas, error)
	FindByID(ctx context.Context, id uint) (domain.Receitas, error)
	Save(ctx context.Context, user domain.Receitas) (domain.Receitas, error)
	Delete(ctx context.Context, user domain.Receitas) error
}
