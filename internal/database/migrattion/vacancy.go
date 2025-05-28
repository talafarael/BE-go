package migrattion

import (
	"gin/internal/models"
	Migrate2 "gin/pkg/base_migrate"
	"log"
)

type VacancyMigration struct {
	Migrate Migrate2.Migrate
}

func (u VacancyMigration) Up() error {
	log.Println("Running UserMigration...")
	return u.Migrate.Migrate(&models.Vacancy{})
}
