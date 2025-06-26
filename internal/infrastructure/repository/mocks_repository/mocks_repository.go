package mocks_repository

import (
	"gin/internal/user/user_models"
	"gin/internal/user/user_repository"
	"gin/internal/user/user_repository/user_mock_repository"
	"gin/internal/vacancy/vacancy_models"
	"gin/internal/vacancy/vacancy_repository"
	"gin/internal/vacancy/vacancy_repository/vacancy_mocks_repository"
)

type Mocks struct {
	users                 map[uint]*user_models.User
	mockUserRepository    user_repository.UserRepository
	vacancy               map[uint]*vacancy_models.Vacancy
	mockVacancyRepository vacancy_repository.VacancyRepository
}

func NewRepository() *Mocks {
	m := &Mocks{
		users: make(map[uint]*user_models.User),
	}

	m.users[1] = &user_models.User{
		ID:       1,
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	return m
}

func (m *Mocks) AddUser(user *user_models.User) {
	m.users[user.ID] = &user_models.User{
		ID:       1,
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
}

func (m *Mocks) Vacancy() vacancy_repository.VacancyRepository {
	if m.mockVacancyRepository != nil {
		return m.mockVacancyRepository
	}

	m.mockVacancyRepository = vacancy_mocks_repository.NewMockVacancyRepository(m.vacancy)

	return m.mockVacancyRepository
}

func (m *Mocks) User() user_repository.UserRepository {
	if m.mockUserRepository != nil {
		return m.mockUserRepository
	}

	m.mockUserRepository = user_mock_repository.NewMockUserRepository(m.users)

	return m.mockUserRepository
}
