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
	db := helpers.InitDB()
	defer db.Close()

	repo := &ripositories.ResourceRepository{DB: db}
	service := &services.ResourceService{Repo: repo}
	controller := &controllers.ResourceController{Service: service}

	http.HandleFunc("/resources", controller.GetAll)
	http.HandleFunc("/resources/create", controller.Create)
	http.HandleFunc("/resources/get", controller.GetByID)
	http.HandleFunc("/resources/update", controller.Update)
	http.HandleFunc("/resources/delete", controller.Delete)

	log.Println("🚀 Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
