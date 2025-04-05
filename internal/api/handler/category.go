package handler

import (
	"net/http"

	error_message "github.com/jonathanmoreiraa/planejja/internal/domain/error"
	"github.com/jonathanmoreiraa/planejja/internal/domain/model"
	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	category_contract "github.com/jonathanmoreiraa/planejja/internal/usecase/category/contract"
	"github.com/jonathanmoreiraa/planejja/pkg/log"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUseCase category_contract.CategoryUseCase
}

func NewCategoryHandler(usecase category_contract.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		categoryUseCase: usecase,
	}
}

func (cr *CategoryHandler) Create(ctx *gin.Context) {
	var category entity.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": error_message.ErrCreateCategory,
		})
		log.NewLogger().Error(err)
		return
	}

	userId, err := GetUserIdByToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":      http.StatusUnprocessableEntity,
			"message":   error_message.ErrCreateCategory,
			"more_info": "Verifique as informações do usuário logado!",
		})
		return
	}

	category.UserID = userId
	_, err = cr.categoryUseCase.Create(ctx.Request.Context(), category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": error_message.ErrCreateCategory,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusOK,
		"message": "Categoria criada com sucesso!",
	})
}

func (cr *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	userId, err := GetUserIdByToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":      http.StatusUnprocessableEntity,
			"message":   error_message.ErrCreateCategory,
			"more_info": "Verifique as informações do usuário logado!",
		})
		return
	}

	categories, err := cr.categoryUseCase.GetAllCategories(ctx.Request.Context(), userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Erro ao encontrar as categorias",
		})
		return
	}

	if len(categories) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": error_message.ErrFindCategory,
		})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (cr *CategoryHandler) FindCategory(ctx *gin.Context) {
	userId, err := GetUserIdByToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":      http.StatusUnprocessableEntity,
			"message":   error_message.ErrCreateCategory,
			"more_info": "Verifique as informações do usuário logado!",
		})
		return
	}

	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Erro ao localizar categoria, verifique os parâmetros e tente novamente!",
		})
		log.NewLogger().Error(err)
		return
	}

	categories, err := cr.categoryUseCase.GetCategory(ctx.Request.Context(), category.Name, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Erro ao encontrar a receita",
		})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}
