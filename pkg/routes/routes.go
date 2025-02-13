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

	api.GET("/users", Handler.FindAll)
	api.GET("/users/:id", Handler.FindByID)
	api.DELETE("/users/:id", Handler.Delete)

	api.GET("/receitas", Handler.FindAll)
	api.GET("/receitas/:id", Handler.FindByID)
	api.POST("/receitas", Handler.Save)
	api.DELETE("/receitas/:id", Handler.Delete)

	api.GET("/despesas", Handler.FindAll)
	api.GET("/despesas/:id", Handler.FindByID)
	api.POST("/despesas", Handler.Save)
	api.DELETE("/despesas/:id", Handler.Delete)

	api.GET("/reservas", Handler.FindAll)
	api.GET("/reservas/:id", Handler.FindByID)
	api.POST("reservas", Handler.Save)
	api.DELETE("/reservas/:id", Handler.Delete)

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
