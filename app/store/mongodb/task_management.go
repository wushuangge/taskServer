package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
)

func InsertManagement(taskManagement _struct.TaskManagement){
	err :=collMap["management"].InsertOne(taskManagement);
	if err != nil {
		fmt.Println(err)
	}
}

func QueryAllManager() (string, error){
	cursor,err := collMap["management"].Find(bson.D{});
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
		var taskManagement _struct.TaskManagement
		if err = cursor.Decode(&taskManagement); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, &taskManagement)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

func QueryConditionManager(key string, value interface{}) (string, error){
	cursor,err := collMap["management"].Find(bson.D{{key, value}});
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
		var taskManagement _struct.TaskManagement
		if err = cursor.Decode(&taskManagement); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, &taskManagement)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

func UpdateManagement(taskManagement _struct.TaskManagement) error {
	update := bson.M{"$set": taskManagement}
	updateOpts := options.Update().SetUpsert(true)
	err := collMap["metadata"].UpdateOne(bson.M{"task_id": taskManagement.TaskID}, update, updateOpts)
	return err
}