package models

type Alert struct {
	ID         string `json:"id" db:"id"`
	Recipients string `json:"recipients" db:"recipients"` // Liste des destinataires
	ResourceID string `json:"resource_id" db:"resource_id"`
	AlertType  string `json:"alert_type" db:"alert_type"` // Ex: "always", "room_change", etc.

}
