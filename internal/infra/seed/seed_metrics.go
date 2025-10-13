package seed

import (
	"time"

	"github.com/jonathanmoreiraa/2cents/internal/domain/model"
	"github.com/jonathanmoreiraa/2cents/pkg/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func SeedMetrics(db *gorm.DB) {
	var count int64
	db.Model(&model.Metric{}).Count(&count)
	if count > 0 {
		return
	}

	metrics := []model.Metric{
		{InvestimentTypeID: 1, Date: time.Date(2025, 10, 9, 3, 0, 0, 0, time.UTC), Value: decimal.NewFromFloat(0.055131)},
		{InvestimentTypeID: 2, Date: time.Date(2025, 10, 1, 3, 0, 0, 0, time.UTC), Value: decimal.NewFromFloat(14.90)},
	}
	if err := db.Create(&metrics).Error; err != nil {
		log.NewLogger().Error(err)
	}
}
