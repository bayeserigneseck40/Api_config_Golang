package services

import (
	"bss.com/create_api_config/internal/models"
	"bss.com/create_api_config/internal/ripositories"
)

type AlertService struct {
	Repo *ripositories.AlertRepository
}

func NewAlertService(repo *ripositories.AlertRepository) *AlertService {
	return &AlertService{Repo: repo}
}

func (s *AlertService) GetAllAlerts() ([]models.Alert, error) {
	return s.Repo.GetAllAlerts()
}

func (s *AlertService) GetAlertByID(id string) (*models.Alert, error) {
	return s.Repo.GetAlertByID(id)
}

func (s *AlertService) CreateAlert(alert *models.Alert) error {
	return s.Repo.CreateAlert(alert)
}

func (s *AlertService) UpdateAlert(alert *models.Alert) error {
	return s.Repo.UpdateAlert(alert)
}

func (s *AlertService) DeleteAlert(id string) error {
	return s.Repo.DeleteAlert(id)
}
