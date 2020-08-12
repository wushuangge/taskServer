package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
)

//insert
func InsertService(document interface{}) error{
	err := collMap["service"].InsertOne(document)
	return err
}

//query all
func QueryAllService() ([]_struct.TaskService, error){
	var all = make([]_struct.TaskService, 0)
	cursor,err := collMap["service"].Find(bson.D{})
	if err != nil {
		return all, err
	}
	if err := cursor.Err(); err != nil {
		return all, err
	}

	for cursor.Next(context.Background()) {
		var taskService _struct.TaskService
		err = cursor.Decode(&taskService);
		if err == nil {
			all = append(all, taskService)
		}
	}
	cursor.Close(context.Background())
	return all, nil
}

//condition query
func QueryConditionService(key string, value interface{}) (string, error){
	cursor,err := collMap["service"].Find(bson.D{{key, value}});
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskService _struct.TaskService
		err = cursor.Decode(&taskService)
		if err == nil {
			all = append(all, &taskService)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//update
func UpdateService(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["service"].UpdateOne(filter, update, updateOpts)
	return err
}
