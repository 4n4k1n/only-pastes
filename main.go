package main

import (
	"log"
	"only-pastes/database"
	_ "only-pastes/docs"
	"only-pastes/handlers"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Only Pastes API
// @version 1.0
// @description A simple pastebin service API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@onlypastes.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @BasePath /api

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/api/paste", handlers.CreatePaste)
	router.GET("/api/paste/:slug", handlers.GetPaste)
	
	router.Static("/static", "./static")
	router.StaticFile("/", "./static/index.html")

	router.GET("/:slug", func(c *gin.Context) {
		c.File("./static/view.html")
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
