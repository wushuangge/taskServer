package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
)

func InsertMetadata(taskMetadata _struct.TaskMetadata){
	err :=collMap["metadata"].InsertOne(taskMetadata);
	if err != nil {
		fmt.Println(err)
	}
}

func QueryAllMetadata() (string, error){
	cursor,err := collMap["metadata"].Find(bson.D{});
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if err := cursor.Err(); err != nil {
		fmt.Println(err)
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskMetadata _struct.TaskMetadata
		if err = cursor.Decode(&taskMetadata); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, &taskMetadata)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

func QueryConditionMetadata(key string, value interface{}) (string, error){
	cursor,err := collMap["metadata"].Find(bson.D{{key, value}});
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if err := cursor.Err(); err != nil {
		fmt.Println(err)
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskMetadata _struct.TaskMetadata
		if err = cursor.Decode(&taskMetadata); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, &taskMetadata)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

func UpdateMetadata(taskMetadata _struct.TaskMetadata) error {
	update := bson.M{"$set": taskMetadata}
	updateOpts := options.Update().SetUpsert(true)
	err := collMap["metadata"].UpdateOne(bson.M{"task_id": taskMetadata.TaskID}, update, updateOpts)
	return err
}