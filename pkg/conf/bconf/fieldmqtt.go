package bconf

type MQTTConfig struct {
	Broker   string // Broker 地址，例如 tcp://127.0.0.1:1883
	ClientID string
}
