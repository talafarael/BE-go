package services

import (
	"gin/internal/repository"
	"gin/pkg/hash"
	"gin/pkg/jwt"
)

type Service interface {
	AuthService
	UserService
	VacancyService
}
type ServiceOptions struct {
	Repo        *repository.Store
	JwtService  *jwt.JwtService
	HashService *hash.HashService
}
type service struct {
	AuthService
	UserService
	VacancyService
}

func NewService(options *ServiceOptions) Service {
	return &service{
		AuthService: NewAuthService(AuthServiceOptions{
			Repo:        *options.Repo,
			JwtService:  *options.JwtService,
			HashService: *options.HashService,
		}),
		UserService: NewUserService(UserServiceOptions{
			Repo: *options.Repo,
		}),
		VacancyService: NewVacancyService(VacancyServiceOptions{
			Repo: *options.Repo,
		}),
	}
}
