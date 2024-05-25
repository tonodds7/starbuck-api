package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.LoadConfig()

	db.InitDB(cfg)
	defer db.DB.Close()

	// Create a new Gorilla Mux router instance
	router := mux.NewRouter()

	// Register routes related to user data
	routes.RegisterRouteUserDate(router)

	// Set up a server using the Gorilla Mux router
	server := &http.Server{
		Addr:    ":8080", // Change port number if needed
		Handler: router,
	}

	// Start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())

	var now string
	err = db.DB.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
	}

	log.Printf("Current time from the database: %s\n", now)
}
