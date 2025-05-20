package mocks_repository

import (
	"fmt"
	"gin/internal/dto"
	"gin/internal/models"
	"gin/internal/repository"

	response_error "gin/pkg/error"
)

type mockUserRepository struct {
	store *Mocks
}
type MockUserRepository interface {
	repository.UserRepository
}

func NewMockUserRepository(store *Mocks) *mockUserRepository {
	return &mockUserRepository{
		store: store,
	}
}

func (m *mockUserRepository) CreateUser(user *dto.RegisterDto) (models.User, error) {
	newUser := models.User{
		ID:    uint(len(m.store.users) + 1),
		Email: user.Email,
	}
	m.store.users[newUser.ID] = &models.User{
		ID:       1,
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	return newUser, nil
}

func (m *mockUserRepository) GetUserByID(id uint) (*models.User, error) {
	user, exists := m.store.users[id]
	if !exists {
		return nil, response_error.ErrUserNotFound
	}

	return user, nil
}

func (m *mockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	for _, user := range m.store.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserRepository) DeleteUser(id int) error {
	_, exists := m.store.users[uint(id)]
	if !exists {
		return fmt.Errorf("user not found")
	}
	delete(m.store.users, uint(id))
	return nil
}
