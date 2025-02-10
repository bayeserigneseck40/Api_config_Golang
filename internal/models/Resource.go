package models

import "github.com/google/uuid"

type Resource struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
	Uid  int       `json:"uid" db:"uid"` // Lien vers le fichier de l'emploi du temps
}
