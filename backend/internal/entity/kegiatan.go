package entity

import (
	"time"
)

type Kegiatan struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Deskripsi  string    `json:"deskripsi"`
	CreatedAt  time.Time `json:"created_at"`
	DeadlineAt time.Time `json:"deadline_at"`
}
