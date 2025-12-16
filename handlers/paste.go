package handlers

import (
	"log"
	"math/rand"
	"only-pastes/database"
	"only-pastes/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func generateSlug() string {
	char_set := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lenght := 8
	result := make([]byte, lenght)

	for i := 0; i < lenght; i++ {
		result[i] = char_set[rand.Intn(len(char_set))]
	}

	return string(result)
}

func calculateExpiration(expires_in string) *time.Time {
	if expires_in == "never" {
		return nil
	}
	if expires_in == "1h" {
		t := time.Now().Add(1 * time.Hour)
		return (&t)
	}
	if expires_in == "1d" {
		t := time.Now().Add(24 * time.Hour)
		return (&t)
	}
	if expires_in == "1w" {
		t := time.Now().Add(7 * 24 * time.Hour)
		return (&t)
	}
	return nil
}

func CreatePaste(ctx *gin.Context) {
	var request models.CreatePasteRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if request.Content == "" {
		ctx.JSON(400, gin.H{"error": "Content cannot be empty"})
		return
	}

	slug := generateSlug()
	expires_at := calculateExpiration(request.ExpiresIn)

	query := `
        INSERT INTO pastes (slug, content, language, expires_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
	`

	var id int
	var created_at time.Time

	err := database.DB.QueryRow(query, slug, request.Content, request.Language, expires_at).Scan(&id, &created_at)

	if err != nil {
		log.Println("Database error:", err)
		ctx.JSON(500, gin.H{"error": "Failed to create paste"})
		return
	}

	base_url := os.Getenv("BASE_URL")
	ctx.JSON(201, gin.H{
		"slug": slug,
		"url":  base_url + "/" + slug,
	})
}
