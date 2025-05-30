package auth_services

import (
<<<<<<< HEAD:internal/services/auth.go
	userModels "gin/internal/models/user"
=======
	"gin/internal/auth/auth_dto"
	"gin/internal/dto"
>>>>>>> 5f8489a (feat:first-step):internal/auth/auth_services/auth_service.go
	"gin/internal/repository"
	"gin/pkg/hash"
	"gin/pkg/jwt"
)

type AuthService interface {
	Register(user *userModels.RegisterDto) (string, error)
	Login(userDto *userModels.LoginDto) (string, error)
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

<<<<<<< HEAD:internal/services/auth.go
func (a *authService) Register(user *userModels.RegisterDto) (string, error) {
=======
func (a *authService) Register(user *auth_dto.RegisterDto) (string, error) {
>>>>>>> 5f8489a (feat:first-step):internal/auth/auth_services/auth_service.go
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

<<<<<<< HEAD:internal/services/auth.go
func (a *authService) Login(userDto *userModels.LoginDto) (string, error) {
=======
func (a *authService) Login(userDto *auth_dto.LoginDto) (string, error) {
>>>>>>> 5f8489a (feat:first-step):internal/auth/auth_services/auth_service.go
	user, err := a.Repo.User().GetUserByEmail(userDto.Email)
	if err != nil {
		return "", err
	}
	isPasswordValid := a.HashService.CheckPasswordHash(userDto.Password, user.Password)
	if !isPasswordValid {
		return "", err
	}
	token, err := a.JwtService.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
