package dto

import "time"

type CreateVacancy struct {
	Vacancy    string
	Status     bool
	Company    string
	UrlComapny string
	UrlChat    string
	DateMeet   time.Time
}
