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

// Création d'une ressource
func (c *ResourceController) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	resource, err := c.Service.CreateResource(req.Name, req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

// Récupération d'une ressource
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

// Mise à jour d'une ressource
func (c *ResourceController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateResource(id, req.Name, req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Suppression d'une ressource
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

// Récupération de toutes les ressources
func (c *ResourceController) GetAll(w http.ResponseWriter, r *http.Request) {
	resources, err := c.Service.GetAllResources()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resources)
}
