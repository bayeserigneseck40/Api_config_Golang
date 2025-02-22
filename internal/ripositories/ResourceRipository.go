package ripositories

import (
	"bss.com/create_api_config/internal/models"
	"database/sql"
	"github.com/google/uuid"
)

type ResourceRepository struct {
	DB *sql.DB
}

// Crée une nouvelle ressource
func (r *ResourceRepository) Create(resource *models.Resource) error {
	query := "INSERT INTO resources (id, name, uca_id) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, resource.ID, resource.Name, resource.UcaId)
	return err
}

// Récupère une ressource par son ID
func (r *ResourceRepository) GetByID(id uuid.UUID) (*models.Resource, error) {
	query := "SELECT id, name, uca_id FROM resources WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	resource := &models.Resource{}
	err := row.Scan(&resource.ID, &resource.Name, &resource.UcaId)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

// Met à jour une ressource
func (r *ResourceRepository) Update(resource *models.Resource) error {
	query := "UPDATE resources SET name = ?, uca_id = ? WHERE id = ?"
	_, err := r.DB.Exec(query, resource.Name, resource.UcaId, resource.ID)
	return err
}

// Supprime une ressource
func (r *ResourceRepository) Delete(id uuid.UUID) error {
	query := "DELETE FROM resources WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}

// Récupère toutes les ressources
func (r *ResourceRepository) GetAll() ([]models.Resource, error) {
	query := "SELECT id, name, uca_id FROM resources"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []models.Resource
	for rows.Next() {
		var resource models.Resource
		if err := rows.Scan(&resource.ID, &resource.Name, &resource.UcaId); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}
