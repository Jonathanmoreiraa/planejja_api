package robot

import (
	"github.com/jonathanmoreiraa/2cents/internal/domain/repository"
	"github.com/jonathanmoreiraa/2cents/pkg/log"
)

const (
	CDI_URL   = "https://api.bcb.gov.br/dados/serie/bcdata.sgs.12/dados?formato=json&dataInicial=${startDate}&dataFinal=${today}"
	SELIC_URL = "https://api.bcb.gov.br/dados/serie/bcdata.sgs.4189/dados?formato=json&dataInicial=${startDate}&dataFinal=${today}"
)

func RunDaily(repo repository.MetricRepository) {
	if err := RunRobot(repo, 1, CDI_URL); err != nil {
		log.NewLogger().Error(err)
	}

	if err := RunRobot(repo, 2, SELIC_URL); err != nil {
		log.NewLogger().Error(err)
	}
}
