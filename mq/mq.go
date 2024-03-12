package mq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-rocketmq/mq/cons"
	"github.com/xxl6097/go-rocketmq/mq/prod"
	"time"
)

type RocketMQ struct {
	producer   *prod.Producer
	consumer   *cons.Consumer
	isWaitChan chan bool
	consopts   []consumer.Option
	prodopts   []producer.Option
}

func NewMQ() *RocketMQ {
	c := &RocketMQ{
		consumer:   cons.New(),
		producer:   prod.New(),
		isWaitChan: make(chan bool),
	}
	return c
}

func (this *RocketMQ) NewRocketMQ(cons_opts []consumer.Option, prod_opts []producer.Option) (error, error) {
	this.consopts = cons_opts
	this.prodopts = prod_opts
	err0 := this.consumer.NewCustomConsumer(cons_opts...)
	if err0 != nil {
		glog.Error("NewConsumer failed ", err0.Error())
	}

	err1 := this.producer.NewCustomConsumer(prod_opts...)
	if err1 != nil {
		glog.Error("NewProducer failed ", err1.Error())
	}
	return err0, err1
}

func (this *RocketMQ) ReConn(err0, err1 error) {
	if err0 != nil {
		for {
			time.Sleep(time.Second * 5)
			err := this.consumer.NewCustomConsumer(this.consopts...)
			if err == nil {
				break
			} else {
				glog.Error("NewConsumer failed delay 5s retry ", err.Error())
			}
		}
	}
	if err1 != nil {
		for {
			time.Sleep(time.Second * 5)
			err := this.producer.NewCustomConsumer(this.prodopts...)
			if err == nil {
				break
			} else {
				fmt.Println("NewProducer failed delay 5s retry ", err.Error())
			}
		}
	}

}

func (this *RocketMQ) InitRocketMQ(servers []string, groupName string) {
	err := this.consumer.NewConsumer(servers, groupName)
	if err != nil {
		//fmt.Println("NewConsumer failed ", err.Error())
		glog.Error("NewConsumer failed ", err.Error())
		for {
			time.Sleep(time.Second * 5)
			err = this.consumer.NewConsumer(servers, groupName)
			if err == nil {
				break
			} else {
				//fmt.Println("NewConsumer failed delay 5s retry ", err.Error())
				glog.Error("NewConsumer failed delay 5s retry ", err.Error())
			}
		}
	}
	err = this.producer.NewProducer(servers)
	if err != nil {
		glog.Error("NewProducer failed ", err.Error())
		for {
			time.Sleep(time.Second * 5)
			err = this.producer.NewProducer(servers)
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
			glog.Error("prod start failed")
			err = this.consumer.Start()
			if err == nil {
				break
			} else {
				glog.Error("prod start failed delay 5s retry ", err.Error())
			}
		}
	}
	glog.Info("cons start sucess")
	err = this.producer.Start()
	if err != nil {
		for {
			time.Sleep(time.Second * 5)
			glog.Error("cons start failed")
			err = this.producer.Start()
			if err == nil {
				break
			} else {
				glog.Error("cons start failed delay 5s retry ", err.Error())
			}
		}
	}
	glog.Info("prod start sucess")
}

func (this *RocketMQ) Init() {
	glog.Info("RocketMQ init")
}

func (this *RocketMQ) Subscribe(topic string, _receiver cons.OnReceiver) {
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
