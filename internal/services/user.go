package services

import (
	userModels "gin/internal/models/user"
	"gin/internal/repository"
)

type UserService interface {
	Get(id string) *userModels.User
}

type userService struct {
	repo repository.Store
}

func NewUserService(repo repository.Store) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) Get(id string) *userModels.User {
	user := &userModels.User{
		Name:  "afa",
		Email: "email@example.com",
	}

	u.repo.User().GetUserByID(1)
	return user
}
