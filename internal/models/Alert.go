package models

import (
	"github.com/google/uuid"
	"time"
)

type Alert struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Recipients []string  `json:"recipients" db:"recipients"` // Liste des destinataires
	ResourceID uuid.UUID `json:"resource_id" db:"resource_id"`
	Condition  string    `json:"condition" db:"condition"` // Ex: "always", "room_change", etc.
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
