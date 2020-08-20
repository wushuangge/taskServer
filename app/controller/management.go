package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

func TaskRegister(data []byte) {
	taskFromService := _struct.TaskFromService{}
	err := json.Unmarshal(data, &taskFromService)
	if err != nil{
		log.Error(err)
		return
	}

	id := getMd5String(taskFromService.ProjectID +
		taskFromService.InstanceID + taskFromService.TaskID + taskFromService.TaskType)

	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"_id", id},
			{"project_id", taskFromService.ProjectID},
			{"instance_id", taskFromService.InstanceID},
			{"task_id", taskFromService.TaskID},
			{"status", taskFromService.Status},
			{"time", taskFromService.CreateTime},
			{"type", taskFromService.TaskType},
			{"url", taskFromService.URL},
			{"increment", taskFromService.Increment},
		}},
	}
	mongodb.UpdateManagement(filter, update, true)
}

func UpdateStatus(data []byte) {
	taskFromService := _struct.TaskFromService{}
	err := json.Unmarshal(data, &taskFromService)
	if err != nil{
		log.Error(err)
		return
	}

	id := getMd5String(taskFromService.ProjectID +
		taskFromService.InstanceID + taskFromService.TaskID + taskFromService.TaskType)

	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"status", taskFromService.Status},
		}},
	}
	mongodb.UpdateManagement(filter, update, false)
}