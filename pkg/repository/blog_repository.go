package repository

import (
	"blogging-app/log"
	"context"
	"fmt"

	"blogging-app/database"
	"blogging-app/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//BlogRepository implimets all methods in BlogRepositoryImpl
type BlogRepository interface {
	CreateBlog(context.Context, models.Blog) (interface{}, error)
	GetAllBlogs(context.Context) ([]*models.Blog, error)
}

// BlogRepositoryImpl **
type BlogRepositoryImpl struct {
	mongoConn database.MongoDBConnInterface
}

//NewBlogRepositoryImpl inject dependancies of DataStore
func NewBlogRepositoryImpl(mongoConn database.MongoDBConnInterface) BlogRepository {
	return &BlogRepositoryImpl{mongoConn: mongoConn}
}

// CreateBlog create blog in database and retub Created Blog
func (blogRepositoryImpl *BlogRepositoryImpl) CreateBlog(ctx context.Context, blog models.Blog) (interface{}, error) {
	mongoCon := blogRepositoryImpl.mongoConn.NewMongoConn(ctx)
	result, err := mongoCon.Database("bloggingapp").Collection("blog").InsertOne(ctx, blog)
	if err != nil {
		//	log.Logger(ctx).Info("in all users service mothod ")
		fmt.Println(err)
		return nil, err
	}

	return result.InsertedID, nil
}

//GetAllBlogs return all Users from database
func (blogRepositoryImpl *BlogRepositoryImpl) GetAllBlogs(ctx context.Context) ([]*models.Blog, error) {
	blogs := []models.Blog{}

	//pass these options to find Method
	findOption := options.Find()

	cur, err := blogRepositoryImpl.mongoConn.NewMongoConn(ctx).Database("bloggingapp").Collection("blog").Find(ctx, bson.M{}, findOption)

	if err != nil {
		log.Logger(ctx).Info("Get All Users Methosd ")
	}

	//finding out multiple documents return a cursur
	//itterating throught cursor allows to decode one dcumnets at a time
	for cur.Next(ctx) {
		blog := models.Blog{}
		err := cur.Decode(&blog)
		if err != nil {
			log.Logger(ctx).Info("Get All Users Methosd ")
		}
		blogs = append(blogs, blog)
	}

	return nil, nil
}
