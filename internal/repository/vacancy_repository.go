package repository

import (
	"gin/internal/dto"
	"gin/internal/models"
)

type VacancyRepository interface {
	CreateVacancy(user *models.User, vacancyDto *dto.CreateVacancyDto) (models.Vacancy, error)
	UpdateVacancy(user *models.User, vacancyDto *dto.UpdateVacancyDto, id uint) (models.Vacancy, error)
	DeleteVacancy(user *models.User, id uint) (bool, error)
}
