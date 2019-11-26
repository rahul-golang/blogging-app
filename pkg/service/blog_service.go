package service

import (
	"context"

	"blogging-app/log"
	"blogging-app/pkg/models"
	"blogging-app/pkg/repository"

	"go.mongodb.org/mongo-driver/bson"
)

// BlogService describes the blog services service.
type BlogService interface {

	// Blog Services Functions

	CreateBlog(context.Context, models.Blog) (interface{}, error)
	GetAllBlogs(context.Context) ([]models.Blog, error)
	UpdateBlog(context.Context, models.Blog) (interface{}, error)
}

//BlogServiceImpl implemts all the BlogService
type BlogServiceImpl struct {
	blogRepository repository.BlogRepository
}

// NewBlogServiceImpl returns a naive, stateless implementation of AppService.
func NewBlogServiceImpl(blogRepository repository.BlogRepository) BlogService {
	return &BlogServiceImpl{blogRepository: blogRepository}
}

//CreateBlog create blog and return id
func (b *BlogServiceImpl) CreateBlog(ctx context.Context, blog models.Blog) (interface{}, error) {

	log.Logger(ctx).Info("Create Blog Request : ", blog)
	resp, err := b.blogRepository.CreateBlog(ctx, blog)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}

	return resp, err
}

//GetAllBlogs retun slice of blogs
func (b *BlogServiceImpl) GetAllBlogs(ctx context.Context) ([]models.Blog, error) {

	log.Logger(ctx).Info("Get ALL Blogs Request")

	//TODO:  Add limit and Offset
	//mongo filter to find all blogs
	filter := bson.M{}
	resp, err := b.blogRepository.FindBlogs(ctx, filter)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return resp, err
}

func (b *BlogServiceImpl) UpdateBlog(ctx context.Context, blog models.Blog) (interface{}, error) {

	log.Logger(ctx).Info("Get ALL Blogs Request")

	//TODO:  Add limit and Offset
	//mongo filter to find all blogs
	filter := bson.M{"_id": bson.M{"$eq": blog.ID}}
	resp, err := b.blogRepository.UpdateBlog(ctx, filter, blog)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return resp, err
}
