package user_mock_repository

import (
	"fmt"
	"gin/internal/auth/auth_dto"
	"gin/internal/user/user_models"
	"gin/internal/user/user_repository"

	response_error "gin/pkg/error"
)

type mockUserRepository struct {
	users map[uint]*user_models.User
}

func NewMockUserRepository(user map[uint]*user_models.User) user_repository.UserRepository {
	return &mockUserRepository{
		users: user,
	}
}

func (m *mockUserRepository) CreateUser(user *auth_dto.RegisterDto) (user_models.User, error) {
	if user.Email == "test@example.com" {
		return user_models.User{}, response_error.ErrUserAlredy
	}
	newUser := user_models.User{
		ID:    uint(len(m.users) + 1),
		Email: user.Email,
	}
	m.users[newUser.ID] = &user_models.User{
		ID:       1,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return newUser, nil
}

func (m *mockUserRepository) GetUserByID(id uint) (*user_models.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, response_error.ErrUserNotFound
	}

	return user, nil
}

func (m *mockUserRepository) GetUserByEmail(email string) (*user_models.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, response_error.ErrUserNotFound
}

func (m *mockUserRepository) DeleteUser(id int) error {
	_, exists := m.users[uint(id)]
	if !exists {
		return fmt.Errorf("user not found")
	}
	delete(m.users, uint(id))
	return nil
}
