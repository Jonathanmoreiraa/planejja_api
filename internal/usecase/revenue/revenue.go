package revenue

import (
	"context"
	"errors"

	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	"github.com/jonathanmoreiraa/planejja/internal/domain/repository"
	services "github.com/jonathanmoreiraa/planejja/internal/usecase/revenue/contract"

	"github.com/go-sql-driver/mysql"
)

type revenueUseCase struct {
	revenueRepo repository.RevenueRepository
}

func NewRevenueUseCase(repo repository.RevenueRepository) services.RevenueUseCase {
	return &revenueUseCase{
		revenueRepo: repo,
	}
}

func (useCase *revenueUseCase) FindByID(ctx context.Context, id int) (entity.Revenue, error) {
	revenue, err := useCase.revenueRepo.FindByID(ctx, id)
	return revenue, err
}

func (useCase *revenueUseCase) FindAll(ctx context.Context) (entity.Revenue, error) {
	revenue, err := useCase.revenueRepo.FindAll(ctx)
	return revenue, err
}

func (useCase *revenueUseCase) Create(ctx context.Context, revenue entity.Revenue) (entity.Revenue, error) {
	revenue, err := useCase.revenueRepo.Create(ctx, revenue)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				err = errors.New("o e-mail informado já foi cadastrado anteriormente")
			}
		}
		return entity.Revenue{}, err
	}

	return revenue, nil
}

func (useCase *revenueUseCase) Update(ctx context.Context, revenue entity.Revenue) error {
	_, err := useCase.FindByID(ctx, revenue.ID)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *revenueUseCase) Delete(ctx context.Context, revenue entity.Revenue) error {
	err := useCase.revenueRepo.Delete(ctx, revenue)

	return err
}
