package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jonathanmoreiraa/planejja/internal/api/middleware"
	"github.com/jonathanmoreiraa/planejja/internal/domain/model"
	user_contract "github.com/jonathanmoreiraa/planejja/internal/usecase/user/contract"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase user_contract.UserUseCase
}

type LoginCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserHandler(usecase user_contract.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) Create(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":      http.StatusUnprocessableEntity,
			"message":   "Erro ao criar a conta com esses dados.",
			"more_info": "Verifique o corpo do requerimento.",
		})

		return
	}

	_, err := cr.userUseCase.Create(ctx.Request.Context(), user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"code":    http.StatusConflict,
			"message": "Erro ao criar a conta!",
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Conta criada com sucesso!",
	})
}

func (cr *UserHandler) Login(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/json")

	var loginCredentials LoginCredentials
	if err := ctx.ShouldBindJSON(&loginCredentials); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao realizar o login, verifique as credenciais!",
		})

		return
	}

	loginData, err := cr.userUseCase.Login(ctx, loginCredentials.Email, loginCredentials.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao realizar o login!",
		})
		return
	}

	expirationTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao realizar o login!",
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

func (cr *UserHandler) Update(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao realizar o login!",
		})
		return
	}

	id, err := GetIdByToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao editar o usuário",
		})
		return
	}

	user.ID = id

	//TODO: adicionar uma validação para também atualizar o email e alterar no repository
	// TODO: adicionar rota para editar a senha (ela pode ser usada no esqueci a senha tb)

	err = cr.userUseCase.Update(ctx.Request.Context(), user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"code":    http.StatusConflict,
			"message": "Erro ao criar a conta!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuário editado com sucesso!",
		"data": gin.H{
			"name":       user.Name,
			"email":      user.Email,
			"birth_date": user.BirthDate,
		},
	})

}

func (cr *UserHandler) FindByID(ctx *gin.Context) {
	id, err := GetIdByToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Erro ao localizar usuário com id informado!",
		})
		return
	}

	user, err := cr.userUseCase.GetUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Erro ao localizar usuário!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"birth_date": user.BirthDate,
	})
}

func GetIdByToken(ctx *gin.Context) (int, error) {
	tokenString := ctx.Request.Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]

	id, err := middleware.ExtractUserId(tokenString)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// func (cr *UserHandler) Delete(ctx *gin.Context) {
// 	paramsId := ctx.Param("id")
// 	id, err := strconv.Atoi(paramsId)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Cannot parse id",
// 		})
// 		return
// 	}

// 	requestCtx := ctx.Request.Context()
// 	user, err := cr.userUseCase.FindByID(requestCtx, id)
// 	if err != nil {
// 		ctx.AbortWithStatus(http.StatusNotFound)
// 	}

// 	if user == (model.User{}) {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": "User is not booking yet",
// 		})
// 		return
// 	}

// 	cr.userUseCase.Delete(ctx, user)

// 	ctx.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
// }
