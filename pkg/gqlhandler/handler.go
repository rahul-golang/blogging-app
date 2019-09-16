package gqlhandler

import (
	"blogging-app/pkg/gql"
	"blogging-app/pkg/service"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

type GraphQlHandlers struct {
	gqlAppSchema *gql.GqlAppSchema
}

func (graphQlHandlers GraphQlHandlers) GetAllUser(w http.ResponseWriter, req *http.Request) {

	//ctx := req.Context()
	//getAllUserResp := &models.GetAllUserResp{}

	//users, err := graphQlHandlers.userService.GetAllUser(ctx)
	// if err != nil {
	// 	getAllUserResp.Message = err.Error()
	// 	getAllUserResp.User = users
	// 	json.NewEncoder(w).Encode(getAllUserResp)
	// 	writeResponse(w, http.StatusInternalServerError)
	// 	return
	// }
	//getAllUserResp.User = users
	//log.Logger(ctx).Info("sucess")
	// err = json.NewEncoder(w).Encode(getAllUserResp)
	// if err != nil {
	// 	getAllUserResp.Message = err.Error()
	// 	getAllUserResp.User = users
	// 	json.NewEncoder(w).Encode(getAllUserResp)
	// 	writeResponse(w, http.StatusInternalServerError)
	// 	return
	// }
	//writeResponse(w, http.StatusOK)

	result := graphql.Do(graphql.Params{
		Schema:        graphQlHandlers.gqlAppSchema.GqlSchema,
		RequestString: req.URL.Query().Get("query"),
	})
	json.NewEncoder(w).Encode(result)
}

func writeResponse(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}

func NewGraphQlHandlers(userService service.UsersService) *GraphQlHandlers {

	gqlAppSchema := gql.NewGqlAppSchema(userService)
	return &GraphQlHandlers{gqlAppSchema: gqlAppSchema}

}
