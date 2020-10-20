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
	if err != nil {
		log.Error(err)
		return
	}

	id := GetMd5String(taskFromService.ProjectID +
		taskFromService.InstanceID + taskFromService.TaskID + taskFromService.TaskType)

	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"_id", id},
			{"project_id", taskFromService.ProjectID},
			{"instance_id", taskFromService.InstanceID},
			{"task_id", taskFromService.TaskID},
			{"status", taskFromService.Status},
			{"type", taskFromService.TaskType},
			{"url", taskFromService.URL},
		}},
	}
	mongodb.UpdateManagement(filter, update, true)
}

func UpdateStatus(data []byte) {
	taskFromService := _struct.TaskFromService{}
	err := json.Unmarshal(data, &taskFromService)
	if err != nil {
		log.Error(err)
		return
	}

	id := GetMd5String(taskFromService.ProjectID +
		taskFromService.InstanceID + taskFromService.TaskID + taskFromService.TaskType)

	if taskFromService.Status == "已完成" {
		filter := bson.M{"_id": id}
		res, err := mongodb.QueryConditionManagement2Interface(filter)
		if err != nil {
			return
		}
		if len(res) == 0 {
			return
		}
		update := bson.D{
			{"$set", bson.D{
				{"_id", id},
				{"project_id", res[0].ProjectID},
				{"instance_id", res[0].InstanceID},
				{"task_id", res[0].TaskID},
				{"status", "已完成"},
				{"time", res[0].CreateTime},
				{"type", res[0].TaskType},
				{"url", res[0].URL},
				{"user", res[0].User},
			}},
		}
		err = mongodb.DeleteManagement(filter)
		if err != nil {
			return
		}
		mongodb.UpdateBackup(filter, update, true)
	} else {
		filter := bson.M{"_id": id}
		update := bson.D{
			{"$set", bson.D{
				{"status", taskFromService.Status},
			}},
		}
		mongodb.UpdateManagement(filter, update, false)
	}
}
