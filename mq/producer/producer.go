package producer

import (
	"context"
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/xxl6097/go-glog/glog"
)

type Producer struct {
	producer *rocketmq.Producer
}

func New() *Producer {
	p := &Producer{
		producer: nil,
	}
	return p
}

func (this *Producer) NewProducer(servers []string) error {
	//连接RocketMQ
	p, err := rocketmq.NewProducer(producer.WithNameServer(servers))
	if err != nil {
		glog.Error("NewProducer失败：", err)
		return err
	}
	this.producer = &p
	return nil
}

func (this *Producer) Start() error {
	if this.producer == nil {
		return errors.New("producer not new,please call function Connect()")
	}
	//启动
	err := (*this.producer).Start()
	if err != nil {
		glog.Error("启动producer错误：", err)
		return err
	}
	return nil
}

// SendAsync 异步方法
func (this *Producer) SendAsync(topic, json string, mq func(ctx context.Context, result *primitive.SendResult, err error)) error {
	if this.producer == nil {
		return errors.New("producer not new,please call function Connect()")
	}
	//实例化消息
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(json),
	}
	//同步发送
	err := (*this.producer).SendAsync(context.Background(), mq, msg)
	if err != nil {
		glog.Errorf("SendAsync message error: %s\n", err)
		return err
	}
	return nil
}

// SendSync 同步方法
func (this *Producer) SendSync(topic, json string) (*primitive.SendResult, error) {
	if this.producer == nil {
		return nil, errors.New("producer not new,please call function Connect()")
	}
	//实例化消息
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(json),
	}
	//同步发送
	res, err := (*this.producer).SendSync(context.Background(), msg)
	if err != nil {
		glog.Errorf("SendSync message error: %s\n", err)
		return nil, err
	} else {
		//fmt.Printf("send message success: result=%s\n", res.String())
		glog.Infof("send message success: result=%s\n", res.String())
	}
	return res, nil
}

func (this *Producer) Shutdown() error {
	if this.producer != nil {
		//关闭连接
		err := (*this.producer).Shutdown()
		if err != nil {
			//fmt.Printf("shutdown producer error: %s", err.Error())
			glog.Errorf("shutdown producer error: %s\n", err.Error())
			return err
		}
	} else {
		return errors.New("producer not new,please call function Connect()")
	}
	return nil
}
