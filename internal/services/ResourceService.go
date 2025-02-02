package services

import (
	"bss.com/create_api_config/internal/models"
	"bss.com/create_api_config/internal/ripositories"
	"errors"
	"github.com/google/uuid"
)

type ResourceService struct {
	Repo *ripositories.ResourceRepository
}

// Création d'une ressource avec validation
func (s *ResourceService) CreateResource(name, url string) (*models.Resource, error) {
	if name == "" || url == "" {
		return nil, errors.New("name and URL cannot be empty")
	}

	resource := &models.Resource{
		ID:   "4",
		Name: name,
		URL:  url,
	}

	err := s.Repo.Create(resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

// Récupération d'une ressource
func (s *ResourceService) GetResourceByID(id uuid.UUID) (*models.Resource, error) {
	return s.Repo.GetByID(id)
}

// Mise à jour d'une ressource
func (s *ResourceService) UpdateResource(id uuid.UUID, name, url string) error {
	existing, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}

	existing.Name = name
	existing.URL = url

	return s.Repo.Update(existing)
}

// Suppression d'une ressource
func (s *ResourceService) DeleteResource(id uuid.UUID) error {
	return s.Repo.Delete(id)
}

// Récupération de toutes les ressources
func (s *ResourceService) GetAllResources() ([]models.Resource, error) {
	return s.Repo.GetAll()
}
