package services

import (
	"bss.com/create_api_config/internal/models"
	"bss.com/create_api_config/internal/ripositories"
)

type EventService struct {
	Repo *ripositories.EventRepository
}

func NewEventService(repo *ripositories.EventRepository) *EventService {
	return &EventService{Repo: repo}
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
	return s.Repo.GetAllEvents()
}

func (s *EventService) GetEventByID(id string) (*models.Event, error) {
	return s.Repo.GetEventByID(id)
}

func (s *EventService) CreateEvent(event *models.Event) error {
	return s.Repo.CreateEvent(event)
}

func (s *EventService) UpdateEvent(event *models.Event) error {
	return s.Repo.UpdateEvent(event)
}

func (s *EventService) DeleteEvent(id string) error {
	return s.Repo.DeleteEvent(id)
}
