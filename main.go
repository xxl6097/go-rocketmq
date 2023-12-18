package main

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/xxl6097/go-rocketmq/mq"
	"time"
)

var nameserver = "10.6.14.3:9876"
var groupname = "clink-tcp-roketmq"
var topicSubs = "clink-any-to-tcp-topic"
var topicPush = "clink-tcp-to-any-topic"

func main() {
	rokmq := mq.NewMQ()
	rokmq.InitRocketMQ(nameserver, groupname)
	//Subscribe必须再Start之前
	rokmq.Subscribe(topicSubs, func(msg *primitive.MessageExt) {
		fmt.Println("-->recv:", string(msg.Body))
	})
	rokmq.Start()
	rokmq.SendSync(topicPush, "hello world "+time.Now().String())
	time.Sleep(time.Second * 60)

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//time.Wait(ctx)
}
