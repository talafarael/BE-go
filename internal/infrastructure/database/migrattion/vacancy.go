package migrattion

import (
	"gin/internal/vacancy/vacancy_models"
	Migrate2 "gin/pkg/base_migrate"
	"log"
)

type VacancyMigration struct {
	Migrate Migrate2.Migrate
}

func (v VacancyMigration) Up() error {
	log.Println("Running VacancyMigration...")
	return v.Migrate.Migrate(&vacancy_models.Vacancy{})
}
