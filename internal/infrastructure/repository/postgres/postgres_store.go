package postgres

import (
	"gin/internal/user/user_repository"
	"gin/internal/user/user_repository/user_postgres"
	"gin/internal/vacancy/vacancy_repository"
	"gin/internal/vacancy/vacancy_repository/vacancy_postgres"

	"gorm.io/gorm"
)

type Repository struct {
	db                *gorm.DB
	userRepository    user_repository.UserRepository
	vacancyRepository vacancy_repository.VacancyRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Vacancy() vacancy_repository.VacancyRepository {
	if r.vacancyRepository != nil {
		return r.vacancyRepository
	}

	r.vacancyRepository = vacancy_postgres.NewVacancyRepo(r.db)

	return r.vacancyRepository
}

func (r *Repository) User() user_repository.UserRepository {
	if r.userRepository != nil {
		return r.userRepository
	}

	r.userRepository = user_postgres.NewUserRepo(r.db)
	return r.userRepository
}
