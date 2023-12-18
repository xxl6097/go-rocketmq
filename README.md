## RocketMQ地址：
    http://10.6.14.3:12581/#/message
    name-server:10.6.14.3:9876
    tcp公共组件--->clink服务   topic:    clink-tcp-to-any-topic
    clink服务--->tcp公共组件   topic:    clink-any-to-tcp-topic

## 添加库
    go get -u gitee.com/xxl6097/go-rocketmq

## 添加RocketMQ依赖
    go get -u github.com/apache/rocketmq-client-go/v2
    go mod tidy

## RocketMQ Go版本代码
https://blog.csdn.net/the_shy_faker/article/details/128965426

## 协议规则

     /**
     * 第三方Tcp客户端唯一标识
     */
    private String clientId;

    /**
     * 任务唯一标识
     */
    private String taskId;

    /**
     * tcp原始16进制值
     */
    private String sourceData;

    /**
     * tcp内容
     */
    private Object bodyData;

    /**
     *  收到消息处理时判断的key，如果不填则为产品Id
     */
    private String matchKey;

    /**
     * CLife产品Id
     */
    private String productId;

    /**
     * 第三方唯一标识
     */
    private String sn;

    /**
     * 1、第三方推送的消息类型  UpCommandCodeEnum
     * 2、推送到第三方的消息类型  DownCommandCodeEnum
     */
    private String commandEnumCode;