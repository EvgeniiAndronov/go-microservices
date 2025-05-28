package midlware

import (
	"github.com/EvgeniiAndronov/auth-service/internal/models"
	"github.com/EvgeniiAndronov/auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMidlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.AuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token := req.Token
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Empty token"})
			return
		}

		userData, err := services.AuthsMidlware(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Wrong token",
			})
			return
		}

		c.Set("userData", userData)
		c.Next()
	}
}
