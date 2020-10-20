package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
	"time"
)

//insert
func InsertBackup(document interface{}) error {
	err := collMap["backup"].InsertOne(document)
	return err
}

//query all
func QueryAllBackup() (string, error) {
	cursor, err := collMap["backup"].Find(bson.D{})
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
func QueryConditionBackup(filter interface{}) (string, error) {
	cursor, err := collMap["backup"].Find(filter)
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

func QueryBackupNum(filter interface{}) (int64, error) {
	return collMap["backup"].CountDocuments(filter)
}

//page query
func QueryPagingBackup(limit int64, skip int64, filter interface{}) ([]interface{}, error) {
	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()
	var findoptions *options.FindOptions
	if limit > 0 {
		findoptions = &options.FindOptions{}
		findoptions.SetLimit(limit)
		findoptions.SetSkip(skip)
	}
	var all = make([]interface{}, 0)
	cursor, err := collMap["backup"].FindPaging(ctx, filter, findoptions)
	if err != nil {
		return all, err
	}
	if err := cursor.Err(); err != nil {
		return all, err
	}
	for cursor.Next(context.Background()) {
		var taskManagement _struct.TaskManagement
		err = cursor.Decode(&taskManagement)
		if err == nil {
			all = append(all, &taskManagement)
		}
	}
	cursor.Close(context.Background())
	return all, nil
}

//update
func UpdateBackup(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["backup"].UpdateOne(filter, update, updateOpts)
	return err
}
