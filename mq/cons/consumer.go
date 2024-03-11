package cons

import (
	"context"
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/xxl6097/go-glog/glog"
)

type OnReceiver func(msg *primitive.MessageExt)

type Consumer struct {
	pushConsumer *rocketmq.PushConsumer
}

func New() *Consumer {
	c := &Consumer{
		pushConsumer: nil,
	}
	return c
}

func (this *Consumer) NewCustomConsumer(opts ...consumer.Option) error {
	c, err := rocketmq.NewPushConsumer(opts...)
	if err != nil {
		glog.Error("NewPushConsumer失败：", err)
		return err
	}
	this.pushConsumer = &c
	return nil
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
	this.pushConsumer = &c
	return nil
}

func (this *Consumer) Shutdown() error {
	if this.pushConsumer != nil {
		//关闭连接
		err := (*this.pushConsumer).Shutdown()
		if err != nil {
			glog.Errorf("shutdown pushConsumer error: %s\n", err.Error())
			return err
		}
	} else {
		return errors.New("pushConsumer not new,please call Connect() function")
	}
	return nil
}

func (this *Consumer) Subscribe(topic string, _receiver OnReceiver) error {
	if this.pushConsumer == nil {
		return errors.New("pushConsumer is nil")
	}
	if err := (*this.pushConsumer).Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
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
	err := (*this.pushConsumer).Start()
	if err != nil {
		glog.Error(err.Error())
		//os.Exit(-1)
		return err
	}
	return nil
}
