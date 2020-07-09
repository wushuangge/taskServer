package main

import (
	"encoding/json"
	"task/test/nsq"
	"time"
)

//Service
type TaskService struct {
	ID				string 	`json:"id"`		     //Service id(唯一标识)
	Name 			string 	`json:"name"`        //名称
	Type        	string 	`json:"type"`        //类别
	URL		        string 	`json:"url"`         //url
	Description		string 	`json:"description"` //描述
}

//HeartBeat
type HeartBeat struct {
	ID				string 	`json:"id"`		     //Service id(唯一标识)
	Time			int64 	`json:"time"`		 //发送心跳的时间
	Reserved		string 	`json:"reserved"`	 //预留
}

func main() {
	url := "127.0.0.1:4150"
	nsq.InitNSQ(url)

	PushHeartBeat()
	PushTaskService()
}

func PushHeartBeat() {
	now := time.Now().Unix() //获取时间戳
	heartBeat := HeartBeat{
		ID:				"123456",
		Time:			now,
		Reserved:		"",
	}
	jsons, err := json.Marshal(heartBeat)
	if err != nil {
		panic(err)
	}
	err = nsq.Publish("heartBeat", []byte(jsons))
	if err != nil {
		panic(err)
	}
}

func PushTaskService() {
	taskService := TaskService{
		ID:				"123456",
		Name: 			"abc",
		Type:			"1",
		URL: 			"http://localhost:8080/v1/task",
		Description: 	"add data",
	}

	jsons, err := json.Marshal(taskService)

	if err != nil {
		panic(err)
	}
	err = nsq.Publish("taskService", []byte(jsons))
	if err != nil {
		panic(err)
	}
}