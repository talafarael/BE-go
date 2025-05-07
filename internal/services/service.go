package services

import "gin/internal/repository"

type Service interface {
	AuthService
	UserService
}

type service struct {
	AuthService
	UserService
}

func NewService(repo repository.Store) Service {
	return &service{
		AuthService: NewAuthService(repo),
		UserService: NewUserService(repo),
	}
}
