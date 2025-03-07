package repository

import (
	"context"

	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	interfaces "github.com/jonathanmoreiraa/planejja/internal/domain/repository"
	database "github.com/jonathanmoreiraa/planejja/internal/infra/database/interface"

	"gorm.io/gorm"
)

type revenueDatabase struct {
	DB *gorm.DB
}

func NewRevenueRepository(Database database.DatabaseProvider) interfaces.RevenueRepository {
	return &revenueDatabase{DB: Database.GetDatabase()}
}

func (database *revenueDatabase) FindByID(ctx context.Context, id int) (entity.Revenue, error) {
	var revenue entity.Revenue

	err := database.DB.First(&revenue, id).Error
	return revenue, err
}

func (database *revenueDatabase) FindAll(ctx context.Context) (entity.Revenue, error) {
	var revenue entity.Revenue

	err := database.DB.Find(&revenue).Error
	return revenue, err
}

func (database *revenueDatabase) Create(ctx context.Context, revenue entity.Revenue) (entity.Revenue, error) {
	err := database.DB.Create(&revenue).Error
	return revenue, err
}

// func (database *revenueDatabase) Update(ctx context.Context, revenue entity.Revenue) error {
// 	err := database.DB.entity(&revenue).Updates(map[string]interface{}{
// 		"name":       revenue.Name,
// 		"password":   revenue.Password,
// 		"birth_date": revenue.BirthDate,
// 		"modified":   time.Now(),
// 	}).Error
// 	return err
// }

func (database *revenueDatabase) Delete(ctx context.Context, revenue entity.Revenue) error {
	err := database.DB.Delete(&revenue).Error
	return err
}
