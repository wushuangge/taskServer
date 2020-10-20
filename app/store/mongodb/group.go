package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
)

//insert
func InsertGroup(document interface{}) error {
	err := collMap["group"].InsertOne(document)
	return err
}

//condition query
func QueryConditionGroup2json(filter interface{}) (string, error) {
	cursor, err := collMap["group"].Find(filter)
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var groupUser _struct.Group
		err = cursor.Decode(&groupUser)
		if err == nil {
			all = append(all, &groupUser)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

func QueryConditionGroup2Struct(filter interface{}) ([]_struct.Group, error) {
	var all = make([]_struct.Group, 0)
	cursor, err := collMap["group"].Find(filter)
	if err != nil {
		return all, err
	}
	if err := cursor.Err(); err != nil {
		return all, err
	}
	for cursor.Next(context.Background()) {
		var group _struct.Group
		err = cursor.Decode(&group)
		if err == nil {
			all = append(all, group)
		}
	}
	cursor.Close(context.Background())
	return all, nil
}

//update
func UpdateGroup(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["group"].UpdateOne(filter, update, updateOpts)
	return err
}

//delete
func DeleteGroup(filter interface{}) error {
	err := collMap["group"].DeleteOne(filter)
	return err
}
