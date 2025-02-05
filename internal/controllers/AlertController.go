package controllers

import (
	"bss.com/create_api_config/internal/models"
	"bss.com/create_api_config/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type AlertController struct {
	Service *services.AlertService
}

func NewAlertController(service *services.AlertService) *AlertController {
	return &AlertController{Service: service}
}

// Récupérer tous les événements
func (c *AlertController) GetAllAlerts(w http.ResponseWriter, r *http.Request) {
	alerts, err := c.Service.GetAllAlerts()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des alerts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alerts)
}

// Récupérer un événement par ID
func (c *AlertController) GetAlertByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	alert, err := c.Service.GetAlertByID(id)
	if err != nil {
		http.Error(w, "Alert non trouvé", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alert)
}

// Créer une nouvelle Alert
func (c *AlertController) CreateAlert(w http.ResponseWriter, r *http.Request) {
	var alert models.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	err := c.Service.CreateAlert(&alert)
	if err != nil {
		http.Error(w, "Impossible de créer l'alert", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alert)
}

// Mettre à jour un événement
func (c *AlertController) UpdateAlert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var alert models.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	alert.ID = id // Assigner l'ID

	err := c.Service.UpdateAlert(&alert)
	if err != nil {
		http.Error(w, "Impossible de mettre à jour l'alert", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alert)
}

// Supprimer un événement
func (c *AlertController) DeleteAlert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := c.Service.DeleteAlert(id)
	if err != nil {
		http.Error(w, "Alert non trouvé", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
