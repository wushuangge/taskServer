package nsq

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
)

func Test(){
	SendHeartBeat()
	SendTaskService()
}

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
}

func SendHeartBeat() {
	url := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	heartBeat := HeartBeat{
		ID:				"1234567",
	}

	jsons, err := json.Marshal(heartBeat)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("err : controllers - tool.go - StructI2json")
	}
	err = producer.Publish("heartBeat", []byte(jsons))
	if err != nil {
		panic(err)
	}
	producer.Stop()
}

func SendTaskService() {
	url := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	taskService := TaskService{
		ID:				"1234567",
		Name: 			"abc",
		Type:			"1",
		URL: 			"http://localhost:8080/v1/task",
		Description: 	"add data",
	}

	jsons, err := json.Marshal(taskService)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("err : controllers - tool.go - StructI2json")
	}
	err = producer.Publish("taskService", []byte(jsons))
	if err != nil {
		panic(err)
	}
	producer.Stop()
}