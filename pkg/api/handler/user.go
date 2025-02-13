package handler

import (
	"fmt"
	"github/jonathanmoreiraa/planejja/pkg/domain"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm/logger"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) FindAll(ctx *gin.Context) {
	users, err := cr.userUseCase.FindAll(ctx.Request.Context())

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &users)

		ctx.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) FindByID(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	user, err := cr.userUseCase.FindByID(ctx.Request.Context(), uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		ctx.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) Save(ctx *gin.Context) {
	var user domain.Users

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Erro ao criar a conta com esses dados.",
			"data":    nil,
		})
		return
	}

	_, err := cr.userUseCase.Save(ctx.Request.Context(), user)
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

func (cr *UserHandler) Update(ctx *gin.Context) {
	var user domain.Users

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao criar a conta.",
			"error":   err,
		})
		return
	}

	user, err := cr.userUseCase.Update(ctx.Request.Context(), user)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	response := Response{}
	copier.Copy(&response, &user)

	ctx.JSON(http.StatusOK, response)
}

func (cr *UserHandler) Delete(ctx *gin.Context) {
	paramsId := ctx.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot parse id",
		})
		return
	}

	requestCtx := ctx.Request.Context()
	user, err := cr.userUseCase.FindByID(requestCtx, uint(id))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	if user == (domain.Users{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User is not booking yet",
		})
		return
	}

	cr.userUseCase.Delete(ctx, user)

	ctx.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
}

func (cr *UserHandler) Login(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/json")

	var loginData LoginData
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Erro ao realizar o login!",
			"error":   nil,
		})

		logger.Default.Error(ctx, err.Error())
		return
	}

	if loginData.Email == "" {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "O e-mail não pode ser vazio!",
			"error":   nil,
		})

		return
	}

	if loginData.Password == "" {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "A senha não pode ser vazia!",
			"error":   nil,
		})

		return
	}

	tokenString, err := cr.userUseCase.Login(ctx, loginData.Email, loginData.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
			"error":   "Erro ao realizar o login!",
		})

		return
	}

	fmt.Println(tokenString)

	ctx.Writer.WriteHeader(http.StatusOK)

	ctx.JSON(http.StatusOK, gin.H{
		"token":      tokenString,
		"expires in": os.Getenv("JWT_EXPIRATION_TIME"),
	})

}
