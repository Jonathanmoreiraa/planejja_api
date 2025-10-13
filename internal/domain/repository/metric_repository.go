package repository

import (
	"context"

	entity "github.com/jonathanmoreiraa/2cents/internal/domain/model"
)

type MetricRepository interface {
	Create(ctx context.Context, metric entity.Metric) (entity.Metric, error)
	GetInvestimentsTypes(ctx context.Context) ([]entity.InvestimentType, error)
	GetLastMetric(ctx context.Context, investimentType int) (entity.Metric, error)
	Update(ctx context.Context, metric entity.Metric) error
	Delete(ctx context.Context, metric entity.Metric) error
}
