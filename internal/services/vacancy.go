package services

import (
	"gin/internal/dto"
	"gin/internal/repository"
)

type vacancyService struct {
	VacancyServiceOptions
}
type VacancyService interface{}

type VacancyServiceOptions struct {
	Repo repository.Store
}

func NewVacancyService() VacancyService {
	return &vacancyService{}
}

func (v *vacancyService) CreateVacancy(vacancy *dto.CreateVacancyDto) {
	// vacancy, err := v.Repo..VacancyRepo
}
