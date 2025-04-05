package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Revenue struct {
	ID          int             `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      int             `json:"user_id,omitempty" gorm:"not null"`
	User        User            `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	Description string          `json:"description" gorm:"not null;type:varchar(255)"`
	DueDate     *time.Time      `json:"due_date" gorm:"type:date"`
	Value       decimal.Decimal `json:"value" gorm:"not null;type:decimal(19,2)"`
	Received    int             `json:"received" gorm:"type:tinyint(1);not null;default:0"`
	CreatedAt   time.Time       `json:"created" gorm:"not null;"`
	UpdatedAt   time.Time       `json:"modified" gorm:"not null;"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
}
