package services

import (
	"github.com/EvgeniiAndronov/auth-service/internal/models"
	"github.com/EvgeniiAndronov/auth-service/internal/repository"
	"github.com/EvgeniiAndronov/auth-service/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req models.LoginRequest) (*models.AuthResponse, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := &models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := repository.CreateUser(user); err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(*user, "secret")
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{Token: token, User: *user}, nil
}
