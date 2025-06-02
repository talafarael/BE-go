package user_postgres

import (
	"gin/internal/auth/auth_dto"
	"gin/internal/user/user_models"
	"gin/internal/user/user_repository"
	response_error "gin/pkg/error"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user_repository.UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (u UserRepo) CreateUser(user *auth_dto.RegisterDto) (user_models.User, error) {
	createdUser := user_models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := u.db.Create(&createdUser).Error; err != nil {
		return user_models.User{}, response_error.ErrUserAlredy
	}

	return createdUser, nil
}

func (u UserRepo) GetUserByEmail(email string) (*user_models.User, error) {
	var user user_models.User
	err := u.db.Preload("Vacancies").Where("email=?", email).First(&user).Error
	if err != nil {
		return &user_models.User{}, response_error.ErrUserNotFound
	}
	return &user, nil
}

func (u UserRepo) GetUserByID(id uint) (*user_models.User, error) {
	var user user_models.User
	err := u.db.Preload("Vacancies").First(&user, id).Error
	if err != nil {
		return &user_models.User{}, response_error.ErrUserNotFound
	}
	return &user, nil
}

func (u UserRepo) UpdateUser(user *user_models.User) error {
	// TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteUser(id int) error {
	// TODO implement me
	panic("implement me")
}
