package main

import (
	"gin/internal/app"
	"gin/internal/config"
	"gin/pkg/database"
	"log"
)

func main() {
	config := config.Config{
		BindAddr: "8080",
		Connection: database.Config{
			Host:     "localhost",
			Port:     5432,
			User:     "myuser",
			Password: "mypassword",
			DBName:   "mydb",
			SSLMode:  "disable",
			TimeZone: "UTC",
		},
	}

	if err := app.NewApp(config).Start(); err != nil {
		log.Fatal(err)
		return
	}
}
