package repository

import (
	"context"
	domain "github/jonathanmoreiraa/planejja/pkg/domain"
	interfaces "github/jonathanmoreiraa/planejja/pkg/repository/interface"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (database *userDatabase) FindAll(ctx context.Context) ([]domain.Users, error) {
	var users []domain.Users
	err := database.DB.Find(&users).Error

	return users, err
}

func (database *userDatabase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	var user domain.Users
	err := database.DB.First(&user, id).Error

	return user, err
}

func (database *userDatabase) FindByEmail(ctx context.Context, email string) (domain.Users, error) {
	var user domain.Users
	err := database.DB.First(&user, "email = ?", email).Error

	return user, err
}

func (database *userDatabase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := database.DB.Save(&user).Error

	return user, err
}

func (database *userDatabase) Update(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := database.DB.Save(&user).Error

	return user, err
}

func (database *userDatabase) Delete(ctx context.Context, user domain.Users) error {
	err := database.DB.Delete(&user).Error

	return err
}
