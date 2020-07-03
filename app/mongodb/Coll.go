package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Coll : class of collection
type Coll struct {
	collection *mongo.Collection
}

// GetCur ...
func (coll *Coll) GetCur() *mongo.Cursor {
	cur, err := coll.collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("err: Coll.go GetCur")
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
		fmt.Println("err: Coll.go GetCur")
	}
	return cur
}

// DeleteOne ...
func (coll *Coll) DeleteOne(p primitive.M) {
	_, err := coll.collection.DeleteOne(context.Background(), p)
	if err != nil {
		fmt.Println("删除错误：Coll.go - DeleteOne")
	}
}
