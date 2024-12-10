package handler

import (
	"github/jonathanmoreiraa/planejja/pkg/domain"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ReceitaHandler struct {
	receitaUseCase services.ReceitaUseCase
}

func NewReceitaHandler(usecase services.ReceitaUseCase) *ReceitaHandler {
	return &ReceitaHandler{
		receitaUseCase: usecase,
	}
}

func (cr *ReceitaHandler) FindAll(c *gin.Context) {
	receitas, err := cr.receitaUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &receitas)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *ReceitaHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	receita, err := cr.receitaUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &receita)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *ReceitaHandler) Save(c *gin.Context) {
	var receita domain.Receitas

	if err := c.BindJSON(&receita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	receita, err := cr.receitaUseCase.Save(c.Request.Context(), receita)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &receita)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *ReceitaHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse id",
		})
		return
	}

	ctx := c.Request.Context()
	receita, err := cr.receitaUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if receita == (domain.Receitas{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Receita is not booking yet",
		})
		return
	}

	cr.receitaUseCase.Delete(ctx, receita)

	c.JSON(http.StatusOK, gin.H{"message": "Receita is deleted successfully"})
}
