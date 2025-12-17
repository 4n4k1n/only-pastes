package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func getFormatedString() string {
	godotenv.Load()

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	if !db_host || !db_port || !db_user || !db_password || !db_name {
		return ""
	}

	formated_string := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db_host, db_port, db_user, db_password, db_name)

	return formated_string
}

func ConnectDatabase() (*sql.DB, error) {
	psql_info := getFormatedString()
	if psql_info == "" {
		return nil, errors.New("Failed to load database configuration: required environment variables not set")
	}

	db, err := sql.Open("postgres", psql_info)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
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
