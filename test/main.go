package main

import (
	"encoding/json"
	"taskdash/test/nsq"
	"time"
)

//Service
type TaskService struct {
	URL				string 	`json:"url"`         //url
	Name			string 	`json:"name"`        //名称
	Reserved		string 	`json:"reserved"`	 //预留
}

//HeartBeat
type HeartBeat struct {
	URL				string 	`json:"url"`         //url
	Time			int64 	`json:"time"`		 //发送心跳的时间
	Reserved		string 	`json:"reserved"`	 //预留
}

func main() {
	addr := "192.168.51.12:4150"
	err := nsq.InitNSQ(addr)
	if err != nil {
		panic(err)
	}
	//PushHeartBeat()
	PushTaskService()
	nsq.StopNSQ()
}

func PushHeartBeat() {
	now := time.Now().Unix() //获取时间戳
	heartBeat := HeartBeat{
		URL: 			"http://localhost:8080/rpost/task",
		Time:			now,
		Reserved:		"123",
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
		Name: 			"annotator",
		URL: 			"https://192.168.51.33:8080/TaskDash/",
		Reserved:		"",
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
