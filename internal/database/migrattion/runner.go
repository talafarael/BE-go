package migrattion

import (
	"gin/pkg/base_migrate"
	"gin/pkg/database"
	"log"
)

func RunMigrations(store database.Database) {
	log.Printf("migration")

	if store == nil {
		log.Fatal("Store cannot be nil")
		return
	}
	migrator := base_migrate.NewMigrate(store)

	migrations := getMigrations(migrator)
	for _, migration := range migrations {
		err := migration.Up()
		if err != nil {
			log.Printf("Migration failed: %v\n", err)
			continue
		}
		log.Println("Migration succeeded!")
	}
}
