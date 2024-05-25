package routes

import (
	"backend/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouteItemData(router *mux.Router) {
	itemDataHandler := handler.NewItemDataHandler()

	groupItemData := router.PathPrefix("/itemdatas").Subrouter()

	groupItemData.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		itemDataHandler.ItemDatasController(r)
	}).Methods("GET")

}
