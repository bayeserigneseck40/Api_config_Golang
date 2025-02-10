package models

type Alert struct {
	ID         string `json:"id" db:"id"`
	Email      string `json:"email" db:"email"` // Liste des destinataires
	ResourceID string `json:"resource_id" db:"resource_id"`
	Oll        string `json:"oll" db:"oll"` // Ex: "always", "room_change", etc.

}
