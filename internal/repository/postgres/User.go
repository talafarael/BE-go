package postgres

import (
	"gin/internal/models"
	"log"
)

type UserRepo struct {
	store *Repository
}

func (u UserRepo) CreateUser(user *models.User) (models.User, error) {
	err := u.store.db.Create(user).Error
	if err != nil {
		return models.User{}, err
	}
	log.Println(user)
	return *user, nil
}

func (u UserRepo) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := u.store.db.First(&user, id).Error
	if err != nil {
		return &models.User{}, err
	}
	log.Println(user)
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
