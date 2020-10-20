package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	_struct "taskdash/app/struct"
)

//insert
func InsertUser(document interface{}) error {
	err := collMap["user"].InsertOne(document)
	return err
}

//condition query
func QueryConditionUser2json(filter interface{}) (string, error) {
	cursor, err := collMap["user"].Find(filter)
	if err != nil {
		return "", err
	}
	if err := cursor.Err(); err != nil {
		return "", err
	}
	var all = make([]interface{}, 0)

	for cursor.Next(context.Background()) {
		var user _struct.User
		err = cursor.Decode(&user)
		if err == nil {
			all = append(all, &user)
		}
	}
	cursor.Close(context.Background())
	return Interfaces2json(all), nil
}

func QueryConditionUser2Struct(filter interface{}) ([]_struct.User, error) {
	var all = make([]_struct.User, 0)
	cursor, err := collMap["user"].Find(filter)
	if err != nil {
		return all, err
	}
	if err := cursor.Err(); err != nil {
		return all, err
	}
	for cursor.Next(context.Background()) {
		var user _struct.User
		err = cursor.Decode(&user)
		if err == nil {
			all = append(all, user)
		}
	}
	cursor.Close(context.Background())
	return all, nil
}

//update
func UpdateUser(filter interface{}, update interface{}, setUpsert bool) error {
	updateOpts := options.Update().SetUpsert(setUpsert)
	err := collMap["user"].UpdateOne(filter, update, updateOpts)
	return err
}

//delete
func DeleteUser(filter interface{}) error {
	err := collMap["user"].DeleteOne(filter)
	return err
}
