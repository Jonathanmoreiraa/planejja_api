package revenue

import (
	"context"

	error_message "github.com/jonathanmoreiraa/planejja/internal/domain/error"
	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	"github.com/jonathanmoreiraa/planejja/internal/domain/repository"
	services "github.com/jonathanmoreiraa/planejja/internal/usecase/revenue/contract"
	"github.com/jonathanmoreiraa/planejja/pkg/util"
)

type revenueUseCase struct {
	revenueRepo repository.RevenueRepository
}

func NewRevenueUseCase(repo repository.RevenueRepository) services.RevenueUseCase {
	return &revenueUseCase{
		revenueRepo: repo,
	}
}

func (useCase *revenueUseCase) Create(ctx context.Context, revenue entity.Revenue) (entity.Revenue, error) {
	revenue, err := useCase.revenueRepo.Create(ctx, revenue)
	if err != nil {
		return entity.Revenue{}, util.ErrorWithMessage(err, error_message.ErrCreateAccount)
	}

	return revenue, nil
}

func (useCase *revenueUseCase) GetAllRevenues(ctx context.Context, userId int) ([]entity.Revenue, error) {
	revenues, err := useCase.revenueRepo.FindAll(ctx, userId)
	if err != nil {
		return []entity.Revenue{}, util.ErrorWithMessage(err, error_message.ErrFindRevenue)
	}

	return revenues, nil
}

func (useCase *revenueUseCase) GetRevenue(ctx context.Context, id int) (entity.Revenue, error) {
	revenue, err := useCase.revenueRepo.FindByID(ctx, id)
	if err != nil {
		return entity.Revenue{}, util.ErrorWithMessage(err, error_message.ErrFindRevenue)
	}

	return revenue, nil
}

func (useCase *revenueUseCase) GetRevenues(ctx context.Context, filters map[string]any) ([]entity.Revenue, error) {
	revenue, err := useCase.revenueRepo.FindByFilter(ctx, filters)
	if err != nil {
		return []entity.Revenue{}, util.ErrorWithMessage(err, error_message.ErrFindRevenue)
	}

	return revenue, nil
}

func (useCase *revenueUseCase) Update(ctx context.Context, revenue entity.Revenue) error {
	err := useCase.revenueRepo.Update(ctx, revenue)
	if err != nil {
		return util.ErrorWithMessage(err, error_message.ErrUpdateRevenue)
	}

	return nil
}

func (useCase *revenueUseCase) Delete(ctx context.Context, revenue entity.Revenue) error {
	err := useCase.revenueRepo.Delete(ctx, revenue)
	if err != nil {
		return util.ErrorWithMessage(err, error_message.ErrUpdateRevenue)
	}

	return nil
}
