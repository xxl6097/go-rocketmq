package mq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-rocketmq/mq/consumer"
	"github.com/xxl6097/go-rocketmq/mq/producer"
	"time"
)

type RocketMQ struct {
	producer   *producer.Producer
	consumer   *consumer.Consumer
	isWaitChan chan bool
}

func NewMQ() *RocketMQ {
	c := &RocketMQ{
		consumer:   consumer.New(),
		producer:   producer.New(),
		isWaitChan: make(chan bool),
	}
	return c
}

func (this *RocketMQ) InitRocketMQ(server, groupName string) {
	err := this.consumer.NewConsumer(server, groupName)
	if err != nil {
		//fmt.Println("NewConsumer failed ", err.Error())
		glog.Error("NewConsumer failed ", err.Error())
		for {
			time.Sleep(time.Second * 5)
			err = this.consumer.NewConsumer(server, groupName)
			if err == nil {
				break
			} else {
				//fmt.Println("NewConsumer failed delay 5s retry ", err.Error())
				glog.Error("NewConsumer failed delay 5s retry ", err.Error())
			}
		}
	}
	err = this.producer.NewProducer(server)
	if err != nil {
		glog.Error("NewProducer failed ", err.Error())
		for {
			time.Sleep(time.Second * 5)
			err = this.producer.NewProducer(server)
			if err == nil {
				break
			} else {
				fmt.Println("NewProducer failed delay 5s retry ", err.Error())
			}
		}
	}
}

func (this *RocketMQ) Start() {
	err := this.consumer.Start()
	if err != nil {
		for {
			time.Sleep(time.Second * 5)
			glog.Error("producer start failed")
			err = this.consumer.Start()
			if err == nil {
				break
			} else {
				glog.Error("producer start failed delay 5s retry ", err.Error())
			}
		}
	}
	glog.Info("consumer start sucess")
	err = this.producer.Start()
	if err != nil {
		for {
			time.Sleep(time.Second * 5)
			glog.Error("consumer start failed")
			err = this.producer.Start()
			if err == nil {
				break
			} else {
				glog.Error("consumer start failed delay 5s retry ", err.Error())
			}
		}
	}
	glog.Info("producer start sucess")
}

func (this *RocketMQ) Subscribe(topic string, _receiver consumer.OnReceiver) {
	err := this.consumer.Subscribe(topic, _receiver)
	if err != nil {
		for {
			time.Sleep(time.Second * 5)
			glog.Error("Subscribe failed")
			err = this.producer.Start()
			if err == nil {
				break
			} else {
				glog.Error("Subscribe failed delay 5s retry ", err.Error())
			}
		}
	}
}

func (this *RocketMQ) SendSync(topic, json string) (*primitive.SendResult, error) {
	return this.producer.SendSync(topic, json)
}

func (this *RocketMQ) SendAsync(topic, json string, mq func(ctx context.Context, result *primitive.SendResult, err error)) error {
	return this.producer.SendAsync(topic, json, mq)
}

func (this *RocketMQ) Shutdown() {
	if this.producer != nil {
		this.producer.Shutdown()
	}
	if this.consumer != nil {
		this.consumer.Shutdown()
	}
}

func (this *RocketMQ) Wait(ctx context.Context) {
	for {
		if <-this.isWaitChan {
			select {
			case <-ctx.Done():
				this.Shutdown()
				glog.Info("任务结束了...")
				return
			}
		}
	}
}
