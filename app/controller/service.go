package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

func ServiceRegister(data []byte) {
	taskService := _struct.TaskService{}
	err := json.Unmarshal(data, &taskService)
	if err != nil{
		log.Error(err)
		return
	}

	filter := bson.M{"_id": taskService.URL}
	update := bson.M{"$set": taskService}
	mongodb.UpdateService(filter, update, true)

	downloadTask(taskService.URL, 0)
}