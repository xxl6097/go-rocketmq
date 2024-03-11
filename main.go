package main

import (
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-rocketmq/mq"
	_ "github.com/xxl6097/go-rocketmq/mq/log"
	"os"
	"time"
)

var nameserver = []string{"10.6.14.3:9876"}
var groupname = "clink-tcp-roketmq"
var topicSubs = "clink-any-to-tcp-topic"
var topicPush = "clink-tcp-to-any-topic"

func main() {
	//export ROCKETMQ_GO_LOG_LEVEL=error
	//os.Setenv("mq.consoleAppender.enabled", "true")
	//golang.ResetLogger()
	// new simpleConsumer inst
	//log.LogDebug = true
	os.Setenv("ROCKETMQ_GO_LOG_LEVEL", "error")
	rokmq := mq.NewMQ()
	rokmq.NewRocketMQ([]consumer.Option{consumer.WithConsumerModel(consumer.BroadCasting)}, []producer.Option{})
	rokmq.InitRocketMQ(nameserver, groupname)
	//Subscribe必须再Start之前
	rokmq.Subscribe(topicSubs, func(msg *primitive.MessageExt) {
		glog.Info("-->recv:", string(msg.Body))
	})
	rokmq.Start()
	rokmq.SendSync(topicPush, "hello world "+time.Now().String())

	select {}

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//time.Wait(ctx)
}

func test() {
	// 创建RocketMQ消费者
}
