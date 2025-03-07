package route

import (
	"os"

	"github.com/jonathanmoreiraa/planejja/internal/api/handler"
	"github.com/jonathanmoreiraa/planejja/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

type HandlerGroup struct {
	UserHandler    *handler.UserHandler
	RevenueHandler *handler.RevenueHandler
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

	api.POST("/revenue/add", Handlers.RevenueHandler.FindByID)
	api.GET("/revenue/:id", Handlers.RevenueHandler.FindByID)
	api.PUT("/revenue/:id", Handlers.RevenueHandler.Update)
	api.DELETE("/revenue/:id", Handlers.RevenueHandler.Delete)

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
