package controller

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

func downloadTask(url string, increment int64) {
	incrementStr := strconv.FormatInt(increment, 10)
	response, err := httpsPostForm(url, "GetAllTasks", incrementStr)
	if err != nil {
		log.Error(err.Error())
		return
	}
	var result []_struct.TaskFromService
	err = json.Unmarshal([]byte(response), &result)
	if err != nil {
		log.Error(err.Error())
		return
	}

	for _, v := range result {
		id := GetMd5String(v.ProjectID + v.InstanceID + v.TaskID + v.TaskType)
		filter := bson.M{"_id": id}
		update := bson.D{
			{"$set", bson.D{
				{"_id", id},
				{"project_id", v.ProjectID},
				{"instance_id", v.InstanceID},
				{"task_id", v.TaskID},
				{"status", v.Status},
				{"type", v.TaskType},
				{"url", v.URL},
			}},
		}
		err := mongodb.UpdateManagement(filter, update, true)
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func httpsPostForm(u string, operation string, increment string) (string, error) {
	formData := url.Values{
		"operation": {operation},
		"increment": {increment},
	}
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.PostForm(u, formData)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
