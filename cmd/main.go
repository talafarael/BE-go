package main

import (
	"gin/internal/app"
	load_config "gin/pkg/loadConfig"
	"log"
)

func main() {
	config := load_config.LoadConfig()

	if err := app.NewApp(config).Start(); err != nil {
		log.Fatal(err)
		return
	}
}
