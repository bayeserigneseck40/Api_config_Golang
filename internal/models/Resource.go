package models

type Resource struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	URL  string `json:"url" db:"url"` // Lien vers le fichier de l'emploi du temps
}
