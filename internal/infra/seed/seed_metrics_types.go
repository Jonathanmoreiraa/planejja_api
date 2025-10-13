package seed

import (
	"github.com/jonathanmoreiraa/2cents/internal/domain/model"
	"github.com/jonathanmoreiraa/2cents/pkg/log"
	"gorm.io/gorm"
)

func SeedInvestimentTypes(db *gorm.DB) {
	var count int64
	db.Model(&model.InvestimentType{}).Count(&count)
	if count > 0 {
		return
	}

	investimentTypes := []model.InvestimentType{
		{Name: "CDI"},
		{Name: "Taxa selic"},
	}

	if err := db.Create(&investimentTypes).Error; err != nil {
		log.NewLogger().Error(err)
	}
}
