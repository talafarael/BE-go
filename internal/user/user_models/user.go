package user_models

import "gin/internal/vacancy/vacancy_models"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Vacancies []vacancy_models.Vacancy `gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	User User `json:"user"`
}
