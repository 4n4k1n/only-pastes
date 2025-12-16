package main

import (
	"log"
	"only-pastes/database"
	"only-pastes/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	router := gin.Default()

	router.POST("/api/paste", handlers.CreatePaste)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
