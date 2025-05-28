package postgres

import (
	"gin/internal/repository"

	"gorm.io/gorm"
)

type Repository struct {
	db                *gorm.DB
	userRepository    *UserRepo
	vacancyRepository *VacancyRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Vacancy() repository.VacancyRepository {
	if r.vacancyRepository != nil {
		return r.vacancyRepository
	}

	r.vacancyRepository = &VacancyRepo{
		store: r,
	}

	return r.vacancyRepository
}

func (r *Repository) User() repository.UserRepository {
	if r.userRepository != nil {
		return r.userRepository
	}

	r.userRepository = &UserRepo{
		store: r,
	}

	return r.userRepository
}
