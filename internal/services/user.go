package services

import (
	userModels "gin/internal/models/user"
	"gin/internal/repository"
	response_error "gin/pkg/error"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Get(ctx *gin.Context) (userModels.User, error)
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

func (u *userService) Get(ctx *gin.Context) (userModels.User, error) {
	userValue, exists := ctx.Get("user")
	if !exists {
		return userModels.User{}, response_error.ErrUnauthorized
	}

	user, ok := userValue.(*userModels.User)
	if !ok {
		return userModels.User{}, response_error.ErrUnauthorized
	}
	return *user, nil
}
