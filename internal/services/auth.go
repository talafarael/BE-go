package services

import (
	customErrors "gin/internal/domain/customErrors"
	"gin/internal/models"
	"gin/internal/repository"
	hash_password "gin/pkg/hashPassword"
)

type AuthService interface {
	Register(user *models.User) (models.User, error)
}

type authService struct {
	repo repository.Store
}

// construcotr
func NewAuthService(repo repository.Store) AuthService {
	return &authService{
		repo: repo,
	}
}

func (a *authService) Register(user *models.User) (models.User, error) {
	hashPassword, err := hash_password.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashPassword
	newUser, err := a.repo.User().CreateUser(user)
	if err != nil {
		return models.User{}, customErrors.ErrUserAlreadyExists
	}
	return newUser, nil
}
