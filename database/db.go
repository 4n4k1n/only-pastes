package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func getFormatedString() string {
	if err := godotenv.Load(); err != nil {
		return ""
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	formated_string := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		db_host, db_port, db_user, db_password, db_name)

	return formated_string
}

func connectDatabase() error {
	psql_info := getFormatedString()
	if psql_info == "" {
		return errors.New("Failed to load .env")
	}

	db, err := sql.Open("postgrs", psql_info)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func RunMigrations(db *sql.DB) error {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS pastes (
			id SERIAL PRIMARY KEY,
			slug VARCHAR(8) UNIQUE NOT NULL,
			content TEXT NOT NULL,
			language VARCHAR(50),
			expires_at TIMESTAMP,
			views INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW()
		);
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	createIndexSQL := `
		CREATE INDEX IF NOT EXISTS idx_slug ON pastes(slug);
		CREATE INDEX IF NOT EXISTS idx_expires_at ON pastes(expires_at);
	`

	_, err = db.Exec(createIndexSQL)
	if err != nil {
		return err
	}

	return nil
}
