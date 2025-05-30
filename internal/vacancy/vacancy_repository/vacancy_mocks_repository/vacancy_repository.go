package mocks_vacancy_repository

import (
	"gin/internal/models"
	"gin/internal/repository"
	"gin/internal/vacancy/vacancy_dto"
	"gin/internal/vacancy/vacancy_models"
)

type MockVacancyRepository struct {
	store *Mocks
}

func NewMockVacancyRepository(store *Mocks) repository.VacancyRepository {
	return &mockVacancyRepository{
		store: store,
	}
}

func (m *mockVacancyRepository) CreateVacancy(user *vacancy_models.User, vacancyDto *vacancy_dto.CreateVacancyDto) (models.Vacancy, error) {
	return models.Vacancy{}, nil
}

func (v *mockVacancyRepository) UpdateVacancy(user *models.User, vacancyDto vacancy_dto.UpdateVacancyDto, id uint) (models.Vacancy, error) {
	return models.Vacancy{}, nil
}

func (m *mockVacancyRepository) DeleteVacancy(user *models.User, id uint) (bool, error) {
	return true, nil
}
