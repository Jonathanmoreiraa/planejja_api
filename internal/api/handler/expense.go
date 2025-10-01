package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	Categories  []int           `json:"categories"`
	Paid        int             `json:"paid"`
	DateStart   string          `json:"date_start"`
	DateEnd     string          `json:"date_end"`
	Status      struct {
		Pending bool `json:"pending"`
		Paid    bool `json:"paid"`
		Overdue bool `json:"overdue"`
		DueSoon bool `json:"due_soon"`
	} `json:"status"`
}

type ExpenseInput struct {
	Description      string          `json:"description"`
	Value            decimal.Decimal `json:"value"`
	DueDate          *time.Time      `json:"due_date"`
	CategoryID       int             `json:"category_id"`
	MultiplePayments bool            `json:"multiple_payments"`
	NumInstallments  int             `json:"num_installments"`
	PaymentDay       int             `json:"payment_day"`
}

type ExpenseResponse struct {
	ID          int             `json:"id"`
	Description string          `json:"description"`
	Value       decimal.Decimal `json:"value"`
	DueDate     *time.Time      `json:"due_date"`
	Paid        int             `json:"paid"`
	Category    string          `json:"category"`
	CategoryID  int             `json:"category_id"`
}

func NewExpenseHandler(expenseUseCase expense_contract.ExpenseUseCase, categoryUseCase category_contract.CategoryUseCase) *ExpenseHandler {
	return &ExpenseHandler{
		expenseUseCase:  expenseUseCase,
		categoryUseCase: categoryUseCase,
	}
}

func (cr *ExpenseHandler) Create(ctx *gin.Context) {
	var input ExpenseInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": error_message.ErrCreateExpense,
		})
		log.NewLogger().Error(err)
		return
	}

	expense := entity.Expense{
		Description: input.Description,
		Value:       input.Value,
		DueDate:     input.DueDate,
		CategoryID:  input.CategoryID,
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

	_, err = cr.expenseUseCase.Create(ctx.Request.Context(), expense, input.MultiplePayments, input.NumInstallments, input.PaymentDay)
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
	var expensesResponse []ExpenseResponse

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

	for _, expense := range expenses {
		category, _ := cr.categoryUseCase.GetCategoryById(ctx.Request.Context(), expense.CategoryID, userId)
		expenseResponse := ExpenseResponse{
			ID:          expense.ID,
			Description: expense.Description,
			Value:       expense.Value,
			DueDate:     expense.DueDate,
			Paid:        expense.Paid,
			Category:    category.Name,
		}
		expensesResponse = append(expensesResponse, expenseResponse)
	}

	ctx.JSON(http.StatusOK, expensesResponse)
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
		log.NewLogger().Error(err)
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
	filters["status"] = expensesFilter.Status

	if len(expensesFilter.Categories) > 0 {
		filters["categories"] = expensesFilter.Categories
	}

	expenses, err := cr.expenseUseCase.GetExpenses(ctx.Request.Context(), filters)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Erro ao encontrar a despesa",
		})
		log.NewLogger().Error(err)
		return
	}

	var expenseResponses []ExpenseResponse

	for _, expense := range expenses {
		category, err := cr.categoryUseCase.GetCategoryById(ctx.Request.Context(), expense.CategoryID, userId)
		if err != nil {
			log.NewLogger().Error(err)
			continue
		}
		expenseResponse := ExpenseResponse{
			ID:          expense.ID,
			Description: expense.Description,
			Value:       expense.Value,
			DueDate:     expense.DueDate,
			Paid:        expense.Paid,
			Category:    category.Name,
			CategoryID:  category.ID,
		}
		expenseResponses = append(expenseResponses, expenseResponse)
	}

	if len(expenseResponses) <= 0 {
		ctx.JSON(http.StatusOK, []ExpenseResponse{})
		return
	}

	ctx.JSON(http.StatusOK, expenseResponses)
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
