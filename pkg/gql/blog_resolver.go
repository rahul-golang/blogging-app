package gql

import (
	"blogging-app/pkg/models"
	"blogging-app/pkg/service"
	"fmt"

	"github.com/graphql-go/graphql"
)

//BlogResolver for all BlogResolver function with blogService Leyer dependancies
type BlogResolver struct {
	blogService service.BlogService
}

//NewBlogResolver inject blogService dependancies
//helps to call  all the app Services
func NewBlogResolver(blogService service.BlogService) *BlogResolver {
	fmt.Println("IN NewResolver")
	return &BlogResolver{blogService: blogService}
}

//CreateBlog blogResolver function call Service layer create function
func (blogResolver BlogResolver) CreateBlog(params graphql.ResolveParams) (interface{}, error) {
	var blog models.Blog

	//
	blog.UserID = params.Args["user_id"].(int)
	blog.Tittle = params.Args["tittle"].(string)
	blog.RelatedTo = params.Args["related_to"].(string)
	blog.Containt = params.Args["containt"].(string)

	blogResp, err := blogResolver.blogService.CreateBlog(params.Context, blog)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(blogResp)
	return blogResp, nil
}

//AllBlogsResolver return all user BlogResolver
func (blogResolver BlogResolver) AllBlogsResolver(param graphql.ResolveParams) (interface{}, error) {
	//fmt.Println("IN AllUserResolver")
	//fmt.Println(blogResolver.blogService)
	blogs, err := blogResolver.blogService.GetAllBlogs(param.Context)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return blogs, nil
}

//NewBlogSchemaImpl creates GraphQL Schema for Application
//excecutes only once when application starts
func (blogResolver *BlogResolver) NewBlogSchemaImpl() graphql.Schema {

	rootBlogMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{

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
				Resolve: blogResolver.CreateBlog,
			},
		},
	})

	var blogRoot = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "BlogRoot",
			Fields: graphql.Fields{

				"blogs": &graphql.Field{
					Type:    graphql.NewList(blogType),
					Resolve: blogResolver.AllBlogsResolver,
				},
			},
		},
	)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    blogRoot,
		Mutation: rootBlogMutation,
	})
	if err != nil {
		fmt.Println("return :", err)
		panic("error")
	}
	return schema

}
