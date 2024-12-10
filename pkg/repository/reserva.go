package repository

import (
	"context"
	domain "github/jonathanmoreiraa/planejja/pkg/domain"
	interfaces "github/jonathanmoreiraa/planejja/pkg/repository/interface"

	"gorm.io/gorm"
)

type reservaDatabase struct {
	DB *gorm.DB
}

func NewReservaRepository(DB *gorm.DB) interfaces.ReservaRepository {
	return &reservaDatabase{DB}
}

func (c *reservaDatabase) FindAll(ctx context.Context) ([]domain.Reservas, error) {
	var reservas []domain.Reservas
	err := c.DB.Find(&reservas).Error

	return reservas, err
}

func (c *reservaDatabase) FindByID(ctx context.Context, id uint) (domain.Reservas, error) {
	var reserva domain.Reservas
	err := c.DB.First(&reserva, id).Error

	return reserva, err
}

func (c *reservaDatabase) Save(ctx context.Context, reserva domain.Reservas) (domain.Reservas, error) {
	err := c.DB.Save(&reserva).Error

	return reserva, err
}

func (c *reservaDatabase) Delete(ctx context.Context, reserva domain.Reservas) error {
	err := c.DB.Delete(&reserva).Error

	return err
}
