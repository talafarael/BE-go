package vacancy_mocks_repository

import (
	"gin/internal/user/user_models"
	"gin/internal/vacancy/vacancy_dto"
	"gin/internal/vacancy/vacancy_models"
	"gin/internal/vacancy/vacancy_repository"
)

type mockVacancyRepository struct {
	vacancy map[uint]*vacancy_models.Vacancy
}

func NewMockVacancyRepository(vacancy map[uint]*vacancy_models.Vacancy) vacancy_repository.VacancyRepository {
	return &mockVacancyRepository{
		vacancy: vacancy,
	}
}

func (m *mockVacancyRepository) CreateVacancy(user *user_models.User, vacancyDto *vacancy_dto.CreateVacancyDto) (vacancy_models.Vacancy, error) {
	return vacancy_models.Vacancy{}, nil
}

func (v *mockVacancyRepository) UpdateVacancy(user *user_models.User, vacancyDto *vacancy_dto.UpdateVacancyDto, id uint) (vacancy_models.Vacancy, error) {
	return vacancy_models.Vacancy{}, nil
}

func (m *mockVacancyRepository) DeleteVacancy(user *user_models.User, id uint) (bool, error) {
	return true, nil
}
