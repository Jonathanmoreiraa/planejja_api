package usecase

import (
	"context"
	"github/jonathanmoreiraa/planejja/pkg/domain"
	interfaces "github/jonathanmoreiraa/planejja/pkg/repository/interface"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"
)

type despesaUseCase struct {
	despesaRepo interfaces.DespesaRepository
}

func NewDespesaUseCase(repo interfaces.DespesaRepository) services.DespesaUseCase {
	return &despesaUseCase{
		despesaRepo: repo,
	}
}

func (c *despesaUseCase) FindAll(ctx context.Context) ([]domain.Despesas, error) {
	despesas, err := c.despesaRepo.FindAll(ctx)
	return despesas, err
}

func (c *despesaUseCase) FindByID(ctx context.Context, id uint) (domain.Despesas, error) {
	despesa, err := c.despesaRepo.FindByID(ctx, id)
	return despesa, err
}

func (c *despesaUseCase) Save(ctx context.Context, despesa domain.Despesas) (domain.Despesas, error) {
	despesa, err := c.despesaRepo.Save(ctx, despesa)

	return despesa, err
}

func (c *despesaUseCase) Delete(ctx context.Context, despesa domain.Despesas) error {
	err := c.despesaRepo.Delete(ctx, despesa)

	return err
}
