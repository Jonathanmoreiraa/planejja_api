package main

import (
	"fmt"

	"github.com/jonathanmoreiraa/2cents/internal/config"
	"github.com/jonathanmoreiraa/2cents/internal/infra/database"
	"github.com/jonathanmoreiraa/2cents/internal/infra/repository"
	robot "github.com/jonathanmoreiraa/2cents/internal/robot/metrics"
	"github.com/jonathanmoreiraa/2cents/pkg/log"
)

// TODO: Utilizar o crontab para rodar o robô diariamente
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.NewLogger().Error(err)
	}

	fmt.Println(&cfg.DbUser, &cfg.DbPassword, &cfg.DbHost, &cfg.DbPort, &cfg.DbName)

	databaseProvider, err := database.NewMySqlDatabase(*cfg)
	if err != nil {
		log.NewLogger().Error(err)
	}
	metricRepository := repository.NewMetricRepository(databaseProvider)

	robot.RunDaily(metricRepository)
}
