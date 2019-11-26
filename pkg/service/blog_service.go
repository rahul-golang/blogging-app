package service

import (
	"context"
	"errors"

	"blogging-app/log"
	"blogging-app/pkg/models"
	"blogging-app/pkg/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogService describes the blog services service.
type BlogService interface {

	// Blog Services Functions

	CreateBlog(context.Context, models.Blog) (interface{}, error)
	GetAllBlogs(context.Context) ([]models.Blog, error)
	UpdateBlog(context.Context, models.Blog) (interface{}, error)
	DeleteBlog(context.Context, string) (interface{}, error)
	GetBlog(context.Context, string) (interface{}, error)
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

//UpdateBlog create filter to update documnet and sent to repository
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

//DeleteBlog create delete filter and pass it to repository
func (b *BlogServiceImpl) DeleteBlog(ctx context.Context, strID string) (interface{}, error) {
	log.Logger(ctx).Info("Delete Blog Request ID : ", strID)

	//String to hex conversion
	id, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		log.Logger(ctx).Error("Error in blogId type conversion String to Hex : ", err)
		return nil, err
	}

	//mongo filter to Delete  blog
	filter := bson.M{"_id": bson.M{"$eq": id}}

	result, err := b.blogRepository.DeleteBlog(ctx, filter)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return result, nil

}

//GetBlog creates filter and pass it to repository
func (b *BlogServiceImpl) GetBlog(ctx context.Context, strID string) (interface{}, error) {
	log.Logger(ctx).Info("Get Blog Request ID : ", strID)

	//String to hex conversion
	id, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		log.Logger(ctx).Error("Error in blogId type conversion String to Hex : ", err)
		return nil, err
	}

	//mongo filter to Delete  blog
	filter := bson.M{"_id": bson.M{"$eq": id}}

	result, err := b.blogRepository.FindBlogs(ctx, filter)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err

	}

	if len(result) <= 0 {
		log.Logger(ctx).Error("No Blog Found")
		return nil, errors.New("No Blog Found")
	}

	return result[0], nil
}
