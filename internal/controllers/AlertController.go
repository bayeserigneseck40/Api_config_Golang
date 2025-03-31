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

// Create crée une nouvelle alerte
// @Summary Crée une alerte
// @Description Crée une nouvelle alerte associée à une ressource
// @Tags Alertes
// @Accept json
// @Produce json
// @Failure 400 {string} string "Données invalides"
// @Failure 500 {string} string "Impossible de créer l'alerte"
// @Router /alerts [post]
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

// GetAll récupère toutes les alertes
// @Summary Liste toutes les alertes
// @Description Récupère toutes les alertes enregistrées
// @Tags Alertes
// @Accept json
// @Produce json
// @Failure 500 {string} string "Erreur lors de la récupération des alertes"
// @Router /alerts [get]
func (c *AlertController) GetAll(w http.ResponseWriter, r *http.Request) {
	alerts, err := c.Service.GetAllAlerts()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des alertes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(alerts)
}

func (c *AlertController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.Service.DeleteAlert(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (c *AlertController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req struct {
		email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateAlert(id, req.email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
