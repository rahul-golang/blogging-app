package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"blogging-app/database"
	"blogging-app/log"
	"blogging-app/pkg/models"
)

//UserRepository implimets all methods in UserRepository
type UserRepository interface {
	CreateUser(ctx context.Context, createReq models.User) (createResp *models.User, err error)
	GetUser(ctx context.Context, id string) (getResp *models.User, err error)
	DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error)
	UpdateUser(ctx context.Context, updateReq models.User) (updateResp *models.User, err error)
	AllUsers(ctx context.Context) (getAllResp []*models.User, err error)
	GetUserProfile(ctx context.Context, id string) (*models.UserProfile, error)
}

//UserRepositoryImpl **
type UserRepositoryImpl struct {
	mongoConn database.MongoDBConnInterface
}

//NewUserRepositoryImpl inject dependancies of DataStore
func NewUserRepositoryImpl(mongoConn database.MongoDBConnInterface) UserRepository {
	return &UserRepositoryImpl{mongoConn: mongoConn}
}

//CreateUser add new record in datastore
func (userRepositoryImpl *UserRepositoryImpl) CreateUser(ctx context.Context, user models.User) (*models.User, error) {

	client := userRepositoryImpl.mongoConn.NewMongoConn(ctx)
	defer client.Disconnect(ctx)
	collection := client.Database("bloggingapp").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	log.Logger(ctx).Info(res)
	//assertion interface type to primitive.ObjectID
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}
	user.ID = id
	return &user, nil
}

//GetUser serch user bu its id retun user fron database
func (userRepositoryImpl *UserRepositoryImpl) GetUser(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}

// DeleteUser delete User From Database
func (userRepositoryImpl *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) (*models.DeleteUserResp, error) {
	return nil, nil
}

//UpdateUser update user in database
func (userRepositoryImpl *UserRepositoryImpl) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

//AllUsers return all Users from database
func (userRepositoryImpl *UserRepositoryImpl) AllUsers(ctx context.Context) (getAllResp []*models.User, err error) {
	return nil, nil
}

//GetUserProfile returns user profile data
func (userRepositoryImpl *UserRepositoryImpl) GetUserProfile(ctx context.Context, id string) (*models.UserProfile, error) {
	return nil, nil
}

// CreateBlog create blog in database and retub Created Blog
func (userRepositoryImpl *UserRepositoryImpl) CreateBlog(ctx context.Context, blogReq models.Blog) (*models.Blog, error) {
	return nil, nil
}

//GetAllBlogs return all Users from database
func (userRepositoryImpl *UserRepositoryImpl) GetAllBlogs(ctx context.Context) ([]*models.Blog, error) {
	return nil, nil
}
