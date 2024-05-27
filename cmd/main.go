package main

import (
	"backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Create a new Gorilla Mux router instance
	router := mux.NewRouter()

	// Register routes related to Item data
	routes.RouteItemData(router)

	// Set up a server using the Gorilla Mux router
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())

}
