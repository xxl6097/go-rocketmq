package consumer

import (
	"context"
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/xxl6097/go-glog/glog"
)

//type OnReceiver interface {
//	onReceive(msg *primitive.MessageExt)
//}

type OnReceiver func(msg *primitive.MessageExt)

type Consumer struct {
	consumer *rocketmq.PushConsumer
}

func New() *Consumer {
	c := &Consumer{
		consumer: nil,
	}
	return c
}

func (this *Consumer) NewCustomConsumer(c *rocketmq.PushConsumer) {
	this.consumer = c
}

func (this *Consumer) NewConsumer(servers []string, groupName string) error {
	//启动recketmq并设置负载均衡的Group
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(servers),
		consumer.WithGroupName(groupName),
	)
	if err != nil {
		glog.Error("NewPushConsumer失败：", err)
		return err
	}
	// 设置消息模式为广播模式
	//c.SetModel(consumer.BroadCasting)
	this.consumer = &c
	return nil
}

func (this *Consumer) Shutdown() error {
	if this.consumer != nil {
		//关闭连接
		err := (*this.consumer).Shutdown()
		if err != nil {
			glog.Errorf("shutdown consumer error: %s\n", err.Error())
			return err
		}
	} else {
		return errors.New("consumer not new,please call Connect() function")
	}
	return nil
}

func (this *Consumer) Subscribe(topic string, _receiver OnReceiver) error {
	if this.consumer == nil {
		return errors.New("consumer is nil")
	}
	if err := (*this.consumer).Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, value := range ext {
			//fmt.Printf("--->%v \n", value)
			if _receiver != nil {
				_receiver(value)
			}
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		glog.Error("Subscribe failed->", err.Error())
		return err
	}
	//fmt.Printf("Subscribe %v sucess\n", topic)
	glog.Infof("Subscribe %v sucess\n", topic)
	return nil
}

func (this *Consumer) Start() error {
	//启动
	err := (*this.consumer).Start()
	if err != nil {
		glog.Error(err.Error())
		//os.Exit(-1)
		return err
	}
	return nil
}
