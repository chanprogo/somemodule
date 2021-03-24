package services

import (
	"fmt"

	"github.com/chanprogo/somemodule/pkg/module/mqtt"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
	// "gopkg.in/redis.v3"
)

func SubscribeData() {

	// 注册需要订阅的消息，连接后会自动订阅
	mqtt.Subscribe(mqtt.SubscribeType{
		Topic:      "topic_event",
		Qos:        byte(2),
		Callback:   TestMessageHandler,
		RetryTimes: 0,
	})

}

// TestMessageHandler 函数定义格式如下
func TestMessageHandler(client gomqtt.Client, msg gomqtt.Message) {
	fmt.Printf("fmt - Subscribe - TOPIC: %s  MSG: %s", msg.Topic(), msg.Payload())
}
