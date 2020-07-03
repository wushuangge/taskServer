package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// DB : class of database
type DB struct {
	name     string
	database *mongo.Database
}

// GetCollByName ...
func (db *DB) GetCollByName(collName string) *Coll {
	coll := new(Coll)
	coll.collection = db.database.Collection(collName)
	return coll
}

// DeleteColl ...
func (db *DB) DeleteColl(coll string) error {
	return db.GetCollByName(coll).collection.Drop(context.Background())
}
