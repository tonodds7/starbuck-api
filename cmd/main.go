package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"log"

	"backend/pkg/tonTest"

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

	tonTest := tonTest.NewTonTestService(db.DB)
	tonTest.DoSomething()

	// userdata.UserDatasController()

	var now string
	err = db.DB.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
	}

	log.Printf("Current time from the database: %s\n", now)
}
