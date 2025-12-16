package models

import "time"

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
