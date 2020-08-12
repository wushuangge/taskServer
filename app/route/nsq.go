package route

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
	"taskdash/app/controller"
)

type NSQServiceRegister struct {
}

type NSQTaskRegister struct {
}

type NSQUpdateStatus struct {
}

type NSQHeartBeat struct {
}

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
	controller.ServiceRegister(msg.Body)
	return nil
}

func (this *NSQTaskRegister) HandleMessage(msg *nsq.Message) error {
	fmt.Println("recv NSQTaskRegister", string(msg.Body))
	controller.TaskRegister(msg.Body)
	return nil
}

func (this *NSQUpdateStatus) HandleMessage(msg *nsq.Message) error {
	fmt.Println("recv NSQUpdateStatus", string(msg.Body))
	controller.UpdateStatus(msg.Body)
	return nil
}

func (this *NSQHeartBeat) HandleMessage(msg *nsq.Message) error {
	fmt.Println("recv NSQHeartBeat",string(msg.Body))
	controller.HeartBeat(msg.Body)
	return nil
}

