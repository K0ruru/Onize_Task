package entity

import (
	"time"

	"github.com/lib/pq"
)

// Album represents an album record.
type Task struct {
	ID         string         `json:"id"`
	Title      string         `json:"title"`
	Task_desc  string         `json:"task_desc"`
	Label      pq.StringArray `json:"label" gorm:"type:text[]"`
	Priority   string         `json:"priority"`
	Status     string         `json:"status"`
	Tags       pq.StringArray `json:"tags" gorm:"type:text[]"`
	Project_id string         `json:"project_id"`
	CreatedAt  time.Time      `json:"created_at"`
	Due        time.Time      `json:"due"`
}
