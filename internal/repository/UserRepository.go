package repository

import userModels "gin/internal/models/user"

type UserRepository interface {
	CreateUser(user *userModels.RegisterDto) (userModels.User, error)
	GetUserByID(id uint) (*userModels.User, error)
	GetUserByEmail(email string) (*userModels.User, error)
	UpdateUser(user *userModels.User) error
	DeleteUser(id int) error
}
