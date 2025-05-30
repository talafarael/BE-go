package user_mock_repository

import (
	"fmt"
	"gin/internal/dto"
	"gin/internal/models"
	"gin/internal/user/user_repository"

	response_error "gin/pkg/error"
)

type mockUserRepository struct {
	store *Mocks
}

func NewMockUserRepository(store *Mocks) user_repository.UserRepository {
	return &mockUserRepository{
		store: store,
	}
}

func (m *mockUserRepository) CreateUser(user *dto.RegisterDto) (models.User, error) {
	if user.Email == "test@example.com" {
		return models.User{}, response_error.ErrUserAlredy
	}
	newUser := models.User{
		ID:    uint(len(m.store.users) + 1),
		Email: user.Email,
	}
	m.store.users[newUser.ID] = &models.User{
		ID:       1,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
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
	return nil, response_error.ErrUserNotFound
}

func (m *mockUserRepository) DeleteUser(id int) error {
	_, exists := m.store.users[uint(id)]
	if !exists {
		return fmt.Errorf("user not found")
	}
	delete(m.store.users, uint(id))
	return nil
}
