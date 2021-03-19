package mqtt

import (
	"time"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

type SubscribeType struct {
	Topic      string
	Qos        byte
	Callback   gomqtt.MessageHandler
	RetryTimes int // 表示订阅失败后的重试次数，如果为0，则表示一直重试下去。
}

var subscribers = make([]SubscribeType, 0)

// 注册订阅消息
func Subscribe(item SubscribeType) {
	subscribers = append(subscribers, item)
}

func subscribe(item SubscribeType) {
	times := 0
	for {
		token, err := subscribeItem(item)
		if err != nil {
			if item.RetryTimes == 0 || times < item.RetryTimes {
				times++
				time.Sleep(3 * time.Second)
				continue
			} else {
				panic(err)
			}
		}
		if token.Wait() && token.Error() != nil {
			if item.RetryTimes == 0 || times < item.RetryTimes {
				times++
				time.Sleep(3 * time.Second)
				continue
			} else {
				panic(token.Error())
			}
		}
		break
	}
}

func subscribeItem(item SubscribeType) (token gomqtt.Token, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		return
	}()
	token = client.Subscribe(item.Topic, item.Qos, item.Callback)
	return
}

// 订阅主题
// if token := client.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
// 	fmt.Println(token.Error())
// 	os.Exit(1)
// }

// 取消订阅
// if token := client.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
// 	os.Exit(1)
// }
