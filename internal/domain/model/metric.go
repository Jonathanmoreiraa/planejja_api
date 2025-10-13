package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Metric struct {
	ID                int             `json:"id" gorm:"primaryKey;autoIncrement"`
	InvestimentTypeID int             `json:"investiment_type_id,omitempty" gorm:"not null"`
	InvestimentType   InvestimentType `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	Date              time.Time       `json:"date" gorm:"not null;type:date"`
	Value             decimal.Decimal `json:"value" gorm:"not null;type:decimal(10,8)"`
	CreatedAt         time.Time       `json:"created" gorm:"not null;"`
	UpdatedAt         time.Time       `json:"modified" gorm:"not null;"`
	DeletedAt         gorm.DeletedAt  `gorm:"index"`
}
