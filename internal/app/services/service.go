package services

import (
	"gin/internal/auth/auth_services"
	"gin/internal/infrastructure/repository"
	"gin/internal/user/user_services"
	"gin/pkg/hash"
	"gin/pkg/jwt"
)

type Service interface {
	auth_services.AuthService
	user_services.UserService
}
type ServiceOptions struct {
	Repo        *repository.Store
	JwtService  *jwt.JwtService
	HashService *hash.HashService
}
type service struct {
	auth_services.AuthService
	user_services.UserService
}

func NewService(options *ServiceOptions) Service {
	return &service{
		AuthService: auth_services.NewAuthService(auth_services.AuthServiceOptions{
			Repo:        *options.Repo,
			JwtService:  *options.JwtService,
			HashService: *options.HashService,
		}),
		UserService: user_services.NewUserService(user_services.UserServiceOptions{
			Repo: *options.Repo,
		}),
	}
}
