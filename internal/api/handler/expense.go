package handler

import (
	"fmt"
	"net/http"
	"strconv"

	error_message "github.com/jonathanmoreiraa/planejja/internal/domain/error"
	"github.com/jonathanmoreiraa/planejja/internal/domain/model"
	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	category_contract "github.com/jonathanmoreiraa/planejja/internal/usecase/category/contract"
	expense_contract "github.com/jonathanmoreiraa/planejja/internal/usecase/expense/contract"
	"github.com/jonathanmoreiraa/planejja/pkg/log"
	"github.com/shopspring/decimal"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	expenseUseCase  expense_contract.ExpenseUseCase
	categoryUseCase category_contract.CategoryUseCase
}

type ExpenseFilters struct {
	Description string          `json:"description"`
	Min         decimal.Decimal `json:"min"`
	Max         decimal.Decimal `json:"max"`
	CategoryID  int             `json:"category_id"`
	Paid        int             `json:"paid"`
	DateStart   string          `json:"date_start"`
	DateEnd     string          `json:"date_end"`
}

func NewExpenseHandler(expenseUseCase expense_contract.ExpenseUseCase, categoryUseCase category_contract.CategoryUseCase) *ExpenseHandler {
	return &ExpenseHandler{
		expenseUseCase:  expenseUseCase,
		categoryUseCase: categoryUseCase,
	}
}

func (cr *ExpenseHandler) Create(ctx *gin.Context) {
	var expense entity.Expense

	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": error_message.ErrCreateExpense,
		})
		log.NewLogger().Error(err)
		return
	}

	userId, err := GetUserIdByToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":      http.StatusUnprocessableEntity,
			"message":   error_message.ErrCreateExpense,
			"more_info": "Verifique as informações do usuário logado!",
		})
		return
	}
	expense.UserID = userId

	category, err := cr.categoryUseCase.GetCategoryById(ctx.Request.Context(), expense.CategoryID, userId)
	if err != nil || category.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":      http.StatusUnprocessableEntity,
			"message":   error_message.ErrCreateExpense,
			"more_info": "Verifique as informações da categoria!",
		})
		return
	}

	_, err = cr.expenseUseCase.Create(ctx.Request.Context(), expense)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": error_message.ErrCreateExpense,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusOK,
		"message": "Despesa criada com sucesso!",
	})
}

func (cr *ExpenseHandler) FindAll(ctx *gin.Context) {
	userId, err := GetUserIdByToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":      http.StatusUnprocessableEntity,
			"message":   error_message.ErrCreateExpense,
			"more_info": "Verifique as informações do usuário logado!",
		})
		return
	}

	expenses, err := cr.expenseUseCase.GetAllExpenses(ctx.Request.Context(), userId)
	if err != nil || len(expenses) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Erro ao encontrar as despesas",
		})
		return
	}

	ctx.JSON(http.StatusOK, expenses)
}

func (cr *ExpenseHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Erro ao identificar o id da despesa",
		})
		return
	}

	expense, err := cr.expenseUseCase.GetExpense(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Erro ao encontrar a despesa",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":          expense.ID,
		"description": expense.Description,
		"due_date":    expense.DueDate,
		"paid":        expense.Paid,
		"value":       expense.Value,
	})
}

func (cr *ExpenseHandler) FindByFilters(ctx *gin.Context) {
	var expensesFilter ExpenseFilters
	if err := ctx.ShouldBindJSON(&expensesFilter); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": error_message.ErrFindExpense,
		})
		log.NewLogger().Error(err)
		return
	}

	userId, err := GetUserIdByToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":      http.StatusUnprocessableEntity,
			"message":   error_message.ErrCreateExpense,
			"more_info": "Verifique as informações do usuário logado!",
		})
		return
	}

	filters := make(map[string]any)
	filters["description"] = expensesFilter.Description
	filters["date_start"] = expensesFilter.DateStart
	filters["date_end"] = expensesFilter.DateEnd
	filters["min"] = expensesFilter.Min
	filters["max"] = expensesFilter.Max
	filters["paid"] = expensesFilter.Paid
	filters["user_id"] = userId

	if int(expensesFilter.CategoryID) > 0 {
		filters["category_id"] = expensesFilter.CategoryID
	}

	expenses, err := cr.expenseUseCase.GetExpenses(ctx.Request.Context(), filters)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Erro ao encontrar a despesa",
		})
		return
	}

	ctx.JSON(http.StatusOK, expenses)
}

func (cr *ExpenseHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Erro ao identificar o id da despesa",
		})
		return
	}

	var expense model.Expense

	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao realizar o login!",
		})
		log.NewLogger().Error(err)
		return
	}

	expense.ID = id

	err = cr.expenseUseCase.Update(ctx.Request.Context(), expense)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"code":    http.StatusConflict,
			"message": "Erro ao criar a conta!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Despesa editada com sucesso!",
		"data": gin.H{
			"id": expense.ID,
		},
	})
}

func (cr *ExpenseHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Erro ao identificar o id da despesa",
		})
		return
	}

	requestCtx := ctx.Request.Context()
	expense, err := cr.expenseUseCase.GetExpense(requestCtx, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	userId, err := GetUserIdByToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao apagar a despesa",
		})
		return
	}
	if expense.UserID != userId {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao apagar a despesa com esse usuário",
		})
		log.NewLogger().Error(fmt.Errorf("As despesa com id %d não está relacionado com o usuário logado com id %d", expense.UserID, userId))
		return
	}

	err = cr.expenseUseCase.Delete(ctx, expense)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao apagar a despesa",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Despesa deletada com sucesso!",
	})
}
