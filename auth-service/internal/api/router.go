package api

import (
	"github.com/EvgeniiAndronov/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
		//authGroup.GET("/me", handlers.Me)
	}

	return router
}
