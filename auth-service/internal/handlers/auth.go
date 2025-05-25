package handlers

import (
	"github.com/EvgeniiAndronov/auth-service/internal/models"
	"github.com/EvgeniiAndronov/auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.RegisterUser(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func Login(c *gin.Context) {
	//var req models.LoginRequest
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//user, err := services.Login(req)
	//if err != nil {
	//	c.JSON(http.StatusConflict, gin.H{"error": err})
	//	return
	//}
	//
	//c.JSON(http.StatusAccepted, gin.H{"user": user})
}
