package repository

import "gin/internal/models"

type UserRepository interface {
	CreateUser(user *models.User) (models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}
