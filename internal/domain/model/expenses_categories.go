package model

type ExpensesCategories struct {
	ID         int      `json:"id" gorm:"primaryKey;autoIncrement"`
	ExpenseID  int      `json:"expense_id,omitempty" gorm:"not null"`
	Expense    Expense  `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
	CategoryID int      `json:"category_id,omitempty" gorm:"not null"`
	Category   Category `json:"-" gorm:"constraint:OnUpdate:CASCADE"`
}
