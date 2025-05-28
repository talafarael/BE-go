package middleware

import (
	"gin/internal/config"
	"gin/internal/repository"
	"gin/pkg/check_auth_header"
	"gin/pkg/jwt"
)

func ConfigMiddleware(repo repository.Store, config *config.Config) *AuthMiddleware {
	jwtService := jwt.NewJwtService(config.Secret)
	checkAuthHeader := check_auth_header.CheckAuthHeader{}

	authMiddleware := NewAuthMiddleware(
		&AuthMiddlewareOptions{
			CheckAuthHeader: &checkAuthHeader,
			JwtService:      &jwtService,
			Repo:            &repo,
		})
	return &authMiddleware
}
