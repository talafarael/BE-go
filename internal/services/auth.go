package services

import (
	userModels "gin/internal/models/user"
	"gin/internal/repository"
	response_error "gin/pkg/error"
	"gin/pkg/hash"
	"gin/pkg/jwt"
)

type AuthService interface {
	Register(user *userModels.RegisterDto) (string, error)
	Login(userDto *userModels.LoginDto) (string, error)
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

func (a *authService) Register(user *userModels.RegisterDto) (string, error) {
	hashPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return "", response_error.ErrInternalServer
	}
	user.Password = hashPassword
	newUser, err := a.repo.User().CreateUser(user)
	if err != nil {
		return "", response_error.ErrUserAlredy
	}
	token, err := jwt.CreateToken(newUser.ID)
	if err != nil {
		return "", response_error.ErrJWTCreationFailed
	}
	return token, nil
}

func (a *authService) Login(userDto *userModels.LoginDto) (string, error) {
	user, err := a.repo.User().GetUserByEmail(userDto.Email)
	if err != nil {
		return "", response_error.ErrPasswordOrEmailNotCorrect
	}
	isPasswordValid := hash.CheckPasswordHash(userDto.Password, user.Password)
	if !isPasswordValid {
		return "", response_error.ErrPasswordOrEmailNotCorrect
	}
	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		return "", response_error.ErrJWTCreationFailed
	}
	return token, nil
}
