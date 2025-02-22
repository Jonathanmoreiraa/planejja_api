// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github/jonathanmoreiraa/planejja/pkg/api/handler"
	"github/jonathanmoreiraa/planejja/pkg/config"
	"github/jonathanmoreiraa/planejja/pkg/database"
	"github/jonathanmoreiraa/planejja/pkg/repository"
	"github/jonathanmoreiraa/planejja/pkg/routes"
	"github/jonathanmoreiraa/planejja/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
