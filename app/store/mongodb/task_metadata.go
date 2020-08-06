package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
	"time"
)

//insert
func InsertMetadata(document interface{}) error {
	err :=collMap["metadata"].InsertOne(document)
	return err
}

//query all
func QueryAllMetadata() (string, error){
	cursor,err := collMap["metadata"].Find(bson.D{})
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskMetadata _struct.TaskMetadata
		err = cursor.Decode(&taskMetadata)
		if err == nil {
			all = append(all, &taskMetadata)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//condition query
func QueryConditionMetadata2Json(filter interface{}) (string, error){
	cursor,err := collMap["metadata"].Find(filter)
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskMetadata _struct.TaskMetadata
		err = cursor.Decode(&taskMetadata)
		if err == nil {
			all = append(all, &taskMetadata)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//condition query
func QueryConditionMetadata2Interface(filter interface{}) ([]_struct.TaskMetadata, error){
	var all = make([]_struct.TaskMetadata, 0)
	cursor, err := collMap["metadata"].Find(filter)
	if err != nil {
		return all, err
	}
	if err := cursor.Err(); err != nil {
		return all, err
	}
	for cursor.Next(context.Background()) {
		var taskMetadata _struct.TaskMetadata
		err = cursor.Decode(&taskMetadata)
		if err == nil {
			all = append(all, taskMetadata)
		}
	}
	cursor.Close(context.Background())
	return all, nil
}

//page query
func QueryPagingMetadata(limit int64, skip int64, key string, value interface{}) (string, error){
	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()
	var findoptions *options.FindOptions
	if limit > 0 {
		findoptions = &options.FindOptions{}
		findoptions.SetLimit(limit)
		findoptions.SetSkip(limit * skip)
	}
	//cursor, err := collMap["metadata"].FindPaging(ctx, bson.M{}, findoptions)
	cursor, err := collMap["metadata"].FindPaging(ctx, bson.D{{key, value}}, findoptions)
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var taskMetadata _struct.TaskMetadata
		err = cursor.Decode(&taskMetadata)
		if err == nil {
			all = append(all, &taskMetadata)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

//update
func UpdateMetadata(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["metadata"].UpdateOne(filter, update, updateOpts)
	return err
}