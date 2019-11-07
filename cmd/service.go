package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"net/http"

	"blogging-app/database"
	"blogging-app/log"
	"blogging-app/pkg/endpoints"
	"blogging-app/pkg/gqlhandler"
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

	httpAddressEnvVar = "HTTP-ADDR"

	MongoDBHostEnvVar = "MONGO-DB-HOST"
	MongoDBUserEnvVar = "MONGO-DB-USER"
	MongoDBPassEnvVar = "MONGO-DB-PASS"
	MongoDBAddrEnvVar = "MONGO-DB-ADDR"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("Blogging-App", flag.ExitOnError)

//HttpAddr http Address
var HttpAddr = fs.String(HTTPAddressFlag, "8080", "HTTP listen address defaults to 8080")

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
	var httpAddr = os.Getenv(httpAddressEnvVar)
	if len(httpAddr) > 0 && (HttpAddr == nil || len(*HttpAddr) == 0) {
		HttpAddr = &httpAddr
	}
}

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

func Run() {

	router := mux.NewRouter()

	//Injected dependancies
	//dataStoreInterface := database.NewMySQLClientConn()
	dataStoreInterface := database.NewMongoDBConn()
	appRepositoryInterface := repository.NewAppRepository(dataStoreInterface)
	appServiceInterface := service.NewBasicAppService(appRepositoryInterface)
	restHandler := gqlhandler.NewGraphQlHandlers(appServiceInterface)
	router.Use(loggingMiddleware)
	endpoints.NewAppRoute(router, restHandler)
	fmt.Println(http.ListenAndServe(":"+httpAddr, router))
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
