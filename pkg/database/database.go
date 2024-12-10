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

	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Receitas{})
	db.AutoMigrate(&domain.Reservas{})
	db.AutoMigrate(&domain.Despesas{})
	db.AutoMigrate(&domain.DespesasParcelas{})
	db.AutoMigrate(&domain.Categorias{})
	db.AutoMigrate(&domain.CategoriasDespesas{})
	db.AutoMigrate(&domain.CategoriasReceitas{})

	return db, dbErr
}
