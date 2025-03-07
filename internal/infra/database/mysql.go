package database

import (
	"fmt"
	"log"

	config "github.com/jonathanmoreiraa/planejja/internal/config"
	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	database "github.com/jonathanmoreiraa/planejja/internal/infra/database/interface"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLProvider struct {
	Db *gorm.DB
}

func NewMySqlDatabase(cfg config.Config) (database.DatabaseProvider, error) {
	mysqlInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=America%sSao_Paulo",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, "%2f",
	)

	db, dbErr := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
		//Logger:                 logger.Default.LogMode(logger.Info),
	})
	if dbErr != nil {
		log.Fatal("cannot load database: ", dbErr)
	}

	err := db.AutoMigrate(&entity.User{}, &entity.Revenue{})
	if err != nil {
		log.Fatal("erro na migração: ", err)
	}

	dbInstance := &MySQLProvider{Db: db}

	return dbInstance, dbErr
}

func (mysql MySQLProvider) GetDatabase() *gorm.DB {
	return mysql.Db
}
