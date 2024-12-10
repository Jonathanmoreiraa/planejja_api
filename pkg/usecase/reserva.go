package usecase

import (
	"context"
	"github/jonathanmoreiraa/planejja/pkg/domain"
	interfaces "github/jonathanmoreiraa/planejja/pkg/repository/interface"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"
)

type reservaUseCase struct {
	reservaRepo interfaces.ReservaRepository
}

func NewReservaUseCase(repo interfaces.ReservaRepository) services.ReservaUseCase {
	return &reservaUseCase{
		reservaRepo: repo,
	}
}

func (c *reservaUseCase) FindAll(ctx context.Context) ([]domain.Reservas, error) {
	reservas, err := c.reservaRepo.FindAll(ctx)
	return reservas, err
}

func (c *reservaUseCase) FindByID(ctx context.Context, id uint) (domain.Reservas, error) {
	reserva, err := c.reservaRepo.FindByID(ctx, id)
	return reserva, err
}

func (c *reservaUseCase) Save(ctx context.Context, reserva domain.Reservas) (domain.Reservas, error) {
	reserva, err := c.reservaRepo.Save(ctx, reserva)

	return reserva, err
}

func (c *reservaUseCase) Delete(ctx context.Context, reserva domain.Reservas) error {
	err := c.reservaRepo.Delete(ctx, reserva)

	return err
}
