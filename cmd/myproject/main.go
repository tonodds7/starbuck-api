package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"log"

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

	// tonTestService := tonTestService.NewTonTestService(db.DB)

	// Example query to test the connection
	var now string
	err = db.DB.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
	}

	log.Printf("Current time from the database: %s\n", now)
}
