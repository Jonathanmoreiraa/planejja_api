package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"not null;type:varchar(255)"`
	Email     string         `json:"email" gorm:"not null;type:varchar(255);unique"`
	Password  string         `json:"password" gorm:"not null;type:varchar(255)"`
	BirthDate *time.Time     `json:"birth_date" gorm:"type:date"`
	CreatedAt time.Time      `gorm:"not null;"`
	UpdatedAt time.Time      `gorm:"not null;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
