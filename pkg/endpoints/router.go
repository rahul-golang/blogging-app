package endpoints

import (
	"github.com/gorilla/mux"
	"github.com/rahul-golang/ecommerce/users/pkg/httphandler"
)

func NewAppRoute(router *mux.Router, restHandler *httphandler.UserHttpHandlers) {
	router.HandleFunc("/users", restHandler.GetUser).Methods("GET")

}
