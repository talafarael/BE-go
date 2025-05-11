package services

import (
	"gin/internal/config"
	"gin/internal/repository"
	"gin/pkg/hash"
	"gin/pkg/jwt"
)

func ConfigService(repo repository.Store, config *config.Config) *Service {
	hashService := hash.NewHashService()
	jwtService := jwt.NewJwtService(config.Secret)

	service := NewService(
		&ServiceOptions{
			Repo:        &repo,
			JwtService:  &jwtService,
			HashService: &hashService,
		})
	return &service
}
