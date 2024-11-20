package main

import (
	"backend-boilerplate/models"
	"backend-boilerplate/routes"
	"backend-boilerplate/utils"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	utils.InitDB()

	// Run migrations
	models.MigrateUser()

	// Setup routes
	router := routes.InitializeRoutes()

	log.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", router)
}
