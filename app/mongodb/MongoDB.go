package mongodb

import (
	// "fmt"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	username = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "27017"
)

// MongoDB : class of Mongo
type MongoDB struct {
	username string
	password string
	ip       string
	port     string
	client   *mongo.Client
}

// SetLinkParameters ...
func (mongoDB *MongoDB) SetLinkParameters(username string, password string, ip string, port string) {
	mongoDB.username = username
	mongoDB.password = password
	mongoDB.ip = ip
	mongoDB.port = port
}

// GetDBUri ...
func (mongoDB *MongoDB) getDBUri() string {
	return "mongodb://" + mongoDB.username + ":" + mongoDB.password + "@" + mongoDB.ip + ":" + mongoDB.port
}

// GetDBByName ...
func (mongoDB *MongoDB) GetDBByName(name string) *DB {
	db := new(DB)
	db.name = name
	db.database = mongoDB.client.Database(name)
	return db
}

// ConnMongoDB ...
func ConnMongoDB() *MongoDB {
	mongoDB := new(MongoDB)
	mongoDB.SetLinkParameters(username, password, ip, port)
	mongoDB.client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoDB.getDBUri()))
	return mongoDB
}

// Ping ...
func (mongoDB *MongoDB) Ping() error {
	return mongoDB.client.Ping(context.Background(), readpref.Primary())
}
