package handler

import (
	"github/jonathanmoreiraa/planejja/pkg/domain"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type DespesaHandler struct {
	despesaUseCase services.DespesaUseCase
}

func NewDespesaHandler(usecase services.DespesaUseCase) *DespesaHandler {
	return &DespesaHandler{
		despesaUseCase: usecase,
	}
}

func (cr *DespesaHandler) FindAll(c *gin.Context) {
	despesas, err := cr.despesaUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &despesas)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *DespesaHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	despesa, err := cr.despesaUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &despesa)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *DespesaHandler) Save(c *gin.Context) {
	var despesa domain.Despesas

	if err := c.BindJSON(&despesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	despesa, err := cr.despesaUseCase.Save(c.Request.Context(), despesa)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &despesa)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *DespesaHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse id",
		})
		return
	}

	ctx := c.Request.Context()
	despesa, err := cr.despesaUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if despesa == (domain.Despesas{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Despesa is not booking yet",
		})
		return
	}

	cr.despesaUseCase.Delete(ctx, despesa)

	c.JSON(http.StatusOK, gin.H{"message": "Despesa is deleted successfully"})
}
