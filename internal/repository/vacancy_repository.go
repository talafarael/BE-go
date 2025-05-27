package repository

import "gin/internal/dto"

type VacancyRepository interface {
	CreateVacancy(vacancy *dto.LoginDto)
}
