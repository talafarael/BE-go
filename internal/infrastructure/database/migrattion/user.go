package migrattion

import (
	"gin/internal/user/user_models"
	Migrate2 "gin/pkg/base_migrate"
	"log"
)

type UserMigration struct {
	Migrate Migrate2.Migrate
}

func (u UserMigration) Up() error {
	log.Println("Running UserMigration...")
	return u.Migrate.Migrate(&user_models.User{})
}
