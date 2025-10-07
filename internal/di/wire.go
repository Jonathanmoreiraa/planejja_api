// wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/jonathanmoreiraa/2cents/internal/api/handler"
	"github.com/jonathanmoreiraa/2cents/internal/api/route"
	"github.com/jonathanmoreiraa/2cents/internal/config"
	"github.com/jonathanmoreiraa/2cents/internal/infra/database"
	"github.com/jonathanmoreiraa/2cents/internal/infra/repository"
	"github.com/jonathanmoreiraa/2cents/internal/usecase/category"
	"github.com/jonathanmoreiraa/2cents/internal/usecase/expense"
	"github.com/jonathanmoreiraa/2cents/internal/usecase/revenue"
	"github.com/jonathanmoreiraa/2cents/internal/usecase/user"

	"github.com/google/wire"
)

func NewHandlerGroup(
	userHandler *handler.UserHandler,
	revenueHandler *handler.RevenueHandler,
	expenseHandler *handler.ExpenseHandler,
	categoryHandler *handler.CategoryHandler,
) route.HandlerGroup {
	return route.HandlerGroup{
		UserHandler:     userHandler,
		RevenueHandler:  revenueHandler,
		ExpenseHandler:  expenseHandler,
		CategoryHandler: categoryHandler,
	}
}

func InitializeAPI(cfg config.Config) (*route.ServerHTTP, error) {
	wire.Build(
		database.NewMySqlDatabase,

		repository.NewUserRepository,
		repository.NewRevenueRepository,
		repository.NewExpenseRepository,
		repository.NewCategoryRepository,

		user.NewUserUseCase,
		revenue.NewRevenueUseCase,
		expense.NewExpenseUseCase,
		category.NewCategoryUseCase,

		handler.NewUserHandler,
		handler.NewRevenueHandler,
		handler.NewExpenseHandler,
		handler.NewCategoryHandler,

		NewHandlerGroup,
		route.NewServerHTTP,
	)
	return &route.ServerHTTP{}, nil
}
