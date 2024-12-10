package repository

import (
	"context"
	domain "github/jonathanmoreiraa/planejja/pkg/domain"
	interfaces "github/jonathanmoreiraa/planejja/pkg/repository/interface"

	"gorm.io/gorm"
)

type despesaDatabase struct {
	DB *gorm.DB
}

func NewDespesaRepository(DB *gorm.DB) interfaces.DespesaRepository {
	return &despesaDatabase{DB}
}

func (c *despesaDatabase) FindAll(ctx context.Context) ([]domain.Despesas, error) {
	var despesas []domain.Despesas
	err := c.DB.Find(&despesas).Error

	return despesas, err
}

func (c *despesaDatabase) FindByID(ctx context.Context, id uint) (domain.Despesas, error) {
	var despesa domain.Despesas
	err := c.DB.First(&despesa, id).Error

	return despesa, err
}

func (c *despesaDatabase) Save(ctx context.Context, despesa domain.Despesas) (domain.Despesas, error) {
	err := c.DB.Save(&despesa).Error

	return despesa, err
}

func (c *despesaDatabase) Delete(ctx context.Context, despesa domain.Despesas) error {
	err := c.DB.Delete(&despesa).Error

	return err
}
