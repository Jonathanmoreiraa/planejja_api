package seed

import (
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	SeedInvestimentTypes(db)
	SeedMetrics(db)
}
