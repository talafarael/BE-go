package mocks_repository

import (
	"gin/internal/models"
	"gin/internal/repository"
)

type Mocks struct {
	users                 map[uint]*models.User
	mockUserRepository    repository.UserRepository
	vacancy               map[uint]*models.Vacancy
	mockVacancyRepository repository.VacancyRepository
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

func (m *Mocks) Vacancy() repository.VacancyRepository {
	if m.mockVacancyRepository != nil {
		return m.mockVacancyRepository
	}

	m.mockVacancyRepository = NewMockVacancyRepository(m)

	return m.mockVacancyRepository
}

func (m *Mocks) User() repository.UserRepository {
	if m.mockUserRepository != nil {
		return m.mockUserRepository
	}

	m.mockUserRepository = NewMockUserRepository(m)

	return m.mockUserRepository
}
