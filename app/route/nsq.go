package route

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"task/app/store/mongodb"
	_struct "task/app/struct"
)

type NSQTaskService struct {

}

type NSQHeartBeat struct {

}

var serviceHeartBeat map[string]int64

func TaskService(addr string, config *nsq.Config) {
	consumer, err := nsq.NewConsumer("taskService", "struggle", config)
	if nil != err {
		fmt.Println("err", err)
		return
	}

	consumer.AddHandler(&NSQTaskService{})
	err = consumer.ConnectToNSQD(addr)
	if nil != err {
		fmt.Println("err", err)
		return
	}
}

func HeartBeat(url string, config *nsq.Config) {
	consumer, err := nsq.NewConsumer("heartBeat", "struggle", config)
	if nil != err {
		fmt.Println("err", err)
		return
	}

	consumer.AddHandler(&NSQHeartBeat{})
	err = consumer.ConnectToNSQD(url)
	if nil != err {
		fmt.Println("err", err)
		return
	}
}

func (this *NSQTaskService) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive NSQTaskService", msg.NSQDAddress, "message:", string(msg.Body))

	taskService := _struct.TaskService{}
	err := json.Unmarshal(msg.Body,&taskService)
	if err!=nil{
		return err
	}
	//save data
	mongodb.InsertService(taskService);
	//post
	response, err := HttpPost(taskService.URL)
	taskMetadata := _struct.TaskMetadata{}
	err = json.Unmarshal([]byte(response),&taskMetadata)
	if err!=nil{
		return err
	}
	//save data
	mongodb.InsertMetadata(taskMetadata);
	return nil
}

func (this *NSQHeartBeat) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive NSQHeartBeat", msg.NSQDAddress, "message:", string(msg.Body))
	heartBeat:=_struct.HeartBeat{}
	err:=json.Unmarshal(msg.Body,&heartBeat)
	if err!=nil{
		return err
	}
	serviceHeartBeat[heartBeat.URL] = heartBeat.Time
	return nil
}
