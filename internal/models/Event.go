package models

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	StartTime  time.Time `json:"start_time" db:"start_time"`
	EndTime    time.Time `json:"end_time" db:"end_time"`
	Location   string    `json:"location" db:"location"`
	ResourceID uuid.UUID `json:"resource_id" db:"resource_id"`
}
