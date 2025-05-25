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
