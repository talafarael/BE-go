package mocks_repository

import (
	"gin/internal/dto"
	"gin/internal/models"
	"gin/internal/repository"
)

type mockVacancyRepository struct {
	store *Mocks
}

func NewMcokVacancyRepository(store *Mocks) repository.VacancyRepository {
	return &mockVacancyRepository{
		store: store,
	}
}
func (m *mockVacancyRepository) CreateVacancy(user *models.User, vacancy *dto.CreateVacancyDto) (models.Vacancy, error)
