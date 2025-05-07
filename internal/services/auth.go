package services

import (
	"gin/internal/models"
	"gin/internal/repository"
)

type AuthService interface {
	Register(user *models.User) models.User
}

type authService struct {
	repo repository.Store
}

// construcotr
func NewAuthService(repo repository.Store) AuthService {
	return &authService{}
}

func (a *authService) Register(user *models.User) models.User {
	return *user
}
