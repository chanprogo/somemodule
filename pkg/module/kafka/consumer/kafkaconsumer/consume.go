package kafkaconsumer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"github.com/wvanbergen/kazoo-go"
)

// RecvMessage ...
func RecvMessage(messageValue []byte) {

	anlmsg := &MyInfo{}

	if err := json.Unmarshal(messageValue, anlmsg); err != nil {
		fmt.Println("Unmarshal faild. ", err.Error())
		return
	}

	MyCacheData <- *anlmsg
	fmt.Println(" - ")
}

func consume() {
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	var zookeeperNodes []string
	zookeeperNodes, config.Zookeeper.Chroot = kazoo.ParseConnectionString("120.79.56.54:2181,112.74.173.118:2181,47.106.188.16:2181")

	//kafkaTopics := strings.Split("topic1,topic2", ",")
	kafkaTopics := []string{"topic001"}

	consumer, err := consumergroup.JoinConsumerGroup("my-consumer-group-001", kafkaTopics, zookeeperNodes, config)
	if err != nil {
		log.Fatalln(err)
	}

	for message := range consumer.Messages() {
		RecvMessage(message.Value)
		consumer.CommitUpto(message)
	}

	defer func() {
		if errClose := consumer.Close(); errClose != nil {
		}
		close(MainChannel)
	}()
}
