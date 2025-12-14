package main

import (
	"log"
	"only-pastes/database"
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

	// Add routes later here

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
