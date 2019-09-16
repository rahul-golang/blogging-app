package endpoints

import (
	"blogging-app/pkg/gqlhandler"

	"github.com/gorilla/mux"
)

func NewAppRoute(router *mux.Router, gqlHandler *gqlhandler.GraphQlHandlers) {
	router.HandleFunc("/users", gqlHandler.GetAllUser).Methods("GET")

}
