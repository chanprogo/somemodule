package mqtt

import (
	"fmt"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

// 跟 MQTT 服务器连接的实例
var client gomqtt.Client

func Start(broker, clientID string) {

	// gomqtt.DEBUG = log.New(os.Stdout, "", 0)
	// gomqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := gomqtt.NewClientOptions() // 获取 MQTT 连接配置项

	// opts.SetKeepAlive(60 * time.Second)
	opts.SetAutoReconnect(true)
	opts.AddBroker(broker)
	if len(clientID) > 0 {
		opts.SetClientID(clientID)
	}

	opts.SetOnConnectHandler(onConnectHandler(opts.OnConnect))
	opts.SetConnectionLostHandler(onConnectionLostHandler(opts.OnConnectionLost))

	// 设置消息回调处理函数
	var f gomqtt.MessageHandler = func(client gomqtt.Client, msg gomqtt.Message) {
		fmt.Printf("fmt - SetDefaultPublishHandler: TOPIC: %s\n", msg.Topic())
		fmt.Printf("fmt - SetDefaultPublishHandler: MSG: %s\n", msg.Payload())
	}
	opts.SetDefaultPublishHandler(f)
	// opts.SetPingTimeout(1 * time.Second)

	client = gomqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic("初始化失败")
	}

}

// 连接上服务器的操作
func onConnectHandler(handler gomqtt.OnConnectHandler) gomqtt.OnConnectHandler {
	return func(c gomqtt.Client) {

		for _, item := range subscribers {
			subscribe(item)
		}
		// handler(c)
	}
}

// 丢失连接的操作(自动重连)
func onConnectionLostHandler(handler gomqtt.ConnectionLostHandler) gomqtt.ConnectionLostHandler {
	return func(c gomqtt.Client, e error) {
		// handler(c, e)
	}
}
