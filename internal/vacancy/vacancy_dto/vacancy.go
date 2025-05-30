package vacancy_dto

import "time"

type CreateVacancyDto struct {
	Vacancy    string    `json:"vacancy"`
	Status     bool      `json:"status"`
	Company    string    `json:"company"`
	UrlCompany string    `json:"url_company"`
	UrlChat    string    `json:"url_chat"`
	DateMeet   time.Time `json:"date_meet"`
}

type UpdateVacancyDto struct {
	Vacancy    string    `json:"vacancy"`
	Status     bool      `json:"status"`
	Company    string    `json:"company"`
	UrlCompany string    `json:"url_company"`
	UrlChat    string    `json:"url_chat"`
	DateMeet   time.Time `json:"date_meet"`
}
