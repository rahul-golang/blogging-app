package endpoints

import (
	"blogging-app/pkg/gqlhandler"

	"github.com/gorilla/mux"
)

//NewAppRoute All Application Routes Are defiend Here
func NewAppRoute(router *mux.Router, gqlHandler *gqlhandler.GraphQlHandlers) {
	router.HandleFunc("/users", gqlHandler.Users).Methods("GET")
	router.HandleFunc("/blogs", gqlHandler.Blogs).Methods("GET")

}
