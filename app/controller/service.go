package controller

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

func ServiceRegister(data []byte) {
	taskService := _struct.TaskService{}
	err := json.Unmarshal(data, &taskService)
	if err != nil {
		log.Error(err)
		return
	}

	if isEmptyString(taskService.Network) ||
		isEmptyString(taskService.Address) ||
		isEmptyString(taskService.Path["taskDash"]) ||
		isEmptyString(taskService.Name) {
		return
	}

	id := GetMd5String(taskService.Network + taskService.Address)
	taskService.ID = id
	taskService.Enable = true
	filter := bson.M{"_id": id}
	update := bson.M{"$set": taskService}
	mongodb.UpdateService(filter, update, true)

	mutex.Lock()
	mapHeartbeat[id] = fillTimer(true, counterRatio)
	mutex.Unlock()

	url := taskService.Network + taskService.Address + "/" + taskService.Path["taskDash"] + "/"
	fmt.Println(url)
	downloadTask(url, 0)
}

func ServiceUpdate(id string, enable bool) {
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"enable", enable},
		}},
	}
	mongodb.UpdateService(filter, update, false)
}

func isEmptyString(s string) bool {
	var isEmpty bool
	if len(s) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}
	return isEmpty
}
