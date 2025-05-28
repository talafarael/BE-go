package models

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	VacancyId []Vacancy `gorm:"foreignKey:UserId"`
}

type UserResponse struct {
	User User `json:"user"`
}
