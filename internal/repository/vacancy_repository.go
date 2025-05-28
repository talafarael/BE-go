package repository

import (
	"gin/internal/dto"
	"gin/internal/models"
)

type VacancyRepository interface {
	CreateVacancy(user *models.User, vacancy *dto.CreateVacancyDto) (models.Vacancy, error)
}
