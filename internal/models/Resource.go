package models

import "github.com/google/uuid"

type Resource struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Name  string    `json:"name" db:"name"`
	UcaId int       `json:"uca_id" db:"uca_id"` // Lien vers le fichier de l'emploi du temps
}
