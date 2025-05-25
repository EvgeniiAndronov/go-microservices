package jwt

import (
	"github.com/EvgeniiAndronov/auth-service/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(user models.User, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
