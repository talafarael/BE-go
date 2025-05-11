package postgres

import (
	"gin/internal/dto"
	"gin/internal/models"
	response_error "gin/pkg/error"
)

type UserRepo struct {
	store *Repository
}

func (u UserRepo) CreateUser(user *dto.RegisterDto) (models.User, error) {
	createdUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := u.store.db.Create(&createdUser).Error; err != nil {
		return models.User{}, response_error.ErrUserAlredy
	}

	return createdUser, nil
}

func (u UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.store.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return &models.User{}, response_error.ErrUserNotFound
	}
	return &user, nil
}

func (u UserRepo) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := u.store.db.First(&user, id).Error
	if err != nil {
		return &models.User{}, response_error.ErrUserNotFound
	}
	return &user, nil
}

func (u UserRepo) UpdateUser(user *models.User) error {
	// TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteUser(id int) error {
	// TODO implement me
	panic("implement me")
}
