package services

import (
	"errors"
	"github.com/EvgeniiAndronov/auth-service/internal/models"
	"github.com/EvgeniiAndronov/auth-service/internal/repository"
	"github.com/EvgeniiAndronov/auth-service/pkg/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

type Secret struct {
	Word string
}

func RegisterUser(req models.LoginRequest) (*models.AuthResponse, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := &models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := repository.CreateUser(user); err != nil {
		return nil, err
	}

	secretWord := LoadSecret().Word

	token, err := jwt.GenerateToken(*user, secretWord)
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{Token: token, User: *user}, nil
}

func AuthUser(req models.LoginRequest) (*models.AuthResponse, error) {
	secretWord := LoadSecret().Word

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := &models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	created, err := repository.FoundUser(user)
	if err != nil && created != true {
		return nil, errors.New("User not found")
	}

	token, err := jwt.GenerateToken(*user, secretWord)
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{Token: token, User: *user}, nil
}

func LoadSecret() Secret {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	return Secret{
		Word: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
