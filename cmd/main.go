package main

import (
	"database/sql"
	"gin/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

type Server struct {
	db     *sql.DB
	router *gin.Engine
}

func (s *Server) SetupRoutes() {
	// Example route: a GET endpoint to check if the server is up
	s.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// You can add more routes here, such as routes for accessing your database, handling POST requests, etc.
}

func main() {
	router := gin.Default()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	server := &Server{db: db, router: router}

	server.SetupRoutes()

	server.router.Run(":8080")
}
