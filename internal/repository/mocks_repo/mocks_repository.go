package mocks_repository

import (
	"gin/internal/models"
	"gin/internal/repository"
)

type Mocks struct {
	users              map[uint]*models.User
	mockUserRepository MockUserRepository
}

func NewRepository() *Mocks {
	m := &Mocks{
		users: make(map[uint]*models.User),
	}

	m.users[1] = &models.User{
		ID:       1,
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	return m
}

func (m *Mocks) AddUser(user *models.User) {
	m.users[user.ID] = &models.User{
		ID:       1,
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
}

func (m *Mocks) User() repository.UserRepository {
	if m.mockUserRepository != nil {
		return m.mockUserRepository
	}

	m.mockUserRepository = NewMockUserRepository(m)

	return m.mockUserRepository
}
