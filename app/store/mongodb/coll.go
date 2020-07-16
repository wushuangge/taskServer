package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (coll *Coll) InsertOne(document interface{}) error {
	_, err := coll.collection.InsertOne(context.Background(), document)
	return err
}

func (coll *Coll) Find(filter interface{}) (*mongo.Cursor, error) {
	cursor, err := coll.collection.Find(context.Background(), filter)
	return cursor, err
}

func (coll *Coll) DeleteOne(p primitive.M) error{
	_, err := coll.collection.DeleteOne(context.Background(), p)
	return err
}

func (coll *Coll) UpdateOne(filter interface{}, update interface{}, opts *options.UpdateOptions) error {
	_, err := coll.collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}
