package handler

import (
	"github/jonathanmoreiraa/planejja/pkg/domain"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ReservaHandler struct {
	reservaUseCase services.ReservaUseCase
}

func NewReservaHandler(usecase services.ReservaUseCase) *ReservaHandler {
	return &ReservaHandler{
		reservaUseCase: usecase,
	}
}

func (cr *ReservaHandler) FindAll(c *gin.Context) {
	reservas, err := cr.reservaUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &reservas)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *ReservaHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	reserva, err := cr.reservaUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &reserva)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *ReservaHandler) Save(c *gin.Context) {
	var reserva domain.Reservas

	if err := c.BindJSON(&reserva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	reserva, err := cr.reservaUseCase.Save(c.Request.Context(), reserva)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &reserva)

		c.JSON(http.StatusOK, response)
	}
}

func (cr *ReservaHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse id",
		})
		return
	}

	ctx := c.Request.Context()
	reserva, err := cr.reservaUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if reserva == (domain.Reservas{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Reserva is not booking yet",
		})
		return
	}

	cr.reservaUseCase.Delete(ctx, reserva)

	c.JSON(http.StatusOK, gin.H{"message": "Reserva is deleted successfully"})
}
