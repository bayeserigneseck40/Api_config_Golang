package services

import (
	"bss.com/create_api_config/internal/models"
	"bss.com/create_api_config/internal/ripositories"
	"github.com/google/uuid"
)

type AlertService struct {
	Repo *ripositories.AlertRepository
}

// Créer une alerte
func (s *AlertService) CreateAlert(email string, resourceID uuid.UUID, oll string) (*models.Alert, error) {
	alert := &models.Alert{
		ID:         uuid.New(),
		Email:      email,
		ResourceID: resourceID,
		Oll:        oll,
	}

	err := s.Repo.CreateAlert(alert)
	if err != nil {
		return nil, err
	}

	return alert, nil
}

// Récupérer toutes les alertes
func (s *AlertService) GetAllAlerts() ([]models.Alert, error) {
	return s.Repo.GetAllAlerts()
}
