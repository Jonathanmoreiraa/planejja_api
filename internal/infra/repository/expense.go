package repository

import (
	"context"
	"time"

	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	interfaces "github.com/jonathanmoreiraa/planejja/internal/domain/repository"
	database "github.com/jonathanmoreiraa/planejja/internal/infra/database/interface"
	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

type expenseDatabase struct {
	DB *gorm.DB
}

func NewExpenseRepository(Database database.DatabaseProvider) interfaces.ExpenseRepository {
	return &expenseDatabase{DB: Database.GetDatabase()}
}

func (database *expenseDatabase) Create(ctx context.Context, expense entity.Expense) (entity.Expense, error) {
	err := database.DB.Create(&expense).Error
	return expense, err
}

func (database *expenseDatabase) FindAll(ctx context.Context, userId int) ([]entity.Expense, error) {
	var expenses []entity.Expense

	err := database.DB.
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		Find(&expenses).Error
	return expenses, err
}

func (database *expenseDatabase) FindByID(ctx context.Context, id int) (entity.Expense, error) {
	var expense entity.Expense

	err := database.DB.First(&expense, id).Error
	return expense, err
}

func (database *expenseDatabase) FindByFilter(ctx context.Context, filters map[string]any) ([]entity.Expense, error) {
	var expenses []entity.Expense

	query := database.DB.
		Where("user_id = ?", filters["user_id"]).
		Where("deleted_at IS NULL")

	if description, ok := filters["description"]; ok && description != "" {
		query = query.Where("description LIKE ?", "%"+description.(string)+"%")
	}
	if dateStart, ok := filters["date_start"]; ok && dateStart != "" {
		query = query.Where("due_date >= ?", dateStart)
	}
	if dateEnd, ok := filters["date_end"]; ok && dateEnd != "" {
		query = query.Where("due_date <= ?", dateEnd)
	}
	if min, ok := filters["min"]; ok && !min.(decimal.Decimal).IsZero() {
		query = query.Where("value >= ?", min)
	}
	if max, ok := filters["max"]; ok && !max.(decimal.Decimal).IsZero() {
		query = query.Where("value <= ?", max)
	}
	if paid, ok := filters["paid"]; ok && paid != "" {
		query = query.Where("paid = ?", paid)
	}
	if category, ok := filters["category_id"]; ok && category != "" {
		query = query.Where("category_id = ?", category)
	}

	err := query.Find(&expenses).Error
	//fmt.Println(query.Debug().Find(&expenses), filters["paid"])

	return expenses, err
}

func (database *expenseDatabase) Update(ctx context.Context, expense entity.Expense) error {
	err := database.DB.Model(&expense).Updates(map[string]interface{}{
		"description": expense.Description,
		"due_date":    expense.DueDate,
		"paid":        expense.Paid,
		"value":       expense.Value,
		"updated_at":  time.Now(),
	}).Error
	return err
}

func (database *expenseDatabase) Delete(ctx context.Context, expense entity.Expense) error {
	err := database.DB.Delete(&expense).Error
	return err
}
