// wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/jonathanmoreiraa/planejja/internal/api/handler"
	"github.com/jonathanmoreiraa/planejja/internal/api/route"
	"github.com/jonathanmoreiraa/planejja/internal/config"
	"github.com/jonathanmoreiraa/planejja/internal/infra/database"
	"github.com/jonathanmoreiraa/planejja/internal/infra/repository"
	"github.com/jonathanmoreiraa/planejja/internal/usecase/revenue"
	"github.com/jonathanmoreiraa/planejja/internal/usecase/user"

	"github.com/google/wire"
)

func NewHandlerGroup(userHandler *handler.UserHandler, revenueHandler *handler.RevenueHandler) route.HandlerGroup {
	return route.HandlerGroup{
		UserHandler:    userHandler,
		RevenueHandler: revenueHandler,
	}
}

func InitializeAPI(cfg config.Config) (*route.ServerHTTP, error) {
	wire.Build(
		database.NewMySqlDatabase,
		repository.NewUserRepository,
		repository.NewRevenueRepository,
		user.NewUserUseCase,
		revenue.NewRevenueUseCase,
		handler.NewUserHandler,
		handler.NewRevenueHandler,
		NewHandlerGroup,
		route.NewServerHTTP,
	)
	return &route.ServerHTTP{}, nil
}
