package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int            `json:"user_id,omitempty" gorm:"not null"`
	User      User           `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	Name      string         `json:"Name" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time      `json:"created" gorm:"not null;"`
	UpdatedAt time.Time      `json:"modified" gorm:"not null;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
