package jwt

import (
	"fmt"
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

func ParseToken(tokenString string, secretKey string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return "wrong pars", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := claims["sub"]
		return sub.(string), nil
	} else {
		return "wrong data", err
	}
}
