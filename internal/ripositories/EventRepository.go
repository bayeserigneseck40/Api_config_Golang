package ripositories

import (
	"bss.com/create_api_config/internal/models"
	"database/sql"
)

type EventRepository struct {
	DB *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{DB: db}
}

// Récupérer tous les événements
func (r *EventRepository) GetAllEvents() ([]models.Event, error) {
	rows, err := r.DB.Query("SELECT id, title, start_time, end_time, location, resource_id FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ID, &event.Title, &event.StartTime, &event.EndTime, &event.Location, &event.ResourceID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

// Récupérer un événement par ID
func (r *EventRepository) GetEventByID(id string) (*models.Event, error) {
	var event models.Event
	err := r.DB.QueryRow("SELECT id, title, start_time, end_time, location, resource_id FROM events WHERE id = ?", id).
		Scan(&event.ID, &event.Title, &event.StartTime, &event.EndTime, &event.Location, &event.ResourceID)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Créer un nouvel événement
func (r *EventRepository) CreateEvent(event *models.Event) error {
	_, err := r.DB.Exec("INSERT INTO events (id, title, start_time, end_time, location, resource_id) VALUES (?, ?, ?, ?, ?, ?)",
		event.ID, event.Title, event.StartTime, event.EndTime, event.Location, event.ResourceID)
	return err
}

// Mettre à jour un événement
func (r *EventRepository) UpdateEvent(event *models.Event) error {
	_, err := r.DB.Exec("UPDATE events SET title = ?, start_time = ?, end_time = ?, location = ?, resource_id = ? WHERE id = ?",
		event.Title, event.StartTime, event.EndTime, event.Location, event.ResourceID, event.ID)
	return err
}

// Supprimer un événement
func (r *EventRepository) DeleteEvent(id string) error {
	_, err := r.DB.Exec("DELETE FROM events WHERE id = ?", id)
	return err
}
