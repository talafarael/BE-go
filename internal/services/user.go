package services

import (
	"gin/internal/models"
	"gin/internal/repository"
)

type UserService interface {
	Get(id string) models.User
}

type userService struct {
	repo repository.Store
}

func NewUserService(repo repository.Store) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) Get(id string) models.User {
	return models.User{
		Name:  "afa",
		Email: "email@example.com",
	}
}
