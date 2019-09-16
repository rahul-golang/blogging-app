package gql

import (
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

type GqlAppSchema struct {
	GqlSchema graphql.Schema
	//userService service.UsersService
}

// var users []models.User = []models.User{
// 	models.User{
// 		//	ID:        1,
// 		Email:     "rahul.shinde@scalent.io",
// 		FirstName: "Rahul",
// 		LastName:  "shinde",
// 		Phone:     "9975227706",
// 	},
// 	models.User{
// 		//gorm.Model.ID: 2,
// 		Email:     "sagar.pawar@scalent.io",
// 		FirstName: "Sagar",
// 		LastName:  "pawar",
// 		Phone:     "7218969895",
// 	},
// }

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

func NewGqlAppSchema(userService service.UsersService) *GqlAppSchema {
	resolver := NewResolver(userService)
	fmt.Println("IN NewGqlAppSchema")
	return &GqlAppSchema{GqlSchema: resolver.NewSchema()}
}
