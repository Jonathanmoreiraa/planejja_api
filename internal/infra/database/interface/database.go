package database

import "gorm.io/gorm"

type DatabaseProvider interface {
	GetDatabase() *gorm.DB
}
