package repository

import (
	"context"
	"time"

	entity "github.com/jonathanmoreiraa/2cents/internal/domain/model"
	interfaces "github.com/jonathanmoreiraa/2cents/internal/domain/repository"
	database "github.com/jonathanmoreiraa/2cents/internal/infra/database/interface"

	"gorm.io/gorm"
)

type categoryDatabase struct {
	DB *gorm.DB
}

func NewCategoryRepository(Database database.DatabaseProvider) interfaces.CategoryRepository {
	return &categoryDatabase{DB: Database.GetDatabase()}
}

func (database *categoryDatabase) Create(ctx context.Context, category entity.Category) (entity.Category, error) {
	err := database.DB.Create(&category).Error
	return category, err
}

func (database *categoryDatabase) FindAll(ctx context.Context, userId int) ([]entity.Category, error) {
	var categories []entity.Category

	err := database.DB.
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		Find(&categories).Error
	return categories, err
}

func (database *categoryDatabase) FindByName(ctx context.Context, name string, userId int) ([]entity.Category, error) {
	var categories []entity.Category

	query := database.DB.
		Where("user_id = ?", userId).
		Where("name LIKE ?", "%"+name+"%").
		Where("deleted_at IS NULL")

	err := query.Find(&categories).Error

	return categories, err
}

func (database *categoryDatabase) FindById(ctx context.Context, id int, userId int) (entity.Category, error) {
	var category entity.Category

	query := database.DB.
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL")

	err := query.Find(&category, id).Error

	return category, err
}

func (database *categoryDatabase) Update(ctx context.Context, category entity.Category) error {
	err := database.DB.Model(&category).Updates(map[string]interface{}{
		"name":       category.Name,
		"updated_at": time.Now(),
	}).Error
	return err
}

func (database *categoryDatabase) Delete(ctx context.Context, category entity.Category) error {
	err := database.DB.Delete(&category).Error
	return err
}
