package contract

import (
	"context"

	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
)

type ExpenseUseCase interface {
	Create(ctx context.Context, revenue entity.Expense) (entity.Expense, error)
	GetAllExpenses(ctx context.Context, userId int) ([]entity.Expense, error)
	GetExpense(ctx context.Context, id int) (entity.Expense, error)
	GetExpenses(ctx context.Context, filters map[string]any) ([]entity.Expense, error)
	Update(ctx context.Context, revenue entity.Expense) error
	Delete(ctx context.Context, revenue entity.Expense) error
}
