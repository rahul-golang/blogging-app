package repository

import (
	"blogging-app/log"
	"context"
	"time"

	"blogging-app/database"
	"blogging-app/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//BlogRepository implimets all methods in BlogRepositoryImpl
type BlogRepository interface {
	CreateBlog(context.Context, models.Blog) (interface{}, error)
	FindBlogs(context.Context, bson.M) ([]models.Blog, error)
	UpdateBlog(context.Context, bson.M, models.Blog) (interface{}, error)
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
func (blogRepositoryImpl *BlogRepositoryImpl) CreateBlog(ctx context.Context,
	blog models.Blog) (interface{}, error) {

	//mongo client connection
	mongoCon := blogRepositoryImpl.mongoConn.NewMongoConn(ctx)
	defer mongoCon.Disconnect(ctx)

	//Update Times Feilds When Created and Updated
	timeNow := time.Now()
	blog.CreatedAt = timeNow
	blog.UpdatedAt = timeNow

	//database and collection
	collection := mongoCon.Database("bloggingapp").Collection("blog")

	//insert opration on mongo collection
	result, err := collection.InsertOne(ctx, blog)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return result.InsertedID, nil
}

//FindBlogs return all Users from database
func (blogRepositoryImpl *BlogRepositoryImpl) FindBlogs(ctx context.Context,
	filter bson.M) ([]models.Blog, error) {

	var blogs []models.Blog

	//pass these options to find Method
	findOption := options.Find()

	//mongo connection
	mongoConn := blogRepositoryImpl.mongoConn.NewMongoConn(ctx)
	defer mongoConn.Disconnect(ctx)

	//collection
	collection := mongoConn.Database("bloggingapp").Collection("blog")
	cur, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}

	//finding out multiple documents return a cursur
	//itterating throught cursor allows to decode one dcumnets at a time
	for cur.Next(ctx) {
		blog := models.Blog{}
		err := cur.Decode(&blog)
		if err != nil {
			log.Logger(ctx).Error(err)
			return nil, err
		}
		//append all blog to blogs
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

//UpdateBlog upadets presented values into blog
func (blogRepositoryImpl *BlogRepositoryImpl) UpdateBlog(ctx context.Context, filter bson.M,
	blog models.Blog) (interface{}, error) {

	//mongo client connection
	mongoCon := blogRepositoryImpl.mongoConn.NewMongoConn(ctx)
	defer mongoCon.Disconnect(ctx)

	//database and collection
	collection := mongoCon.Database("bloggingapp").Collection("blog")

	//update documnet
	updateDoc := bson.M{}

	//Final Update Object
	data := bson.M{}

	//like are incrimented by one
	if blog.Likes != 0 {
		data["$inc"] = bson.M{"likes": 1}
	}
	//set to  new tittle if any
	if len(blog.Tittle) > 0 {
		updateDoc["tittle"] = blog.Tittle
	}
	//set to  new value if any
	if len(blog.RelatedTo) > 0 {
		updateDoc["related_to"] = blog.RelatedTo
	}

	//set to  new value if any
	if len(blog.Containt) > 0 {
		updateDoc["containt"] = blog.Containt
	}

	//add updated time of documnets
	updateDoc["updated_at"] = time.Now()

	//final data to update
	data["$set"] = updateDoc

	//insert opration on mongo collection
	result, err := collection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return result, nil
}
