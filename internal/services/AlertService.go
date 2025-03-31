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

// Suppression d'une Alert
func (s *AlertService) DeleteAlert(id uuid.UUID) error {
	return s.Repo.DeleteAlert(id)
}

// Mise à jour d'une ressource
func (s *AlertService) UpdateAlert(id uuid.UUID, email string) error {
	existing, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}

	existing.ID = id
	existing.Email = email

	return s.Repo.UpdateAlert(existing)
}
