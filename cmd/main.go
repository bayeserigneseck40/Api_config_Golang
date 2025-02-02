package main

import (
	"bss.com/create_api_config/internal/controllers"
	"bss.com/create_api_config/internal/helpers"
	"bss.com/create_api_config/internal/ripositories"
	"bss.com/create_api_config/internal/services"
	"log"
	"net/http"
)

func main() {
	// Initialisation de la base de données
	db, _ := helpers.InitDB()
	defer db.Close()

	// Initialisation des composants pour Resource
	resourceRepo := &ripositories.ResourceRepository{DB: db}
	resourceService := &services.ResourceService{Repo: resourceRepo}
	resourceController := &controllers.ResourceController{Service: resourceService}

	// Initialisation des composants pour Event
	eventRepo := &ripositories.EventRepository{DB: db}
	eventService := &services.EventService{Repo: eventRepo}
	eventController := &controllers.EventController{Service: eventService}

	// Routes pour Resource
	http.HandleFunc("/resources", resourceController.GetAll)
	http.HandleFunc("/resources/create", resourceController.Create)
	http.HandleFunc("/resources/get", resourceController.GetByID)
	http.HandleFunc("/resources/update", resourceController.Update)
	http.HandleFunc("/resources/delete", resourceController.Delete)

	// Routes pour Event
	http.HandleFunc("/events", eventController.GetAllEvents)
	http.HandleFunc("/events/create", eventController.CreateEvent)
	http.HandleFunc("/events/get", eventController.GetEventByID)
	http.HandleFunc("/events/update", eventController.UpdateEvent)
	http.HandleFunc("/events/delete", eventController.DeleteEvent)

	log.Println("🚀 Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
