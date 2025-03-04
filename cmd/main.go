package main

import (
	"bss.com/create_api_config/internal/controllers"
	"bss.com/create_api_config/internal/helpers"
	"bss.com/create_api_config/internal/ripositories"
	"bss.com/create_api_config/internal/services"
	"log"
	"net/http"
)

// @title Config
// @version 1.0
// @description Documentation de l'API Config.
// @host localhost:8080
// @BasePath /
func main() {
	// Initialisation de la base de données
	db, _ := helpers.InitDB()
	defer db.Close()
	// Définition du routeur
	//router := mux.NewRouter()
	// Initialisation des composants pour Resource
	resourceRepo := &ripositories.ResourceRepository{DB: db}
	resourceService := &services.ResourceService{Repo: resourceRepo}
	resourceController := &controllers.ResourceController{Service: resourceService}

	// Initialisation des composants pour Alert
	alertRepo := &ripositories.AlertRepository{DB: db}
	alertService := &services.AlertService{Repo: alertRepo}
	alertController := &controllers.AlertController{Service: alertService}
	// Routes pour Resource
	http.HandleFunc("/resources", resourceController.GetAll)
	http.HandleFunc("/resources/create", resourceController.Create)
	http.HandleFunc("/resources/get", resourceController.GetByID)
	http.HandleFunc("/resources/update", resourceController.Update)
	http.HandleFunc("/resources/delete", resourceController.Delete)

	// Routes pour Alert
	http.HandleFunc("/alerts", alertController.GetAll)
	http.HandleFunc("/alerts/create", alertController.Create)
	//http.HandleFunc("/alerts/{delete}", alertController.DeleteAlert)
	//http.HandleFunc("/alerts/update", alertController.UpdateAlert)
	log.Println("🚀 Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
