package route

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

type NSQServiceRegister struct {
}

type NSQTaskRegister struct {
}

type NSQUpdateStatus struct {
}

type NSQHeartBeat struct {
}

var serviceHeartBeat map[string]int64

func ServiceRegister(addr string, config *nsq.Config) error{
	consumer, err := nsq.NewConsumer("serviceRegister", "struggle", config)
	if nil != err {
		log.Error(err)
		return err
	}

	consumer.AddHandler(&NSQServiceRegister{})
	err = consumer.ConnectToNSQD(addr)
	if nil != err {
		log.Error(err)
		return err
	}
	return err
}

func TaskRegister(url string, config *nsq.Config) error {
	consumer, err := nsq.NewConsumer("taskRegister", "struggle", config)
	if nil != err {
		log.Error(err)
		return err
	}

	consumer.AddHandler(&NSQTaskRegister{})
	err = consumer.ConnectToNSQD(url)
	if nil != err {
		log.Error(err)
		return err
	}
	return err
}

func UpdateStatus(url string, config *nsq.Config) error {
	consumer, err := nsq.NewConsumer("updateStatus", "struggle", config)
	if nil != err {
		log.Error(err)
		return err
	}

	consumer.AddHandler(&NSQUpdateStatus{})
	err = consumer.ConnectToNSQD(url)
	if nil != err {
		log.Error(err)
		return err
	}
	return err
}

func HeartBeat(url string, config *nsq.Config) error {
	consumer, err := nsq.NewConsumer("heartBeat", "struggle", config)
	if nil != err {
		log.Error(err)
		return err
	}

	consumer.AddHandler(&NSQHeartBeat{})
	err = consumer.ConnectToNSQD(url)
	if nil != err {
		log.Error(err)
		return err
	}
	return err
}

func (this *NSQServiceRegister) HandleMessage(msg *nsq.Message) error {
	fmt.Println("recv NSQServiceRegister", string(msg.Body))

	taskService := _struct.TaskService{}
	err := json.Unmarshal(msg.Body,&taskService)
	if err != nil{
		log.Error(err)
		return err
	}

	filter := bson.M{"_id": taskService.URL}
	update := bson.M{"$set": taskService}
	mongodb.UpdateService(filter, update, true)

	postFormToService(taskService.URL, 0)
	return nil
}

func (this *NSQTaskRegister) HandleMessage(msg *nsq.Message) error {
	fmt.Println("recv NSQTaskRegister", string(msg.Body))
	taskFromService := _struct.TaskFromService{}
	err := json.Unmarshal(msg.Body,&taskFromService)
	if err != nil{
		log.Error(err)
		return err
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
		}},
	}
	mongodb.UpdateManagement(filter, update, true)

	return nil
}

func (this *NSQUpdateStatus) HandleMessage(msg *nsq.Message) error {
	fmt.Println("recv NSQUpdateStatus", string(msg.Body))
	taskFromService := _struct.TaskFromService{}
	err := json.Unmarshal(msg.Body,&taskFromService)
	if err != nil{
		log.Error(err)
		return err
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

	return nil
}

func (this *NSQHeartBeat) HandleMessage(msg *nsq.Message) error {
	fmt.Println("recv NSQHeartBeat",string(msg.Body))
	heartBeat:=_struct.HeartBeat{}
	err:=json.Unmarshal(msg.Body,&heartBeat)
	if err != nil{
		log.Error(err)
		return err
	}
	serviceHeartBeat[heartBeat.URL] = heartBeat.Time
	return nil
}

