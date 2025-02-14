package http

import (
	"github/jonathanmoreiraa/planejja/pkg/api/handler"
	"github/jonathanmoreiraa/planejja/pkg/api/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(Handler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())

	engine.POST("/login", Handler.Login)
	engine.POST("/user/new", Handler.Save)

	api := engine.Group("/api", middleware.AuthorizationMiddleware)
	api.GET("/user/:id", Handler.FindByID)
	api.PUT("/user/:id", Handler.Update)
	api.DELETE("/user/:id", Handler.Delete)

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
