package gql

import (
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

//UserSchema having GraphQL Schema dependancies
type UserSchema struct {
	UserSchema graphql.Schema
}

//NewUserSchema inialize GraphQL Schema
//Send Service dpendancies for Services to GraphQL Resolver functions
func NewUserSchema(userService service.UserService) *UserSchema {
	resolver := NewUserResolver(userService)
	fmt.Println("IN NewGqlAppSchema")
	return &UserSchema{UserSchema: resolver.NewUserSchemaImpl()}
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"first_name": &graphql.Field{
				Type: graphql.String,
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"deleted_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"blogs": &graphql.Field{
				Type: graphql.NewList(blogType),
			},
		},
	},
)

var userProfileType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "userprofile",
		Fields: graphql.Fields{
			// "id": &graphql.Field{
			// 	Type: graphql.Int,
			// },
			// "first_name": &graphql.Field{
			// 	Type: graphql.String,
			// },
			// "last_name": &graphql.Field{
			// 	Type: graphql.String,
			// },
			// "email": &graphql.Field{
			// 	Type: graphql.String,
			// },
			// "phone": &graphql.Field{
			// 	Type: graphql.String,
			// },
			// "username": &graphql.Field{
			// 	Type: graphql.String,
			// },
			// "password": &graphql.Field{
			// 	Type: graphql.String,
			// },
			// "created_at": &graphql.Field{
			// 	Type: graphql.DateTime,
			// },
			// "deleted_at": &graphql.Field{
			// 	Type: graphql.DateTime,
			// },
			// "updated_at": &graphql.Field{
			// 	Type: graphql.DateTime,
			// },
			"user": &graphql.Field{
				Type: userType,
			},
			"followers": &graphql.Field{
				Type: graphql.NewList(followerType),
			},
			"blogs": &graphql.Field{
				Type: graphql.NewList(blogType),
			},
		},
	},
)
