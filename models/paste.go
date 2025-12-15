package models

import (
	"math/rand"
	"only-pastes/database"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type Paste struct {
	ID        int
	Slug      string     `json:"slug"`
	Content   string     `json:"content"`
	Language  string     `json:"language"`
	ExpiresAt *time.Time `json:"expires_at"`
	Views     int        `json:"views"`
	CreatedAt time.Time  `json:"created_at"`
}

type CreatePasteRequest struct {
	Content   string `json:"content"`
	Language  string `json:"language"`
	ExpiresIn string `json:"expires_in"`
}

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
	var request CreatePasteRequest

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
	`

	var id int
	var created_at time.Time

	err := database.DB.QueryRow(query, slug, request.Content, request.Language, expires_at).Scan(&id, &created_at)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create paste"})
        return
	}

	base_url := os.Getenv("BASE_URL")
	ctx.JSON(201, gin.H{
        "slug": slug,
        "url": base_url + "/" + slug,
    })
}
