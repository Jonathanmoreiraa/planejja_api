package usecase

import (
	"context"
	"github/jonathanmoreiraa/planejja/pkg/domain"
	interfaces "github/jonathanmoreiraa/planejja/pkg/repository/interface"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"
)

type receitaUseCase struct {
	receitaRepo interfaces.ReceitaRepository
}

func NewReceitaUseCase(repo interfaces.ReceitaRepository) services.ReceitaUseCase {
	return &receitaUseCase{
		receitaRepo: repo,
	}
}

func (c *receitaUseCase) FindAll(ctx context.Context) ([]domain.Receitas, error) {
	receitas, err := c.receitaRepo.FindAll(ctx)
	return receitas, err
}

func (c *receitaUseCase) FindByID(ctx context.Context, id uint) (domain.Receitas, error) {
	receita, err := c.receitaRepo.FindByID(ctx, id)
	return receita, err
}

func (c *receitaUseCase) Save(ctx context.Context, receita domain.Receitas) (domain.Receitas, error) {
	receita, err := c.receitaRepo.Save(ctx, receita)

	return receita, err
}

func (c *receitaUseCase) Delete(ctx context.Context, receita domain.Receitas) error {
	err := c.receitaRepo.Delete(ctx, receita)

	return err
}
