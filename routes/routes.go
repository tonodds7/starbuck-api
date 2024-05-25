package routes

import (
	"backend/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouteUserDate(router *mux.Router) {
	userDataHandler := handler.NewUseDataHandler()

	groupUserData := router.PathPrefix("/userdatas").Subrouter()

	groupUserData.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		userDataHandler.UserDatasController(r)
	}).Methods("GET")

}
