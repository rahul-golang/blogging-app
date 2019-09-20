package gql

import (
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

//AppSchema having GraphQL Schema dependancies
type AppSchema struct {
	GqlSchema graphql.Schema
}

var blogType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "blog",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"tittle": &graphql.Field{
				Type: graphql.String,
			},
			"related_to": &graphql.Field{
				Type: graphql.String,
			},
			"containt": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.Int,
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
		},
	},
)

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
		},
	},
)

//NewGqlAppSchema inialize GraphQL Schema
//Send Service dpendancies for Services to GraphQL Resolver functions
func NewGqlAppSchema(appService service.AppService) *AppSchema {
	resolver := NewResolver(appService)
	fmt.Println("IN NewGqlAppSchema")
	return &AppSchema{GqlSchema: resolver.NewSchema()}
}
