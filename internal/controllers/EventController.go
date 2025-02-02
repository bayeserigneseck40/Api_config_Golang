package controllers

import (
	"bss.com/create_api_config/internal/models"
	"bss.com/create_api_config/internal/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type EventController struct {
	Service *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{Service: service}
}

// Récupérer tous les événements
func (c *EventController) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := c.Service.GetAllEvents()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des événements", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// Récupérer un événement par ID
func (c *EventController) GetEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	event, err := c.Service.GetEventByID(id)
	if err != nil {
		http.Error(w, "Événement non trouvé", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// Créer un nouvel événement
func (c *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	err := c.Service.CreateEvent(&event)
	if err != nil {
		http.Error(w, "Impossible de créer l'événement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

// Mettre à jour un événement
func (c *EventController) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	event.ID = id // Assigner l'ID

	err := c.Service.UpdateEvent(&event)
	if err != nil {
		http.Error(w, "Impossible de mettre à jour l'événement", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// Supprimer un événement
func (c *EventController) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := c.Service.DeleteEvent(id)
	if err != nil {
		http.Error(w, "Événement non trouvé", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
