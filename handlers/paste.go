package handlers

import (
	"database/sql"
	"log"
	"math/rand"
	"only-pastes/database"
	"only-pastes/models"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func generateSlug() string {
	char_set := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lenght := 6
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

	expires_at := calculateExpiration(request.ExpiresIn)

	query := `
        INSERT INTO pastes (slug, content, language, expires_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
	`

	var id int
	var created_at time.Time
	var slug string
	var err error

	maxRetries := 100
	for attempt := 0; attempt < maxRetries; attempt++ {
		slug = generateSlug()
		err = database.DB.QueryRow(query, slug, request.Content, request.Language, expires_at).Scan(&id, &created_at)

		if err == nil {
			break
		}

		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			log.Printf("Slug collision detected (attempt %d/%d): %s", attempt+1, maxRetries, slug)
			continue
		}

		log.Println("Database error:", err)
		ctx.JSON(500, gin.H{"error": "Failed to create paste"})
		return
	}

	if err != nil {
		log.Println("Failed to generate unique slug after", maxRetries, "attempts")
		ctx.JSON(500, gin.H{"error": "Failed to generate unique identifier"})
		return
	}

	base_url := os.Getenv("BASE_URL")
	ctx.JSON(201, gin.H{
		"slug": slug,
		"url":  base_url + "/" + slug,
	})
}

func GetPaste(ctx *gin.Context) {
	slug := ctx.Param("slug")

	query := `SELECT id, slug, content, language, expires_at, views, created_at 
            FROM pastes 
            WHERE slug = $1`

	var paste models.Paste

	err := database.DB.QueryRow(query, slug).Scan(
		&paste.ID,
		&paste.Slug,
		&paste.Content,
		&paste.Language,
		&paste.ExpiresAt,
		&paste.Views,
		&paste.CreatedAt,
	)

	if err == sql.ErrNoRows {
		ctx.JSON(404, gin.H{"error": "Paste not found"})
		return
	}
	if err != nil {
		log.Println("Database error:", err)
		ctx.JSON(500, gin.H{"error": "Failed to retrieve paste"})
		return
	}

	if paste.ExpiresAt != nil && paste.ExpiresAt.Before(time.Now()) {
		ctx.JSON(404, gin.H{"error": "Paste has expired"})
		return
	}

	database.DB.Exec("UPDATE pastes SET views = views + 1 WHERE slug = $1", slug)

	ctx.JSON(200, paste)
}
