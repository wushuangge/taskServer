package nsq

import (
	"errors"
	"github.com/nsqio/go-nsq"
)

type NSQProducer struct {
	p *nsq.Producer
}

var producer *NSQProducer

func InitNSQ(addr string) error{
	producer = new(NSQProducer)
	producer.p = newNSQProducer(addr)
	if producer.p == nil {
		return errors.New("init nsq err")
	}
	return nil
}

func StopNSQ()  {
	if producer.p != nil {
		producer.p.Stop()
	}
}

func newNSQProducer(url string) *nsq.Producer {
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		return nil
	}
	return producer
}

func Publish(topic string, body []byte) error {
	err := producer.p.Publish(topic, body)
	return err
}
