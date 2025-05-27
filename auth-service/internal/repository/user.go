package repository

import (
	"github.com/EvgeniiAndronov/auth-service/internal/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(connection *gorm.DB) {
	db = connection
}

func CreateUser(user *models.User) error {
	return db.Create(user).Error
}

func FoundUser(userInput *models.User) (bool, error) {
	var userFound *models.User

	if err := db.Where("email = ?", userInput.Email).First(&userFound).Error; err != nil {
		return false, err
	}

	if userFound.PasswordHash != userInput.PasswordHash {
		return false, nil
	} else if userFound.PasswordHash == userInput.PasswordHash {
		return true, nil
	}
	return false, nil
}
