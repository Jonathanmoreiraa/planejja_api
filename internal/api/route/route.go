package route

import (
	"net/http"
	"os"

	"github.com/jonathanmoreiraa/planejja/internal/api/handler"
	"github.com/jonathanmoreiraa/planejja/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

type HandlerGroup struct {
	UserHandler     *handler.UserHandler
	RevenueHandler  *handler.RevenueHandler
	ExpenseHandler  *handler.ExpenseHandler
	CategoryHandler *handler.CategoryHandler
	// Adicione outros handlers aqui futuramente
}

func NewServerHTTP(Handlers HandlerGroup) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())

	engine.POST("/login", Handlers.UserHandler.Login)
	engine.POST("/user/new", Handlers.UserHandler.Create)

	api := engine.Group("/api", middleware.AuthorizationMiddleware)
	api.GET("/user/:id", Handlers.UserHandler.FindByID)
	api.PUT("/user/:id", Handlers.UserHandler.Update)
	// api.DELETE("/user/:id", Handlers.UserHandler.Delete)

	api.POST("/revenue/add", Handlers.RevenueHandler.Create)
	api.GET("/revenue/:id", Handlers.RevenueHandler.FindByID)
	api.GET("/revenues", Handlers.RevenueHandler.FindAll)
	api.POST("/revenue/filter", Handlers.RevenueHandler.FindByFilters)
	api.PUT("/revenue/:id", Handlers.RevenueHandler.Update)
	api.DELETE("/revenue/:id", Handlers.RevenueHandler.Delete)

	api.POST("/category/add", Handlers.CategoryHandler.Create)
	api.GET("/categories", Handlers.CategoryHandler.GetAllCategories)
	api.POST("/category", Handlers.CategoryHandler.FindCategory)

	api.POST("/expense/add", Handlers.ExpenseHandler.Create)
	api.GET("/expense/:id", Handlers.ExpenseHandler.FindByID)
	api.GET("/expenses", Handlers.ExpenseHandler.FindAll)
	api.POST("/expense/filter", Handlers.ExpenseHandler.FindByFilters)
	api.PUT("/expense/:id", Handlers.ExpenseHandler.Update)
	api.DELETE("/expense/:id", Handlers.ExpenseHandler.Delete)

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   "Rota não encontrada",
			"message": "Verifique a URL e tente novamente",
		})
	})

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(getEnv("PORT", ":8080"))
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return ":" + value
	}
	return defaultValue
}
