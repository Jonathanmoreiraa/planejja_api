package model

import (
	"time"

	"gorm.io/gorm"
)

type Revenue struct {
	ID          int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      int            `json:"user_id,omitempty" gorm:"not null"`
	User        User           `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	Description string         `json:"description" gorm:"not null;type:varchar(255)"`
	DueDate     *time.Time     `json:"due_date" gorm:"type:date"`
	Value       string         `json:"value" gorm:"not null;type:varchar(255)"`
	Received    int            `json:"received" gorm:"tinyint(1);default:0"`
	CreatedAt   time.Time      `json:"created" gorm:"not null;"`
	UpdatedAt   time.Time      `json:"modified" gorm:"not null;"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
