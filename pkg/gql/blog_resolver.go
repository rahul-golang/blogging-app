package gql

import (
	"blogging-app/log"
	"blogging-app/pkg/models"
	"blogging-app/pkg/service"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/graphql-go/graphql"
)

//BlogResolver for all BlogResolver function with blogService Leyer dependancies
type BlogResolver struct {
	blogService service.BlogService
}

//NewBlogResolver inject blogService dependancies
//helps to call  all the app Services
func NewBlogResolver(blogService service.BlogService) *BlogResolver {
	return &BlogResolver{blogService: blogService}
}

//CreateBlog blogResolver function call Service layer create function
func (blogResolver BlogResolver) CreateBlog(params graphql.ResolveParams) (interface{}, error) {
	var blog models.Blog
	var err error

	//get context for reading request specific attributes
	var ctx = params.Context

	//get userId from resolver Params
	strID := params.Args["user_id"].(string)

	//String to hex conversion
	blog.UserID, err = primitive.ObjectIDFromHex(strID)
	if err != nil {
		log.Logger(ctx).Error("Error in userid type conversion String to Hex : ", err)
		return nil, err
	}

	blog.Tittle = params.Args["tittle"].(string)
	blog.RelatedTo = params.Args["related_to"].(string)
	blog.Containt = params.Args["containt"].(string)
	//blog.Likes = params.Args["likes"]

	//call to blogservice method
	blogResp, err := blogResolver.blogService.CreateBlog(params.Context, blog)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}

	return blogResp, nil
}

//AllBlogsResolver return all user BlogResolver
func (blogResolver BlogResolver) AllBlogsResolver(param graphql.ResolveParams) (interface{}, error) {

	//get request context for request specific values
	ctx := param.Context
	//call blogservice methods
	blogs, err := blogResolver.blogService.GetAllBlogs(ctx)
	if err != nil {
		log.Logger(ctx).Error(err)
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
		panic("error in blog schema creation")
	}
	return schema

}
