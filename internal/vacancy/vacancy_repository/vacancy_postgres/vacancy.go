package vacancy_postgres

import (
	"gin/internal/user/user_models"
	"gin/internal/vacancy/vacancy_dto"
	"gin/internal/vacancy/vacancy_models"
	"gin/internal/vacancy/vacancy_repository"
	response_error "gin/pkg/error"
	"log"
	"time"

	"gorm.io/gorm"
)

type VacancyRepo struct {
	db *gorm.DB
}

func NewVacancyRepo(db *gorm.DB) vacancy_repository.VacancyRepository {
	return &VacancyRepo{
		db: db,
	}
}

func (v *VacancyRepo) CreateVacancy(user *user_models.User, vacancyDto *vacancy_dto.CreateVacancyDto) (vacancy_models.Vacancy, error) {
	createVacancy := vacancy_models.Vacancy{
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
	err := v.db.Model(&user).Association("Vacancies").Append(&createVacancy)
	if err != nil {
		log.Printf("%s", err)
		return vacancy_models.Vacancy{}, response_error.ErrVacancyCreateError
	}
	return createVacancy, nil
}

func (v *VacancyRepo) UpdateVacancy(user *user_models.User, vacancyDto *vacancy_dto.UpdateVacancyDto, id uint) (vacancy_models.Vacancy, error) {
	updateVacancy := vacancy_models.Vacancy{
		Vacancy:    vacancyDto.Vacancy,
		Status:     vacancyDto.Status,
		Company:    vacancyDto.Company,
		UrlCompany: vacancyDto.UrlCompany,
		UrlChat:    vacancyDto.UrlChat,
		DateMeet:   vacancyDto.DateMeet,
		Time:       time.Now(),
	}
	var updatedVacancy vacancy_models.Vacancy
	err := v.db.Model(&updateVacancy).
		Where("id = ? AND user_id = ?", id, user.ID).
		Updates(updateVacancy).
		Error
	if err != nil {
		return vacancy_models.Vacancy{}, err
	}
	err = v.db.Where("id = ? AND user_id = ?", id, user.ID).First(&updatedVacancy).Error
	if err != nil {
		return vacancy_models.Vacancy{}, err
	}

	return updatedVacancy, nil
}

func (v *VacancyRepo) DeleteVacancy(user *user_models.User, id uint) (bool, error) {
	result := v.db.Where("id = ? AND user_id = ?", id, user.ID).Delete(&vacancy_models.Vacancy{})
	if result.Error != nil {
		return false, response_error.ErrInvalidClaims
	}
	if result.RowsAffected == 0 {
		return false, response_error.ErrVacancyNotFound
	}
	return true, nil
}
