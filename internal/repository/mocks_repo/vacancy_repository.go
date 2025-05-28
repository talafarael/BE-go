package mocks_repository

import (
	"gin/internal/dto"
	"gin/internal/models"
	"gin/internal/repository"
)

type mockVacancyRepository struct {
	store *Mocks
}

func NewMockVacancyRepository(store *Mocks) repository.VacancyRepository {
	return &mockVacancyRepository{
		store: store,
	}
}

func (m *mockVacancyRepository) CreateVacancy(user *models.User, vacancyDto *dto.CreateVacancyDto) (models.Vacancy, error) {
	return models.Vacancy{}, nil
}

func (v *mockVacancyRepository) UpdateVacancy(user *models.User, vacancyDto *dto.UpdateVacancyDto, id uint) (models.Vacancy, error) {
	return models.Vacancy{}, nil
}

func (m *mockVacancyRepository) DeleteVacancy(user *models.User, id uint) (bool, error) {
	return true, nil
}
