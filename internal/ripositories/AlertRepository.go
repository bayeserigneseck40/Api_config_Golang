package ripositories

import (
	"bss.com/create_api_config/internal/models"
	"database/sql"
)

type AlertRepository struct {
	DB *sql.DB
}

func NewAlertRepository(db *sql.DB) *AlertRepository {
	return &AlertRepository{DB: db}
}

// Récupérer tous les événements
func (r *AlertRepository) GetAllAlerts() ([]models.Alert, error) {
	rows, err := r.DB.Query("SELECT id, recipient, resource_id, alert_type FROM alerts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []models.Alert
	for rows.Next() {
		var alert models.Alert
		if err := rows.Scan(&alert.ID, &alert.Recipients, &alert.ResourceID, &alert.AlertType); err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

// Récupérer un événement par ID
func (r *AlertRepository) GetAlertByID(id string) (*models.Alert, error) {
	var alert models.Alert
	err := r.DB.QueryRow("SELECT id, recipient, resource_id, alert_type FROM alerts WHERE id = ?", id).
		Scan(&alert.ID, &alert.Recipients, &alert.ResourceID, &alert.AlertType)

	if err != nil {
		return nil, err
	}
	return &alert, nil
}

// Créer un nouvel événement
func (r *AlertRepository) CreateAlert(alert *models.Alert) error {
	_, err := r.DB.Exec("INSERT INTO alerts (id, recipient, resource_id, alert_type) VALUES (?, ?, ?, ?)",
		alert.ID, alert.Recipients, alert.ResourceID, alert.AlertType)
	return err
}

// Mettre à jour un événement
func (r *AlertRepository) UpdateAlert(alert *models.Alert) error {
	_, err := r.DB.Exec("UPDATE alerts SET recipient = ? , resource_id = ?,alert_type = ? WHERE id = ?",
		alert.Recipients, alert.ResourceID, alert.AlertType, alert.ID)
	return err
}

// Supprimer un événement
func (r *AlertRepository) DeleteAlert(id string) error {
	_, err := r.DB.Exec("DELETE FROM alerts WHERE id = ?", id)
	return err
}
