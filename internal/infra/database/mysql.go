package database

import (
	"fmt"

	config "github.com/jonathanmoreiraa/2cents/internal/config"
	entity "github.com/jonathanmoreiraa/2cents/internal/domain/model"
	database "github.com/jonathanmoreiraa/2cents/internal/infra/database/interface"
	"github.com/jonathanmoreiraa/2cents/internal/infra/seed"
	"github.com/jonathanmoreiraa/2cents/pkg/log"

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
		log.NewLogger().Error(dbErr)
	}

	err := db.AutoMigrate(
		&entity.User{},
		&entity.Revenue{},
		&entity.Expense{},
		&entity.Category{},
		&entity.Saving{},
		&entity.InvestimentType{},
		&entity.Metric{},
	)
	if err != nil {
		log.NewLogger().Error(err)
	}

	dbInstance := &MySQLProvider{Db: db}
	seed.Run(dbInstance.Db)

	return dbInstance, dbErr
}

func (mysql MySQLProvider) GetDatabase() *gorm.DB {
	return mysql.Db
}
