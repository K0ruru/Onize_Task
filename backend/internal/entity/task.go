package entity

import (
	"time"
)

// Album represents an album record.
type Task struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Task_desc   string      `json:"task_desc"`
	Label       string      `json:"label"`
	Priority    string      `json:"priority"`
	Status      string      `json:"status"`
	Tags        string      `json:"tags"`
	Project_id 	string 	    `json:"project_id"`
	CreatedAt   time.Time   `json:"created_at"`
	Due         time.Time   `json:"due"`
}
