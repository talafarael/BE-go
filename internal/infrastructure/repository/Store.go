package repository

import (
	"gin/internal/user/user_repository"
	"gin/internal/vacancy/vacancy_repository"
)

// Repository implement from interface Store
type Store interface {
	User() user_repository.UserRepository
	Vacancy() vacancy_repository.VacancyRepository
}
