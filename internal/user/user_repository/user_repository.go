package user_repository

import (
	"gin/internal/auth/auth_dto"
	"gin/internal/user/user_models"
)

type UserRepository interface {
	CreateUser(user *auth_dto.RegisterDto) (user_models.User, error)
	GetUserByID(id uint) (*user_models.User, error)
	GetUserByEmail(email string) (*user_models.User, error)
	DeleteUser(id int) error
}
