package routes

import (
	"backend/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouteItemData(router *mux.Router) {
	itemDataHandler := handler.NewItemDataHandler()

	groupItemData := router.PathPrefix("/itemdatas").Subrouter()

	// groupItemData.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	itemDataHandler.ItemDatasController(r)
	// }).Methods("GET")

	groupItemData.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		itemDataHandler.GetAllItem(w, r)
	}).Methods("GET")

	groupItemData.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method == http.MethodPost {
			// Parse the ID parameter from the URL path
			vars := mux.Vars(r)
			id := vars["id"]

			// Handle the POST request with the provided ID
			// You can access the ID and other data from the request body
			// For example:
			// requestData := r.FormValue("someField")

			// Send a response indicating that the request was successful
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "POST request with ID %s was successful", id)
			return
		}

		// Respond with Method Not Allowed (405) for other HTTP methods
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}).Methods("POST")

}
