package controllers

import (
	"bss.com/create_api_config/internal/services"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type ResourceController struct {
	Service *services.ResourceService
}

// @Summary Crée une ressource
// @Description Crée une nouvelle ressource avec un nom et un identifiant UCA
// @Tags Ressources
// @Accept json
// @Produce json
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /resources [post]
func (c *ResourceController) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Ucaid int    `json:"uca_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	resource, err := c.Service.CreateResource(req.Name, req.Ucaid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

// @Summary Récupère une ressource par ID
// @Description Retourne une ressource spécifique en fonction de son ID
// @Tags Ressources
// @Accept json
// @Produce json
// @Param id query string true "ID de la ressource"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Resource not found"
// @Router /resources [get]
func (c *ResourceController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	resource, err := c.Service.GetResourceByID(id)
	if err != nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(resource)
}

// Update met à jour une ressource existante
// @Summary Met à jour une ressource
// @Description Met à jour une ressource via son ID
// @Tags Ressources
// @Accept json
// @Produce json
// @Param id query string true "ID de la ressource"
// @Success 200 {string} string "Resource updated successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /resources/{id} [put]
func (c *ResourceController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Name  string `json:"name"`
		Ucaid int    `json:"uca_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateResource(id, req.Name, req.Ucaid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete supprime une ressource
// @Summary Supprime une ressource
// @Description Supprime une ressource via son ID
// @Tags Ressources
// @Accept json
// @Produce json
// @Param id query string true "ID de la ressource"
// @Success 204 {string} string "Resource deleted successfully"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal server error"
// @Router /resources/{id} [delete]
func (c *ResourceController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.Service.DeleteResource(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAll récupère toutes les ressources
// @Summary Liste toutes les ressources
// @Description Récupère la liste de toutes les ressources disponibles
// @Tags Ressources
// @Accept json
// @Produce json
// @Failure 500 {string} string "Internal server error"
// @Router /resources [get]
func (c *ResourceController) GetAll(w http.ResponseWriter, r *http.Request) {
	resources, err := c.Service.GetAllResources()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resources)
}
