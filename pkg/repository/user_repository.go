package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"blogging-app/database"
	"blogging-app/log"
	"blogging-app/pkg/models"
)

//UserRepository implimets all methods in UserRepository
type UserRepository interface {
	CreateUser(context.Context, models.User) (*models.User, error)
	GetUsers(context.Context, bson.M) ([]models.User, error)
	DeleteUser(context.Context, bson.M) (interface{}, error)
	UpdateUser(context.Context, bson.M, models.User) (*models.User, error)
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

	//  mongo client connection
	client := userRepositoryImpl.mongoConn.NewMongoConn(ctx)
	defer client.Disconnect(ctx)

	//Update Times Feilds When Created and Updated
	timeNow := time.Now()
	user.CreatedAt = timeNow
	user.UpdatedAt = timeNow

	//  mongo client Collection and Db
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

//GetUsers serch user and returns listof users
func (userRepositoryImpl *UserRepositoryImpl) GetUsers(ctx context.Context, filter bson.M) ([]models.User, error) {
	var users []models.User

	findOptiond := options.Find()

	//  mongo client connection
	mongoClient := userRepositoryImpl.mongoConn.NewMongoConn(ctx)
	defer mongoClient.Disconnect(ctx)

	//db and collection
	collection := mongoClient.Database("bloggingapp").Collection("users")

	//fetch data from mongo database on tahe basis of filters
	cur, err := collection.Find(ctx, filter, findOptiond)
	if err != nil {
		log.Logger(ctx).Errorf("Error in finding mongo users : %v", err)
		return nil, err
	}

	//if cursor having more documents its will itterate over a single documnets
	for cur.Next(ctx) {
		var user models.User
		if err = cur.Decode(&user); err != nil {
			log.Logger(ctx).Errorf("Error in decoding cursor into user model  : %v ", err)
		}
		users = append(users, user)

	}

	if err := cur.Err(); err != nil {
		log.Logger(ctx).Errorf("Error in finding mongo users : %v", err)

		return nil, err
	}

	return users, nil
}

// DeleteUser delete User From Database
func (userRepositoryImpl *UserRepositoryImpl) DeleteUser(ctx context.Context, filter bson.M) (interface{}, error) {

	//specify delete optons if any
	deleteOptions := options.Delete()

	//  mongo client connection
	mongoClient := userRepositoryImpl.mongoConn.NewMongoConn(ctx)
	defer mongoClient.Disconnect(ctx)

	//db and collection
	collection := mongoClient.Database("bloggingapp").Collection("users")

	//fetch data from mongo database on tahe basis of filters
	result, err := collection.DeleteOne(ctx, filter, deleteOptions)
	if err != nil {
		log.Logger(ctx).Errorf("Error in Deleting mongo users : %v", err)
		return nil, err
	}
	return result, nil
}

//UpdateUser update user in database
func (userRepositoryImpl *UserRepositoryImpl) UpdateUser(ctx context.Context, filter bson.M, user models.User) (*models.User, error) {

	//update options
	updateOption := options.Update()

	//  mongo client connection
	mongoClient := userRepositoryImpl.mongoConn.NewMongoConn(ctx)
	defer mongoClient.Disconnect(ctx)

	//db and collection
	collection := mongoClient.Database("bloggingapp").Collection("users")

	update := bson.M{}
	//Checking for update only those values who has been requested for update
	//
	if len(user.FirstName) > 0 {
		update["first_name"] = user.FirstName
	}
	if len(user.LastName) > 0 {
		update["last_name"] = user.LastName
	}
	if len(user.Email) > 0 {
		update["user_email"] = user.Email
	}
	if len(user.Password) > 0 {
		update["password"] = user.Password
	}
	if len(user.Phone) > 0 {
		update["user_phone"] = user.Phone
	}
	update["times.updatedat"] = time.Now()

	data := bson.M{"$set": update}
	//fetch data from mongo database on tahe basis of filters
	result, err := collection.UpdateOne(ctx, filter, data, updateOption)
	if err != nil {
		log.Logger(ctx).Errorf("Error in finding mongo users : %v", err)
		return nil, err
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
	return nil, nil
}
