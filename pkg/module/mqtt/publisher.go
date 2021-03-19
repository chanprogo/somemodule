package mqtt

// 通用发布消息接口
func Publish(topic string, payload interface{}, qos byte, retained bool) (err error) {

	// 发布消息
	// token := client.Publish("testtopic/1", 0, false, "Hello World")
	// token.Wait()

	token := client.Publish(topic, qos, retained, payload)

	if token.Wait() && token.Error() != nil {
		err = token.Error()
	}

	return
}
