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

		vars := mux.Vars(r)
		id := vars["id"]

		itemDataHandler.GetItemDetailById(w, r, id)
	}).Methods("GET")

	groupItemData.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		itemDataHandler.OrderItem(w, r)
	}).Methods("PATCH")

}
