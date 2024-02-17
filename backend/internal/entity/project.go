package entity

import (
	"time"
)

// Album represents an album record.
type Project struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	User_id 	string 		`json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
