package database

import (
	config "github/jonathanmoreiraa/planejja/pkg/config"
	domain "github/jonathanmoreiraa/planejja/pkg/domain"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	mysqlInfo := config.Config.GetDSN(cfg)
	db, dbErr := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		log.Fatal("cannot load database: ", dbErr)
	}

	err := db.AutoMigrate(&domain.Users{})
	if err != nil {
		log.Fatal("erro na migração: ", err)
	}

	return db, dbErr
}
