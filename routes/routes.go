package routes

import (
	"backend/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RouteItemData(router *mux.Router) {
	itemDataHandler := handler.NewItemDataHandler()

	groupItemData := router.PathPrefix("/itemdatas").Subrouter()

	groupItemData.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		itemDataHandler.GetAllItem(w, r)
	}).Methods("GET")

	groupItemData.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {

		// Extract the ID from the URL path
		vars := mux.Vars(r)
		id := vars["id"]

		itemDataHandler.GetItemDetailById(w, r, id)
	}).Methods("GET")

	groupItemData.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Handle the POST request
		itemDataHandler.OrderItem(w, r)
	}).Methods("POST")

}
