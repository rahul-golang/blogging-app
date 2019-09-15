package gqlhandler

import (
	"github.com/graphql-go/graphql"
)

type Root struct {
	Query *graphql.Object
}

func (graphQlHandlers GraphQlHandlers) UserResolver(params graphql.ResolveParams) (interface{}, error) {

	//name, ok := params.Args["name"].(string)
	//if ok {
	users, err := graphQlHandlers.userService.GetAllUser(params.Context)
	if err != nil {

		return nil, nil
	}
	return users, nil
	//}

}
