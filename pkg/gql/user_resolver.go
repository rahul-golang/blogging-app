package gql

import (
	"blogging-app/pkg/models"
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

//UserResolver for all UserResolver function with UserService Leyer dependancies
type UserResolver struct {
	userService service.UserService
}

//NewUserResolver inject appservice dependancies
//helps to call  all the app Services
func NewUserResolver(userService service.UserService) *UserResolver {
	fmt.Println("IN NewResolver")
	return &UserResolver{userService: userService}
}

//AllUserResolver return all user UserResolver
func (resolver UserResolver) AllUserResolver(param graphql.ResolveParams) (interface{}, error) {
	fmt.Println("IN AllUserResolver")
	fmt.Println(resolver.userService)
	users, err := resolver.userService.GetAllUser(param.Context)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

//CreateUser resolver function to create user Record
func (resolver UserResolver) CreateUser(params graphql.ResolveParams) (interface{}, error) {
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
		return nil, err
	}
	return users, nil
}

//NewUserSchemaImpl creates GraphQL Schema for User Schema
//excecutes only once when application starts
func (resolver *UserResolver) NewUserSchemaImpl() graphql.Schema {

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
