package vacancy_postgres

import (
	"gin/internal/models"
	"gin/internal/vacancy_dto"
	response_error "gin/pkg/error"
	"log"
	"time"
)

type VacancyRepo struct {
	store *Repository
}

func (v *VacancyRepo) CreateVacancy(user *models.User, vacancyDto *vacancy_dto.CreateVacancyDto) (models.Vacancy, error) {
	createVacancy := models.Vacancy{
		Vacancy:    vacancyDto.Vacancy,
		Status:     vacancyDto.Status,
		Company:    vacancyDto.Company,
		UrlCompany: vacancyDto.UrlCompany,
		UrlChat:    vacancyDto.UrlChat,
		DateMeet:   vacancyDto.DateMeet,
		Time:       time.Now(),
		UserID:     user.ID,
	}
	log.Printf("create")
	err := v.store.db.Model(&user).Association("Vacancies").Append(&createVacancy)
	if err != nil {
		log.Printf("%s", err)
		return models.Vacancy{}, response_error.ErrVacancyCreateError
	}
	return createVacancy, nil
}

func (v *VacancyRepo) UpdateVacancy(user *models.User, vacancyDto *vacancy_dto.UpdateVacancyDto, id uint) (models.Vacancy, error) {
	updateVacancy := models.Vacancy{
		Vacancy:    vacancyDto.Vacancy,
		Status:     vacancyDto.Status,
		Company:    vacancyDto.Company,
		UrlCompany: vacancyDto.UrlCompany,
		UrlChat:    vacancyDto.UrlChat,
		DateMeet:   vacancyDto.DateMeet,
		Time:       time.Now(),
	}
	var updatedVacancy models.Vacancy
	err := v.store.db.Model(&updateVacancy).
		Where("id = ? AND user_id = ?", id, user.ID).
		Updates(updateVacancy).
		Error
	if err != nil {
		return models.Vacancy{}, err
	}
	err = v.store.db.Where("id = ? AND user_id = ?", id, user.ID).First(&updatedVacancy).Error
	if err != nil {
		return models.Vacancy{}, err
	}

	return updatedVacancy, nil
}

func (v *VacancyRepo) DeleteVacancy(user *models.User, id uint) (bool, error) {
	result := v.store.db.Where("id = ? AND user_id = ?", id, user.ID).Delete(&models.Vacancy{})
	if result.Error != nil {
		return false, response_error.ErrInvalidClaims
	}
	if result.RowsAffected == 0 {
		return false, response_error.ErrVacancyNotFound
	}
	return true, nil
}
