package vacancy_repository

import (
	"gin/internal/user/user_models"
	"gin/internal/vacancy/vacancy_dto"
	"gin/internal/vacancy/vacancy_models"
)

type VacancyRepository interface {
	CreateVacancy(user *user_models.User, vacancyDto *vacancy_dto.CreateVacancyDto) (vacancy_models.Vacancy, error)
	UpdateVacancy(user *user_models.User, vacancyDto *vacancy_dto.UpdateVacancyDto, id uint) (vacancy_models.Vacancy, error)
	DeleteVacancy(user *user_models.User, id uint) (bool, error)
}
