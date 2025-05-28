package api

import (
	"github.com/EvgeniiAndronov/auth-service/internal/handlers"
	"github.com/EvgeniiAndronov/auth-service/internal/midlware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/register", handlers.Register) // /api/v1/auth/register/
		authGroup.POST("/login", handlers.Login)       // /api/v1/auth/login/
	}

	meGroup := router.Group("/api/v1/")
	meGroup.Use(midlware.AuthMidlware())
	{
		meGroup.GET("/me", handlers.Me) // /api/v1/auth/me
	}

	return router
}
