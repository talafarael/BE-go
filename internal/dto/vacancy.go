package dto

import "time"

type CreateVacancyDto struct {
	Vacancy    string
	Status     bool
	Company    string
	UrlComapny string
	UrlChat    string
	DateMeet   time.Time
}
