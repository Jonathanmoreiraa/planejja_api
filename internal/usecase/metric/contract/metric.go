package contract

import (
	"context"

	entity "github.com/jonathanmoreiraa/2cents/internal/domain/model"
)

type MetricUseCase interface {
	Create(ctx context.Context, metric entity.Metric) (entity.Metric, error)
	GetLastMetric(ctx context.Context, investimentType int) (entity.Metric, error)
	// SimulateInvestments(ctx context.Context, initialValue float64, monthlyContribution float64, months int) ([]string, error)
}
