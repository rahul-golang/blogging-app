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
		log.Logger(ctx).Error("Error in blogId type conversion String to Hex : ", err)
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

//UpdateBlog update blog
func (blogResolver BlogResolver) UpdateBlog(params graphql.ResolveParams) (interface{}, error) {
	var blog models.Blog
	var err error

	//get context for reading request specific attributes
	var ctx = params.Context

	//get userId from resolver Params
	strID := params.Args["id"].(string)

	//String to hex conversion
	blog.ID, err = primitive.ObjectIDFromHex(strID)
	if err != nil {
		log.Logger(ctx).Error("Error in blogId type conversion String to Hex : ", err)
		return nil, err
	}

	blog.Tittle = params.Args["tittle"].(string)
	blog.RelatedTo = params.Args["related_to"].(string)
	blog.Containt = params.Args["containt"].(string)
	blog.Likes = params.Args["likes"].(int)

	//call to blogservice method
	blogResp, err := blogResolver.blogService.UpdateBlog(params.Context, blog)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return blogResp, nil
}

//DeleteBlog Resolver function
func (blogResolver BlogResolver) DeleteBlog(params graphql.ResolveParams) (interface{}, error) {

	//get context for reading request specific attributes
	var ctx = params.Context

	//get userId from resolver Params
	strID := params.Args["id"].(string)

	result, err := blogResolver.blogService.DeleteBlog(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return result, nil
}

//GetBlog resolver function
func (blogResolver BlogResolver) GetBlog(params graphql.ResolveParams) (interface{}, error) {

	//get context for reading request specific attributes
	var ctx = params.Context

	//get userId from resolver Params
	strID := params.Args["id"].(string)

	result, err := blogResolver.blogService.GetBlog(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return result, nil
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
			"updateblog": &graphql.Field{
				Type: blogType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"tittle": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"related_to": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"containt": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"likes": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 0,
					},
				},
				Resolve: blogResolver.UpdateBlog,
			},
			"deleteblog": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: blogResolver.DeleteBlog,
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
				"blog": &graphql.Field{
					Type: blogType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: blogResolver.GetBlog,
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
