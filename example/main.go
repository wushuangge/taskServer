package main

import (
	"encoding/json"
	"taskdash/test/nsq"
	"time"
)

//Service
type TaskService struct {
	ID       string            `bson:"_id"`      //id
	Network  string            `bson:"network"`  //网络类型
	Address  string            `bson:"address"`  //地址
	Name     string            `bson:"name"`     //名称
	Path     map[string]string `bson:"path"`     //路径
	Reserved string            `bson:"reserved"` //预留
}

//HeartBeat
type HeartBeat struct {
	ID       string      `bson:"_id"`      //id
	Network  string      `bson:"network"`  //网络类型
	Address  string      `bson:"address"`  //地址
	Time     int64       `bson:"time"`     //发送心跳的时间
	Reserved interface{} `bson:"reserved"` //预留
}

type TaskFromService struct {
	ID         string `bson:"_id"`
	ProjectID  string
	InstanceID string
	TaskID     string
	Status     string
	TaskType   string
	URL        string
	EditInfo   map[string]string
	CreateTime int64
}

func main() {
	addr := "192.168.51.12:4150"
	err := nsq.InitNSQ(addr)
	if err != nil {
		panic(err)
	}
	//PushHeartBeat()
	PushServiceRegister()
	//PushTaskRegister()
	nsq.StopNSQ()
}

func PushHeartBeat() {
	now := time.Now().Unix() //获取时间戳
	heartBeat := HeartBeat{
		Network:  "https://",
		Address:  "192.168.51.12:8080",
		Time:     now,
		Reserved: "123",
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

func PushTaskRegister() {
	taskFromService := TaskFromService{
		ID:  "111111111",
		URL: "https://192.168.51.33:8080/TaskDash/",
	}

	jsons, err := json.Marshal(taskFromService)

	if err != nil {
		panic(err)
	}
	err = nsq.Publish("taskRegister", []byte(jsons))
	if err != nil {
		panic(err)
	}
}

func PushServiceRegister() {
	var s = make(map[string]string)
	s["mapDeck"] = "taskDB"
	s["taskDash"] = "taskDB"
	taskService := TaskService{
		Network: "https://",
		Address: "192.168.51.33:8080",
		Name:    "test",
		Path:    s,
	}
	jsons, err := json.Marshal(taskService)

	if err != nil {
		panic(err)
	}
	err = nsq.Publish("serviceRegister", []byte(jsons))
	if err != nil {
		panic(err)
	}
}
