package handler

import (
	"github/jonathanmoreiraa/planejja/pkg/domain"
	services "github/jonathanmoreiraa/planejja/pkg/usecase/interface"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

//TODO: Assim que iniciar a tela no front, criar a função para recuperar a senha

type UserHandler struct {
	userUseCase services.UserUseCase
}

type LoginCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao localizar usuário com id informado!",
			"data":    nil,
			"error":   "Erro ao localizar usuário com id informado!",
		})
		return
	}

	user, err := cr.userUseCase.FindByID(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Erro ao localizar usuário!",
			"data":    nil,
			"error":   "Erro ao localizar usuário!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"password":   user.Password,
		"birth_date": user.BirthDate,
	})
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

	if err := ctx.ShouldBindJSON(&user); err != nil {
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
	user.ID = uint(id)

	//TODO: adicionar uma validação para também atualizar o email e alterar no repository

	err = cr.userUseCase.Update(ctx.Request.Context(), user)
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
			"name":       user.Name,
			"email":      user.Email,
			"password":   user.Password,
			"birth_date": user.BirthDate,
		},
	})

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

	var loginCredentials LoginCredentials
	if err := ctx.ShouldBindJSON(&loginCredentials); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Erro ao realizar o login, verifique as credenciais!",
			"error":   nil,
		})

		logger.Default.Error(ctx, err.Error())
		return
	}

	loginData, err := cr.userUseCase.Login(ctx, loginCredentials.Email, loginCredentials.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
			"error":   "Erro ao realizar o login!",
		})

		return
	}

	expirationTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Erro ao realizar o login!",
			"error":   "Erro ao realizar o login!",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": gin.H{
			"access_token": loginData["token"],
			"token_type":   "Bearer",
			"expires_in":   time.Now().Add(time.Hour * time.Duration(expirationTime)).Unix(),
		},
		"user": gin.H{
			"id": loginData["userId"],
		},
	})
}
