package service

import (
	"context"

	"blogging-app/log"
	"blogging-app/pkg/models"
	"blogging-app/pkg/repository"
)

// BlogService describes the blog services service.
type BlogService interface {

	// Blog Services Functions

	CreateBlog(context.Context, models.Blog) (interface{}, error)
	GetAllBlogs(context.Context) ([]*models.Blog, error)
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
func (b *BlogServiceImpl) GetAllBlogs(ctx context.Context) ([]*models.Blog, error) {
	allRecordResp, err := b.blogRepository.GetAllBlogs(ctx)
	if err != nil {
		return nil, err
	}

	return allRecordResp, err

}
