package services

import (
	"gin/internal/models"
	"gin/internal/repository"
	response_error "gin/pkg/error"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Get(ctx *gin.Context) (*models.User, error)
}

type userService struct {
	UserServiceOptions
}
type UserServiceOptions struct {
	Repo repository.Store
}

func NewUserService(options UserServiceOptions) UserService {
	return &userService{
		UserServiceOptions: options,
	}
}

func (u *userService) Get(ctx *gin.Context) (*models.User, error) {
	userValue, exists := ctx.Get("user")
	if !exists {
		return &models.User{}, response_error.ErrUnauthorized
	}

	user, ok := userValue.(*models.User)
	if !ok {
		return &models.User{}, response_error.ErrUnauthorized
	}
	return user, nil
}
