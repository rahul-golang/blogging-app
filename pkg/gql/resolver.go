package gql

import (
	"blogging-app/pkg/models"
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	userService service.UsersService
}

func NewResolver(userService service.UsersService) *Resolver {
	fmt.Println("IN NewResolver")
	return &Resolver{userService: userService}
}
func (resolver Resolver) AllUserResolver(param graphql.ResolveParams) (interface{}, error) {
	fmt.Println("IN AllUserResolver")
	fmt.Println(resolver.userService)
	users, err := resolver.userService.GetAllUser(param.Context)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

func (resolver Resolver) CreateUser(params graphql.ResolveParams) (interface{}, error) {
	var user models.User

	//
	//user.ID = uint(params.Args["id"].(int))
	user.FirstName = params.Args["first_name"].(string)
	user.LastName = params.Args["last_name"].(string)
	user.Email = params.Args["email"].(string)
	user.Phone = params.Args["phone"].(string)
	user.Username = params.Args["username"].(string)
	user.Password = params.Args["password"].(string)
	var createUserReq models.CreateUserReq
	createUserReq.User = user
	users, err := resolver.userService.CreateUser(params.Context, createUserReq)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}
func (resolver *Resolver) NewSchema() graphql.Schema {

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createuser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{

					// "id": &graphql.ArgumentConfig{
					// 	Type: graphql.NewNonNull(graphql.Int),
					// },
					"first_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"last_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					// "created_at": &graphql.ArgumentConfig{
					// 	Type: graphql.NewNonNull(graphql.String),
					// },
					// "deleted_at": &graphql.ArgumentConfig{
					// 	Type: graphql.NewNonNull(graphql.String),
					// },
					// "updated_at": &graphql.ArgumentConfig{
					// 	Type: graphql.NewNonNull(graphql.String),
					// },
				},
				Resolve: resolver.CreateUser,
			},
		},
	})

	var userRoot = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserRoot",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Type:    graphql.NewList(userType),
					Resolve: resolver.AllUserResolver,

					// Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					// 	// id := params.Args["id"].(string)
					// 	// filtered := Filter(users, func(v models.User) bool {

					// 	// 	return id == id
					// 	// })
					// 	return users, nil
					// },
				},
			},
		},
	)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    userRoot,
		Mutation: rootMutation,
	})
	if err != nil {
		fmt.Println("return :", err)
		panic("error")
	}
	return schema

}
