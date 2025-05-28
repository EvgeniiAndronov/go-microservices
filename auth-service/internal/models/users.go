package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null;size:50"`
	Email        string `gorm:"unique;not null;size:100"`
	PasswordHash string `gorm:"not null;size:100" json:"-"`
	LastLogin    time.Time
	IsActive     bool `gorm:"default:true"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Username string `json:"username" binding:"required,min=5,max=32"`
}

type AuthRequest struct {
	Token string `json:"token"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
