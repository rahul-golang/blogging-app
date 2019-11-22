package handler

import (
	"blogging-app/pkg/gql"
	"blogging-app/pkg/service"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

//UserHandlersImpl for handler Functions
type UserHandlersImpl struct {
	userSchema *gql.UserSchema
}

// Users handler Function
func (userHandlersImpl UserHandlersImpl) Users(w http.ResponseWriter, req *http.Request) {
	result := graphql.Do(graphql.Params{
		Schema:        userHandlersImpl.userSchema.UserSchema,
		RequestString: req.URL.Query().Get("query"),
		Context:       req.Context(),
	})
	json.NewEncoder(w).Encode(result)
}

//NewUserHandlerImpl inits dependancies for graphQL and Handlers
func NewUserHandlerImpl(userService service.UserService) *UserHandlersImpl {

	userSchema := gql.NewUserSchema(userService)
	return &UserHandlersImpl{userSchema: userSchema}

}
