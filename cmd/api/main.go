package main

import (
	config "github.com/jonathanmoreiraa/planejja/internal/config"
	di "github.com/jonathanmoreiraa/planejja/internal/di"
	"github.com/jonathanmoreiraa/planejja/pkg/log"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.NewLogger().Error(configErr)
	}

	server, diErr := di.InitializeAPI(*config)
	if diErr != nil {
		log.NewLogger().Error(diErr)
	}

	server.Start()
}
