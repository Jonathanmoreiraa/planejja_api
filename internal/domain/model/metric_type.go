package model

import (
	"time"

	"gorm.io/gorm"
)

type InvestimentType struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time      `json:"created" gorm:"not null;"`
	UpdatedAt time.Time      `json:"modified" gorm:"not null;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
