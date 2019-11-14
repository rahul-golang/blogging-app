package gql

import "github.com/graphql-go/graphql"

var followerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "followers",
		Fields: graphql.Fields{
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
			"follower_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
