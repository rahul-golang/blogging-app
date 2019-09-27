package service

import (
	"context"
	"fmt"

	"blogging-app/pkg/models"
	"blogging-app/pkg/repository"
)

// AppService describes the service.
type AppService interface {
	// User Service Functions
	CreateUser(ctx context.Context, createReq models.CreateUserReq) (createResp *models.CreateUserResp, err error)
	GetAllUser(ctx context.Context) (allRecordResp []*models.User, err error)
	UpdateUser(ctx context.Context, upadteReq models.User) (updateResp *models.User, err error)
	DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error)
	GetUser(ctx context.Context, id string) (createResp *models.User, err error)

	// Blog Services Functions

	CreateBlog(ctx context.Context, blogReq models.Blog) (blogResp *models.Blog, err error)
	GetAllBlogs(ctx context.Context) ([]*models.Blog, error)

	GetUserProfile(ctx context.Context, id string) (*models.UserProfile, error)
}

type basicAppService struct {
	repositoryInterface repository.RepositoryInterface
}

func (b *basicAppService) CreateUser(ctx context.Context, createReq models.CreateUserReq) (*models.CreateUserResp, error) {
	user, err := b.repositoryInterface.CreateUser(ctx, createReq.User)
	if err != nil {
		return nil, err
	}

	return &models.CreateUserResp{
		Message: "record created sucessfully",
		User:    user,
	}, err
}
func (b *basicAppService) GetAllUser(ctx context.Context) (allRecordResp []*models.User, err error) {
	fmt.Println("in all users service mothod")
	//log.Logger(ctx).Info("in all users service mothod ")

	allRecordResp, err = b.repositoryInterface.AllUsers(ctx)
	fmt.Println(allRecordResp)
	return allRecordResp, err
}
func (b *basicAppService) UpdateUser(ctx context.Context, upadteReq models.User) (updateResp *models.User, err error) {
	updateResp, err = b.repositoryInterface.UpdateUser(ctx, upadteReq)
	return updateResp, err
}
func (b *basicAppService) DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error) {
	deleteResp, err = b.repositoryInterface.DeleteUser(ctx, id)
	return deleteResp, err
}
func (b *basicAppService) GetUser(ctx context.Context, id string) (createResp *models.User, err error) {
	fmt.Println("id", id)
	createResp, err = b.repositoryInterface.GetUser(ctx, id)
	return createResp, err
}

func (b *basicAppService) GetUserProfile(ctx context.Context, id string) (userProfile *models.UserProfile, err error) {
	fmt.Println("id", id)
	userProfile, err = b.repositoryInterface.GetUserProfile(ctx, id)
	return userProfile, err
}

func (b *basicAppService) CreateBlog(ctx context.Context, blogReq models.Blog) (blogResp *models.Blog, err error) {

	blog, err := b.repositoryInterface.CreateBlog(ctx, blogReq)
	if err != nil {
		return nil, err
	}

	return blog, err
}

func (b *basicAppService) GetAllBlogs(ctx context.Context) ([]*models.Blog, error) {
	allRecordResp, err := b.repositoryInterface.GetAllBlogs(ctx)
	if err != nil {
		return nil, err
	}

	return allRecordResp, err

}

// NewBasicAppService returns a naive, stateless implementation of AppService.
func NewBasicAppService(repositoryInterface repository.RepositoryInterface) AppService {
	return &basicAppService{repositoryInterface: repositoryInterface}
}
