package gql

import (
	"blogging-app/log"
	"blogging-app/pkg/models"
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ctx := param.Context
	users, err := resolver.userService.GetAllUser(ctx)

	if err != nil {
		log.Logger(ctx).Errorf("Error from Service : %v ", err)
		return nil, err
	}
	return users, nil
}

//CreateUser resolver function to create user Record
func (resolver UserResolver) CreateUser(params graphql.ResolveParams) (interface{}, error) {
	var user models.User
	user.FirstName = params.Args["first_name"].(string)
	user.LastName = params.Args["last_name"].(string)
	user.Email = params.Args["email"].(string)
	user.Phone = params.Args["phone"].(string)
	user.Username = params.Args["username"].(string)
	user.Password = params.Args["password"].(string)
	ctx := params.Context

	users, err := resolver.userService.CreateUser(ctx, user)
	if err != nil {
		log.Logger(ctx).Errorf("Error from Service : %v ", err)
		return nil, err
	}
	return users, nil
}

//UpdateUser resolver function to create user Record
func (resolver UserResolver) UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	//get context from request
	ctx := params.Context

	var user models.User
	id := params.Args["id"].(string)

	// string to primitive.ObjectID conversion
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Logger(ctx).Error("Error from Service : ", err)
		return nil, err
	}

	user.ID = objectID
	user.FirstName = params.Args["first_name"].(string)
	user.LastName = params.Args["last_name"].(string)
	user.Email = params.Args["email"].(string)
	user.Phone = params.Args["phone"].(string)
	user.Username = params.Args["username"].(string)
	user.Password = params.Args["password"].(string)

	users, err := resolver.userService.UpdateUser(ctx, user)
	if err != nil {
		log.Logger(ctx).Errorf("Error from Service : %v ", err)
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
				},
				Resolve: resolver.CreateUser,
			},
			"updateuser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{

					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"first_name": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"last_name": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"email": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"phone": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"username": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"password": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: resolver.UpdateUser,
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
