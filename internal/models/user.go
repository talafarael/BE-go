package models

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Vacancies []Vacancy `gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	User User `json:"user"`
}
