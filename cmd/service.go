package cmd

import (
	"blogging-app/pkg/handler"
	"errors"
	"flag"
	"fmt"
	"os"

	"net/http"

	"blogging-app/database"
	"blogging-app/log"
	"blogging-app/pkg/endpoints"
	"blogging-app/pkg/repository"
	"blogging-app/pkg/service"

	"github.com/gorilla/mux"
)

var (
	//HTTPAddressFlag cli flag name for http port
	HTTPAddressFlag = "http-add"

	//MongoDBHostFlag cli flag name for mongodb host
	MongoDBHostFlag = "mongodb-host"

	//MongoDBUserFlag cli flag name for mongodb username
	MongoDBUserFlag = "mongodb-user"

	//MongoDBPassFlag cli flag name for mongodb password
	MongoDBPassFlag = "mongodb-pass"

	//MongoDBAddrFlag cli flag name for mongodb port
	MongoDBAddrFlag = "mongodb-addr"

	//DebugAddressFlag cli flag name for debug address
	DebugAddressFlag = "debug-addr"

	//HTTPAddressEnvVar **
	HTTPAddressEnvVar = "HTTP-ADDR"

	//MongoDBHostEnvVar **
	MongoDBHostEnvVar = "MONGO-DB-HOST"

	//MongoDBUserEnvVar **
	MongoDBUserEnvVar = "MONGO-DB-USER"

	//MongoDBPassEnvVar **
	MongoDBPassEnvVar = "MONGO-DB-PASS"

	//MongoDBAddrEnvVar **
	MongoDBAddrEnvVar = "MONGO-DB-ADDR"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("Blogging-App", flag.ExitOnError)

//HTTPAddr http Address
var HTTPAddr = fs.String(HTTPAddressFlag, "8080", "HTTP listen address defaults to 8080")

//MongoDBHost mongodb  hostname
var MongoDBHost = fs.String(MongoDBHostFlag, "", "Hostname for mongoDB")

//MongoDBUser mongodb username
var MongoDBUser = fs.String(MongoDBUserFlag, "", "Username for mongoDB")

//MongoDBPass mongodb password
var MongoDBPass = fs.String(MongoDBPassFlag, "", "Password for mongoDB")

//MongoDBAddr mongodb port address
var MongoDBAddr = fs.String(MongoDBAddrFlag, "27017", "Port Number for mongoDB defaults to 27017")

func init() {
	flag.Parse()
}

//GetEnviromentVariables **
func GetEnviromentVariables() {

	//get mongoDBHost from enviroment variables
	var mongoDBHost = os.Getenv(MongoDBHostEnvVar)
	if len(mongoDBHost) > 0 && (MongoDBHost == nil || len(*MongoDBHost) == 0) {
		MongoDBHost = &mongoDBHost
	}

	//get mongoDBUser from enviroments variables
	var mongoDBUser = os.Getenv(MongoDBUserEnvVar)
	if len(mongoDBUser) > 0 && (MongoDBUser == nil || len(*MongoDBUser) == 0) {
		MongoDBUser = &mongoDBUser
	}

	//get mongoDBPass from enviroments variables
	var mongoDBPass = os.Getenv(MongoDBPassEnvVar)
	if len(mongoDBPass) > 0 && (MongoDBPass == nil || len(*MongoDBPass) == 0) {
		MongoDBPass = &mongoDBPass
	}

	//get mongoDBAddr from enviroments variables
	var mongoDBAddr = os.Getenv(MongoDBAddrEnvVar)
	if len(mongoDBAddr) > 0 && (MongoDBAddr == nil || len(*MongoDBAddr) == 0) {
		MongoDBAddr = &mongoDBAddr
	}

	//get httpAddr from enviroments variables
	var httpAddr = os.Getenv(HTTPAddressEnvVar)
	if len(httpAddr) > 0 && (HTTPAddr == nil || len(*HTTPAddr) == 0) {
		HTTPAddr = &httpAddr
	}
}

//ValidateFlags ckecks the flags and update
func ValidateFlags() error {
	GetEnviromentVariables()
	flagMessage := "is a requird flag"
	if MongoDBUser == nil || len(*MongoDBUser) == 0 {
		return errors.New(MongoDBUserFlag + flagMessage)
	}
	if MongoDBPass == nil || len(*MongoDBPass) == 0 {
		return errors.New(MongoDBPassFlag + flagMessage)
	}
	return nil
}

//Run **
func Run() {

	router := mux.NewRouter()

	//MongoDB Dependancy
	dataStore := database.NewMongoDBConn()

	//User Services dependancies
	userRepository := repository.NewUserRepositoryImpl(dataStore)
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := handler.NewUserHandlerImpl(userService)

	endpoints.NewUserRoute(router, userHandler)

	//Blog Services dependancies
	blogRepository := repository.NewBlogRepositoryImpl(dataStore)
	blogService := service.NewBlogServiceImpl(blogRepository)
	blogHandler := handler.NewBlogHandlersImpl(blogService)
	endpoints.NewBlogRoute(router, blogHandler)

	router.Use(loggingMiddleware)
	fmt.Println(http.ListenAndServe(":"+*HTTPAddr, router))
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
