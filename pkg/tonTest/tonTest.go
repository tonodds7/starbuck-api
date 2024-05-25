package tonTest

import (
	"database/sql"
	"log"
)

type tonTestService struct {
	DB *sql.DB
}

func NewTonTestService(db *sql.DB) *tonTestService {
	return &tonTestService{DB: db}
}

func (s *tonTestService) DoSomething() {
	var result string
	err := s.DB.QueryRow("SELECT 'service works!'").Scan(&result)
	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
	}

	log.Printf("Service result: %s\n", result)
}
