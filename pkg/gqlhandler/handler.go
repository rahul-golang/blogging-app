package gqlhandler

import (
	"blogging-app/pkg/gql"
	"blogging-app/pkg/service"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

//GraphQlHandlers for handler Functions
type GraphQlHandlers struct {
	gqlAppSchema *gql.AppSchema
}

// Users handler Function
func (graphQlHandlers GraphQlHandlers) Users(w http.ResponseWriter, req *http.Request) {
	result := graphql.Do(graphql.Params{
		Schema:        graphQlHandlers.gqlAppSchema.GqlSchema,
		RequestString: req.URL.Query().Get("query"),
	})
	json.NewEncoder(w).Encode(result)
}

//Blogs handler function
func (graphQlHandlers GraphQlHandlers) Blogs(w http.ResponseWriter, req *http.Request) {
	result := graphql.Do(graphql.Params{
		Schema:        graphQlHandlers.gqlAppSchema.GqlSchema,
		RequestString: req.URL.Query().Get("query"),
	})
	json.NewEncoder(w).Encode(result)
}

func writeResponse(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}

//NewGraphQlHandlers inits dependancies for graphQL and Handlers
func NewGraphQlHandlers(appService service.AppService) *GraphQlHandlers {

	gqlAppSchema := gql.NewGqlAppSchema(appService)
	return &GraphQlHandlers{gqlAppSchema: gqlAppSchema}

}
