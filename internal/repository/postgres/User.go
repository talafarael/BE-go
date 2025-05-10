package postgres

import (
	userModels "gin/internal/models/user"
)

type UserRepo struct {
	store *Repository
}

func (u UserRepo) CreateUser(user *userModels.RegisterDto) (userModels.User, error) {
	createdUser := userModels.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := u.store.db.Create(&createdUser).Error; err != nil {
		return userModels.User{}, err
	}

	return createdUser, nil
}

func (u UserRepo) GetUserByEmail(email string) (*userModels.User, error) {
	var user userModels.User
	err := u.store.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return &userModels.User{}, err
	}
	return &user, nil
}

func (u UserRepo) GetUserByID(id uint) (*userModels.User, error) {
	var user userModels.User
	err := u.store.db.First(&user, id).Error
	if err != nil {
		return &userModels.User{}, err
	}
	return &user, nil
}

func (u UserRepo) UpdateUser(user *userModels.User) error {
	// TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteUser(id int) error {
	// TODO implement me
	panic("implement me")
}
