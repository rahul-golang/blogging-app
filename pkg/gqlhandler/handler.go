package gqlhandler

import (
	"blogging-app/pkg/models"
	"blogging-app/pkg/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type GraphQlHandlers struct {
	userService service.UsersService
	GqlSchema   graphql.Schema
}

func UserHandler() {

}

func (graphQlHandlers GraphQlHandlers) GetAllUser(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	getAllUserResp := &models.GetAllUserResp{}

	users, err := graphQlHandlers.userService.GetAllUser(ctx)
	if err != nil {
		getAllUserResp.Message = err.Error()
		getAllUserResp.User = users
		json.NewEncoder(w).Encode(getAllUserResp)
		writeResponse(w, http.StatusInternalServerError)
		return
	}
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
	writeResponse(w, http.StatusOK)

	result := graphql.Do(graphql.Params{
		Schema:        graphQlHandlers.GqlSchema,
		RequestString: req.URL.Query().Get("query"),
	})
	json.NewEncoder(w).Encode(result)
}

func writeResponse(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}

func NewGraphQlHandlers(userService service.UsersService) *GraphQlHandlers {

	graphqlsssss := GraphQlHandlers{userService: userService}
	var userRoot = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserRoot",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Type:    graphql.NewList(userType),
					Resolve: graphqlsssss.UserResolver,
				},
			},
		},
	)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: userRoot,
	})
	if err != nil {
		fmt.Println("return :")
		panic("error")
	}
	return &graphqlsssss
}
