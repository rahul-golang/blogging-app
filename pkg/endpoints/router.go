package endpoints

import (
	"blogging-app/pkg/handler"

	"github.com/gorilla/mux"
)

//NewUserRoute All Application Routes Are defiend Here
func NewUserRoute(router *mux.Router, handler *handler.UserHandlersImpl) {
	router.HandleFunc("/users", handler.Users).Methods("GET")

}

//NewBlogRoute All Application Routes Are defiend Here
func NewBlogRoute(router *mux.Router, handler *handler.BlogHandlersImpl) {
	router.HandleFunc("/blogs", handler.Blogs).Methods("GET")

}
