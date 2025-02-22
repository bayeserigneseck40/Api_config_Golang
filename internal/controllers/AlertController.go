package controllers

import (
	"bss.com/create_api_config/internal/services"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type AlertController struct {
	Service *services.AlertService
}

// Créer une alerte (POST /alerts/create)
func (c *AlertController) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email      string    `json:"email"`
		ResourceID uuid.UUID `json:"resource_id"`
		Oll        string    `json:"oll"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	alert, err := c.Service.CreateAlert(input.Email, input.ResourceID, input.Oll)
	if err != nil {
		http.Error(w, "Impossible de créer l'alerte", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alert)
}

// Récupérer toutes les alertes (GET /alerts)
func (c *AlertController) GetAll(w http.ResponseWriter, r *http.Request) {
	alerts, err := c.Service.GetAllAlerts()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des alertes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(alerts)
}
