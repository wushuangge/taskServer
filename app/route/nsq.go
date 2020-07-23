package route

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

type NSQTaskService struct {

}

type NSQHeartBeat struct {

}

type NSQUpdateStatus struct {

}

var serviceHeartBeat map[string]int64

func TaskService(addr string, config *nsq.Config) error{
	consumer, err := nsq.NewConsumer("taskService", "struggle", config)
	if nil != err {
		fmt.Println("err", err)
		return err
	}

	consumer.AddHandler(&NSQTaskService{})
	err = consumer.ConnectToNSQD(addr)
	if nil != err {
		fmt.Println("err", err)
		return err
	}
	return err
}

func UpdateStatus(url string, config *nsq.Config) error {
	consumer, err := nsq.NewConsumer("updateStatus", "struggle", config)
	if nil != err {
		fmt.Println("err", err)
		return err
	}

	consumer.AddHandler(&NSQUpdateStatus{})
	err = consumer.ConnectToNSQD(url)
	if nil != err {
		fmt.Println("err", err)
		return err
	}
	return err
}

func HeartBeat(url string, config *nsq.Config) error {
	consumer, err := nsq.NewConsumer("heartBeat", "struggle", config)
	if nil != err {
		fmt.Println("err", err)
		return err
	}

	consumer.AddHandler(&NSQHeartBeat{})
	err = consumer.ConnectToNSQD(url)
	if nil != err {
		fmt.Println("err", err)
		return err
	}
	return err
}

func (this *NSQTaskService) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive NSQTaskService", msg.NSQDAddress, "message:", string(msg.Body))

	taskService := _struct.TaskService{}
	err := json.Unmarshal(msg.Body,&taskService)
	if err!=nil{
		return err
	}
	//存在则更新，不存在则插入
	mongodb.UpdateService(taskService);
	////post
	response, err := HttpsPostForm(taskService.URL,"GetNewTasks","0")
	var result []_struct.TaskFromService
	err = json.Unmarshal([]byte(response),&result)
	if err!=nil{
		return err
	}

	for _, v := range result {
		taskMetadata := _struct.TaskMetadata{}
		taskMetadata.ID 		= GetMd5String(v.ProjectID + v.InstanceID + v.TaskID)
		taskMetadata.ProjectID 	= v.ProjectID
		taskMetadata.InstanceID = v.InstanceID
		taskMetadata.TaskID 	= v.TaskID
		taskMetadata.Status 	= v.Status
		taskMetadata.CreateTime = v.CreateTime
		taskMetadata.DataType 	= v.DataType
		taskMetadata.Reserved 	= ""
		//存在则更新，不存在则插入
		mongodb.UpdateMetadata(taskMetadata);
	}
	return nil
}

func (this *NSQUpdateStatus) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive UpdateStatus", msg.NSQDAddress, "message:", string(msg.Body))
	taskFromService := _struct.TaskFromService{}
	err := json.Unmarshal(msg.Body,&taskFromService)
	if err!=nil{
		return err
	}
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


func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
