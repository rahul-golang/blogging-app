package handler

import (
	"blogging-app/pkg/gql"
	"blogging-app/pkg/service"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

//BlogHandlersImpl for handler Functions
type BlogHandlersImpl struct {
	blogSchema *gql.BlogSchema
}

//Blogs handler function
func (blogHandlersImpl BlogHandlersImpl) Blogs(w http.ResponseWriter, req *http.Request) {
	result := graphql.Do(graphql.Params{
		Schema:        blogHandlersImpl.blogSchema.BlogSchema,
		RequestString: req.URL.Query().Get("query"),
		Context:       req.Context(),
	})
	json.NewEncoder(w).Encode(result)
}

func writeResponse(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}

//NewBlogHandlersImpl inits dependancies for graphQL and Handlers
func NewBlogHandlersImpl(blogService service.BlogService) *BlogHandlersImpl {

	blogSchema := gql.NewBlogSchema(blogService)
	return &BlogHandlersImpl{blogSchema: blogSchema}

}
