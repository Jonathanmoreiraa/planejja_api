// wire.go
//go:build wireinject
// +build wireinject

package di

import (
	handler "github/jonathanmoreiraa/planejja/pkg/api/handler"
	config "github/jonathanmoreiraa/planejja/pkg/config"
	database "github/jonathanmoreiraa/planejja/pkg/database"
	repository "github/jonathanmoreiraa/planejja/pkg/repository"
	http "github/jonathanmoreiraa/planejja/pkg/routes"
	"github/jonathanmoreiraa/planejja/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		database.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
