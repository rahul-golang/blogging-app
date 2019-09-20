package gql

import (
	"blogging-app/pkg/models"
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

//Resolver for all Resolver function with AppService Leyer dependancies
type Resolver struct {
	appService service.AppService
}

//NewResolver inject appservice dependancies
//helps to call  all the app Services
func NewResolver(appService service.AppService) *Resolver {
	fmt.Println("IN NewResolver")
	return &Resolver{appService: appService}
}

//AllUserResolver return all user Resolver
func (resolver Resolver) AllUserResolver(param graphql.ResolveParams) (interface{}, error) {
	fmt.Println("IN AllUserResolver")
	fmt.Println(resolver.appService)
	users, err := resolver.appService.GetAllUser(param.Context)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

//CreateUser resolver function to create user Record
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
	users, err := resolver.appService.CreateUser(params.Context, createUserReq)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}

//CreateBlog resolver function call Service layer create function
func (resolver Resolver) CreateBlog(params graphql.ResolveParams) (interface{}, error) {
	var blog models.Blog

	//
	blog.UserID = params.Args["user_id"].(int)
	blog.Tittle = params.Args["tittle"].(string)
	blog.RelatedTo = params.Args["related_to"].(string)
	blog.Containt = params.Args["containt"].(string)

	blogResp, err := resolver.appService.CreateBlog(params.Context, blog)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(blogResp)
	return blogResp, nil
}

//AllBlogsResolver return all user Resolver
func (resolver Resolver) AllBlogsResolver(param graphql.ResolveParams) (interface{}, error) {
	fmt.Println("IN AllUserResolver")
	fmt.Println(resolver.appService)
	blogs, err := resolver.appService.GetAllBlogs(param.Context)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return blogs, nil
}

//NewSchema creates GraphQL Schema for Application
//excecutes only once when application starts
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
			"createblog": &graphql.Field{
				Type: blogType,
				Args: graphql.FieldConfigArgument{

					// "id": &graphql.ArgumentConfig{
					// 	Type: graphql.NewNonNull(graphql.Int),
					// },
					"tittle": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"related_to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"containt": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
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
				Resolve: resolver.CreateBlog,
			},
		},
	})

	var userRoot = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "AppRoot",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Type:    graphql.NewList(userType),
					Resolve: resolver.AllUserResolver,
				},
				"blogs": &graphql.Field{
					Type:    graphql.NewList(blogType),
					Resolve: resolver.AllBlogsResolver,
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
