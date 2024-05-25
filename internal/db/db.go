package db

import (
	"backend/internal/config"
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func InitDB(cfg config.Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	log.Println("Successfully connected to the database!")
}
