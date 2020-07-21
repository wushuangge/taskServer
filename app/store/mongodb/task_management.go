package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	_struct "taskdash/app/struct"
	"time"
)

//插入
func InsertManagement(taskManagement _struct.TaskManagement){
	err :=collMap["management"].InsertOne(taskManagement);
	if err != nil {
		fmt.Println(err)
	}
}

//查询所有
func QueryAllManagement() (string, error){
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

//条件查询
func QueryConditionManagement(key string, value interface{}) (string, error){
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

//更新，存在则更新，不存在则插入
func UpdateManagement(taskManagement _struct.TaskManagement) error {
	update := bson.M{"$set": taskManagement}
	updateOpts := options.Update().SetUpsert(true)
	err := collMap["management"].UpdateOne(bson.M{"_id": taskManagement.TaskID}, update, updateOpts)
	return err
}

//分页查询
func QueryPagingManagement(limit int64, index int64) (string, error){
	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()
	var findoptions *options.FindOptions
	if limit > 0 {
		findoptions = &options.FindOptions{}
		findoptions.SetLimit(limit)
		findoptions.SetSkip(limit * index)
	}
	cursor, err := collMap["management"].FindPaging(ctx, bson.M{}, findoptions)
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

//test
func TestInsertManagement(){
	for i := 0; i < 20; i++ {
		var tmp string
		tmp = strconv.Itoa(i)
		id := "sy-hn-" + tmp
		taskManagement := _struct.TaskManagement{
			TaskID:      id,
			DataType:    "2",
			Status:		 "running",
			User:        "zhangsan",
			Checker:	 "lisi",
			Group:	 	 "mx",
			Reserved:    "",
		}
		UpdateManagement(taskManagement);
	}

	res,_ := QueryConditionManagement("user","zhangsan")
	fmt.Println(res)
}