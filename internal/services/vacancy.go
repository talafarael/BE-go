package services

import (
	"gin/internal/dto"
	"gin/internal/models"
	"gin/internal/repository"
	response_error "gin/pkg/error"

	"github.com/gin-gonic/gin"
)

type vacancyService struct {
	VacancyServiceOptions
}
type VacancyService interface {
	CreateVacancy(ctx *gin.Context, vacancy *dto.CreateVacancyDto) (bool, error)
	UpdateVacancy(ctx *gin.Context, vacancy *dto.UpdateVacancyDto, id uint) (bool, error)
	DeleteVacancy(ctx *gin.Context, id uint) (bool, error)
}

type VacancyServiceOptions struct {
	Repo repository.Store
}

func NewVacancyService(options VacancyServiceOptions) VacancyService {
	return &vacancyService{
		VacancyServiceOptions: options,
	}
}

func (v *vacancyService) CreateVacancy(ctx *gin.Context, vacancy *dto.CreateVacancyDto) (bool, error) {
	userValue, exists := ctx.Get("user")
	if !exists {
		return false, response_error.ErrUnauthorized
	}
	user, ok := userValue.(*models.User)
	if !ok {
		return false, response_error.ErrUnauthorized
	}

	_, err := v.Repo.Vacancy().CreateVacancy(user, vacancy)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (v *vacancyService) UpdateVacancy(ctx *gin.Context, vacancy *dto.UpdateVacancyDto, id uint) (bool, error) {
	userValue, exists := ctx.Get("user")
	if !exists {
		return false, response_error.ErrUnauthorized
	}
	user, ok := userValue.(*models.User)
	if !ok {
		return false, response_error.ErrUnauthorized
	}
	_, err := v.Repo.Vacancy().UpdateVacancy(user, vacancy, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (v *vacancyService) DeleteVacancy(ctx *gin.Context, id uint) (bool, error) {
	userValue, exists := ctx.Get("user")
	if !exists {
		return false, response_error.ErrUnauthorized
	}
	user, ok := userValue.(*models.User)
	if !ok {
		return false, response_error.ErrUnauthorized
	}
	_, err := v.Repo.Vacancy().DeleteVacancy(user, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
