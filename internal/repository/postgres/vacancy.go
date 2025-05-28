package postgres

import (
	"gin/internal/dto"
	"gin/internal/models"
	response_error "gin/pkg/error"
	"time"
)

type VacancyRepo struct {
	store *Repository
}

func (v *VacancyRepo) CreateVacancy(user *models.User, vacancy *dto.CreateVacancyDto) (models.Vacancy, error) {
	createVacancy := models.Vacancy{
		Vacancy:    vacancy.Vacancy,
		Status:     vacancy.Status,
		Company:    vacancy.Company,
		UrlComapny: vacancy.UrlComapny,
		UrlChat:    vacancy.UrlChat,
		DateMeet:   vacancy.DateMeet,
		Time:       time.Now(),
	}
	if err := v.store.db.Model(&user).Association("Vacancies").Append(&createVacancy); err != nil {
		return models.Vacancy{}, response_error.ErrVacancyCreateError
	}
	return createVacancy, nil
}
