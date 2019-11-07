package repository

import (
	"context"
	"fmt"
	"time"

	"blogging-app/database"
	"blogging-app/pkg/models"
)

//RepositoryInterface implimets all methods in AppRepository
type RepositoryInterface interface {
	CreateUser(ctx context.Context, createReq models.User) (createResp *models.User, err error)
	GetUser(ctx context.Context, id string) (getResp *models.User, err error)
	DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error)
	UpdateUser(ctx context.Context, updateReq models.User) (updateResp *models.User, err error)
	AllUsers(ctx context.Context) (getAllResp []*models.User, err error)
	GetUserProfile(ctx context.Context, id string) (*models.UserProfile, error)

	//Blog Repository FUnctions
	CreateBlog(ctx context.Context, blogReq models.Blog) (*models.Blog, error)
	GetAllBlogs(ctx context.Context) ([]*models.Blog, error)
}

// AppRepository **
type AppRepository struct {
	mongoConn database.MongoDBConnInterface
}

//NewAppRepository inject dependancies of DataStore
func NewAppRepository(mongoConn database.MongoDBConnInterface) RepositoryInterface {
	return &AppRepository{mongoConn: mongoConn}
}

//CreateUser add new record in datastore
func (appRepository *AppRepository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {

	client := appRepository.mongoConn.NewMongoConn()
	collection := client.Database("bloggingapp").Collection("users")

	//var err error
	fmt.Println(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	id := res.InsertedID
	user.ID = id.(int)
	return &user, nil
}

//GetUser serch user bu its id retun user fron database
func (appRepository *AppRepository) GetUser(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}

// DeleteUser delete User From Database
func (appRepository *AppRepository) DeleteUser(ctx context.Context, id string) (*models.DeleteUserResp, error) {
	return nil, nil
}

//UpdateUser update user in database
func (appRepository *AppRepository) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

//AllUsers return all Users from database
func (appRepository *AppRepository) AllUsers(ctx context.Context) (getAllResp []*models.User, err error) {
	return nil, nil
}

//GetUserProfile returns user profile data
func (appRepository *AppRepository) GetUserProfile(ctx context.Context, id string) (*models.UserProfile, error) {
	return nil, nil
}

// CreateBlog create blog in database and retub Created Blog
func (appRepository *AppRepository) CreateBlog(ctx context.Context, blogReq models.Blog) (*models.Blog, error) {
	return nil, nil
}

//GetAllBlogs return all Users from database
func (appRepository *AppRepository) GetAllBlogs(ctx context.Context) ([]*models.Blog, error) {
	return nil, nil
}
