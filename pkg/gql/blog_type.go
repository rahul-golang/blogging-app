package gql

import (
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

//BlogSchema having GraphQL Schema dependancies
type BlogSchema struct {
	BlogSchema graphql.Schema
}

//NewBlogSchema inialize GraphQL Schema
//Send Service dpendancies for Services to GraphQL Resolver functions
func NewBlogSchema(blogService service.BlogService) *BlogSchema {
	resolver := NewBlogResolver(blogService)
	fmt.Println("IN NewGqlAppSchema")
	return &BlogSchema{BlogSchema: resolver.NewBlogSchemaImpl()}
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
