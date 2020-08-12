package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
	"time"
)

//insert
func InsertManagement(document interface{}) error{
	err := collMap["management"].InsertOne(document)
	return err
}

//query all
func QueryAllManagement() (string, error){
	cursor,err := collMap["management"].Find(bson.D{})
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskManagement _struct.TaskManagement
		err = cursor.Decode(&taskManagement)
		if err == nil {
			all = append(all, &taskManagement)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//condition query
func QueryConditionManagement(filter interface{}) (string, error){
	cursor,err := collMap["management"].Find(filter)
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskManagement _struct.TaskManagement
		err = cursor.Decode(&taskManagement)
		if err == nil {
			all = append(all, &taskManagement)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//page query
func QueryPagingManagement(limit int64, skip int64, filter interface{}) (string, error){
	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()
	var findoptions *options.FindOptions
	if limit > 0 {
		findoptions = &options.FindOptions{}
		findoptions.SetLimit(limit)
		findoptions.SetSkip(limit * skip)
	}
	//cursor, err := collMap["management"].FindPaging(ctx, bson.M{}, findoptions)
	cursor, err := collMap["management"].FindPaging(ctx, filter, findoptions)
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskManagement _struct.TaskManagement
		err = cursor.Decode(&taskManagement)
		if err == nil {
			all = append(all, &taskManagement)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//update
func UpdateManagement(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["management"].UpdateOne(filter, update, updateOpts)
	return err
}