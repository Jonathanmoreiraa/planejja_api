package repository

import (
	"context"
	domain "github/jonathanmoreiraa/planejja/pkg/domain"
	interfaces "github/jonathanmoreiraa/planejja/pkg/repository/interface"

	"gorm.io/gorm"
)

type receitaDatabase struct {
	DB *gorm.DB
}

func NewReceitaRepository(DB *gorm.DB) interfaces.ReceitaRepository {
	return &receitaDatabase{DB}
}

func (c *receitaDatabase) FindAll(ctx context.Context) ([]domain.Receitas, error) {
	var receitas []domain.Receitas
	err := c.DB.Find(&receitas).Error

	return receitas, err
}

func (c *receitaDatabase) FindByID(ctx context.Context, id uint) (domain.Receitas, error) {
	var receita domain.Receitas
	err := c.DB.First(&receita, id).Error

	return receita, err
}

func (c *receitaDatabase) Save(ctx context.Context, receita domain.Receitas) (domain.Receitas, error) {
	err := c.DB.Save(&receita).Error

	return receita, err
}

func (c *receitaDatabase) Delete(ctx context.Context, receita domain.Receitas) error {
	err := c.DB.Delete(&receita).Error

	return err
}
