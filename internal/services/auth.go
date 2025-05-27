package services

import (
	"gin/internal/dto"
	"gin/internal/repository"
	response_error "gin/pkg/error"
	"gin/pkg/hash"
	"gin/pkg/jwt"
)

type AuthService interface {
	Register(user *dto.RegisterDto) (string, error)
	Login(userDto *dto.LoginDto) (string, error)
}

type authService struct {
	AuthServiceOptions
}
type AuthServiceOptions struct {
	Repo        repository.Store
	JwtService  jwt.JwtService
	HashService hash.HashService
}

// construcotr
func NewAuthService(options AuthServiceOptions) AuthService {
	return &authService{
		AuthServiceOptions: options,
	}
}

func (a *authService) Register(user *dto.RegisterDto) (string, error) {
	hashPassword, err := a.HashService.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashPassword
	newUser, err := a.Repo.User().CreateUser(user)
	if err != nil {
		return "", err
	}
	token, err := a.JwtService.CreateToken(newUser.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authService) Login(userDto *dto.LoginDto) (string, error) {
	user, err := a.Repo.User().GetUserByEmail(userDto.Email)
	if err != nil {
		return "", err
	}
	isPasswordValid := a.HashService.CheckPasswordHash(userDto.Password, user.Password)
	if !isPasswordValid {
		return "", response_error.ErrUnauthorized
	}
	token, err := a.JwtService.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
