package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
)

func InsertService(taskService _struct.TaskService) error{
	err := collMap["service"].InsertOne(taskService);
	return err
}

func QueryAllService() (string, error){
	cursor,err := collMap["service"].Find(bson.D{});
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
		var taskService _struct.TaskService
		if err = cursor.Decode(&taskService); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, &taskService)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

func QueryConditionService(key string, value interface{}) (string, error){
	cursor,err := collMap["service"].Find(bson.D{{key, value}});
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
		var taskService _struct.TaskService
		if err = cursor.Decode(&taskService); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, &taskService)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

func UpdateService(taskService _struct.TaskService) error {
	update := bson.M{"$set": taskService}
	updateOpts := options.Update().SetUpsert(true)
	err := collMap["service"].UpdateOne(bson.M{"url": taskService.URL}, update, updateOpts)
	return err
}