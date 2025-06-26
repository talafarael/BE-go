package user_services

import (
	"gin/internal/infrastructure/repository"
	"gin/internal/user/user_models"
	response_error "gin/pkg/error"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Get(ctx *gin.Context) (*user_models.User, error)
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

func (u *userService) Get(ctx *gin.Context) (*user_models.User, error) {
	userValue, exists := ctx.Get("user")
	if !exists {
		return &user_models.User{}, response_error.ErrUnauthorized
	}

	user, ok := userValue.(*user_models.User)
	if !ok {
		return &user_models.User{}, response_error.ErrUnauthorized
	}
	return user, nil
}
