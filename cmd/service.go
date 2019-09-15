package cmd

import (
	"flag"

	"net/http"

	"blogging-app/database"
	"blogging-app/log"
	"blogging-app/pkg/endpoints"
	"blogging-app/pkg/gqlhandler"
	"blogging-app/pkg/repository"
	"blogging-app/pkg/service"

	"github.com/gorilla/mux"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("users", flag.ExitOnError)
var debugAddr = fs.String("debug.addr", ":8080", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")

func Run() {

	router := mux.NewRouter()

	//Injected dependancies
	dataStoreInterface := database.NewMySQLClientConn()
	userRepositoryInterface := repository.NewUserRepository(dataStoreInterface)
	usersServiceInterface := service.NewBasicUsersService(userRepositoryInterface)
	//restHandler := httphandler.NewUserRestHandler(usersServiceInterface)
	restHandler := gqlhandler.NewGraphQlHandlers(usersServiceInterface)
	router.Use(loggingMiddleware)
	endpoints.NewAppRoute(router, restHandler)

	http.ListenAndServe(":8081", router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Do stuff here
		//	log.Println(r.RequestURI)
		req = req.WithContext(log.WithRqID(req.Context()))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, req)
	})
}
