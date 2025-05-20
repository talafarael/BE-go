package repository

import (
	"gin/internal/dto"
	"gin/internal/models"
)

type UserRepository interface {
	CreateUser(user *dto.RegisterDto) (models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(id int) error
}
