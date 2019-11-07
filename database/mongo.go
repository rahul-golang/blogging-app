package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoDBConn holding ref to mongo Client
type MongoDBConn struct {
	mongoConn *mongo.Client
}

//MongoDBConnInterface defined behaviour of MongoDbConn
type MongoDBConnInterface interface {
	NewMongoConn() *mongo.Client
}

//NewMongoDBConn inject dependacies
func NewMongoDBConn() MongoDBConnInterface {
	return &MongoDBConn{}
}

//ClientOptions return MongoClientOptions
func (mongoDBConn MongoDBConn) ClientOptions() *options.ClientOptions {
	return options.Client().ApplyURI("mongodb://localhost:27017")
}

//NewMongoConn return new client for query oprations
func (mongoDBConn MongoDBConn) NewMongoConn() *mongo.Client {
	client, err := mongo.NewClient(mongoDBConn.ClientOptions())
	if err != nil {
		panic("Error in create Database Connection")
	}
	return client
}
