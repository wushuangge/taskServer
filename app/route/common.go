package route

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

func postFormToService(url string, timeStamp int64){
	creatTime := strconv.FormatInt(timeStamp,10)
	response, err := HttpsPostForm(url,"GetNewTasks", creatTime)
	if err != nil{
		log.Error(err.Error())
		return
	}
	var result []_struct.TaskFromService
	err = json.Unmarshal([]byte(response),&result)
	if err != nil{
		log.Error(err.Error())
		return
	}

	for _, v := range result {
		id := getMd5String(v.ProjectID + v.InstanceID + v.TaskID + v.TaskType)
		filter := bson.M{"_id": id}
		update := bson.D{
			{"$set", bson.D{
				{"_id", id},
				{"project_id", v.ProjectID},
				{"instance_id", v.InstanceID},
				{"task_id", v.TaskID},
				{"status", v.Status},
				{"time", v.CreateTime},
				{"type", v.TaskType},
				{"url", v.URL},
			}},
		}
		err :=mongodb.UpdateManagement(filter, update, true)
		if err != nil{
			log.Error(err.Error())
		}
	}
}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
