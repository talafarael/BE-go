package services

import (
	userModels "gin/internal/models/user"
	"gin/internal/repository"
)

type UserService interface {
	Get(id string) (userModels.User, error)
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

func (u *userService) Get(token string) (userModels.User, error) {
	//id, err := u.JwtService.VerifyToken(token)
	//if err != nil {
	//return userModels.User{}, err
	//}

	user, err := u.Repo.User().GetUserByID(1)
	if err != nil {
		return userModels.User{}, err
	}
	return *user, nil
}
