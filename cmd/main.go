package main

import (
	"gin/internal/app"
	"gin/pkg/load_config"
	"log"
)

func main() {
	config := load_config.LoadConfig()

	if err := app.NewApp(config).Start(); err != nil {
		log.Fatal(err)
		return
	}
}
