package app

import (
	"database/sql"
	"gin/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

type App struct {
	db     *sql.DB
	router *gin.Engine
}

func (s *App) configureRouter() {
	// Example route: a GET endpoint to check if the server is up
	s.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	s.router.Use()
	// You can add more routes here, such as routes for accessing your database, handling POST requests, etc.
}

func Start() {
	router := gin.Default()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app := &App{db: db, router: router}

	app.configureRouter()
	app.router.Run(":8080")
}
