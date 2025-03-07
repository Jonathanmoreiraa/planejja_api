package handler

import (
	"net/http"
	"strconv"

	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	revenue_contract "github.com/jonathanmoreiraa/planejja/internal/usecase/revenue/contract"

	"github.com/gin-gonic/gin"
)

type RevenueHandler struct {
	revenueUseCase revenue_contract.RevenueUseCase
}

func NewRevenueHandler(usecase revenue_contract.RevenueUseCase) *RevenueHandler {
	return &RevenueHandler{
		revenueUseCase: usecase,
	}
}

//TODO: criar uma rota para realizar o CRUD primeiramente, em seguida irá existir uma rota para filtrar por categoria (uma ou mais)

func (cr *RevenueHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao localizar usuário com id informado!",
			"data":    nil,
			"error":   "Erro ao localizar usuário com id informado!",
		})
		return
	}

	revenue, err := cr.revenueUseCase.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Erro ao localizar usuário!",
			"data":    nil,
			"error":   "Erro ao localizar usuário!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"oi": revenue,
	})
}

func (cr *RevenueHandler) Create(ctx *gin.Context) {
	var revenue entity.Revenue

	if err := ctx.ShouldBindJSON(&revenue); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Erro ao criar a conta com esses dados.",
			"data":    nil,
		})
		return
	}

	_, err := cr.revenueUseCase.Create(ctx.Request.Context(), revenue)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error":   "Erro ao criar a conta!",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Conta criada com sucesso!",
		"data":    nil,
	})
}

func (cr *RevenueHandler) Update(ctx *gin.Context) {
	var revenue entity.Revenue

	if err := ctx.ShouldBindJSON(&revenue); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Erro ao criar a conta com esses dados.",
			"data":    nil,
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao editar o usuário",
			"data":    nil,
			"error":   "Erro ao editar o usuário",
		})
		return
	}
	revenue.ID = id

	//TODO: adicionar uma validação para também atualizar o email e alterar no repository

	err = cr.revenueUseCase.Update(ctx.Request.Context(), revenue)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": "Erro ao criar a conta!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuário editado com sucesso!",
		"data": gin.H{
			"oi": revenue,
		},
	})

}

func (cr *RevenueHandler) Delete(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse id",
		})
		return
	}

	requestCtx := ctx.Request.Context()
	revenue, err := cr.revenueUseCase.FindByID(requestCtx, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	if revenue == (entity.Revenue{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Revenue is not booking yet",
		})
		return
	}

	cr.revenueUseCase.Delete(ctx, revenue)

	ctx.JSON(http.StatusOK, gin.H{"message": "Revenue was deleted successfully"})
}
