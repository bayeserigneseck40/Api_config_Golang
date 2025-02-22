package models

import "github.com/google/uuid"

type Alert struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"` // Liste des destinataires
	ResourceID uuid.UUID `json:"resource_id" db:"resource_id"`
	Oll        string    `json:"oll" db:"oll"` // Ex: "always", "room_change", etc.

}
