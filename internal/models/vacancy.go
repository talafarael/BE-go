package models

import "time"

type Vacancy struct {
	ID         uint `gorm:"primaryKey"`
	Vacancy    string
	Time       time.Time
	Status     bool
	Company    string
	UrlComapny string
	UrlChat    string
	DateMeet   time.Time
}
