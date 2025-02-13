package domain

import "time"

type Users struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"not null;type:varchar(255)"`
	Email     string     `json:"email" gorm:"not null;type:varchar(255);unique"`
	Password  string     `json:"password" gorm:"not null;type:varchar(255)"`
	BirthDate *time.Time `json:"birth_date"`
	Created   time.Time  `json:"created" gorm:"not null;autoCreateTime:true"`
	Modified  time.Time  `json:"modified" gorm:"not null;autoUpdateTime:true"`
}
