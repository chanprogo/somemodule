package multipleproducer

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Shopify/sarama"
)

// MyInfo ...
type MyInfo struct {
	MyID     int    `json:"myID"`
	MyName   string `json:"myName"`
	MyTime   string `json:"myTime"`
	MyEnable bool   `json:"myEnable"`
}

// ProducerAnalisysChans 解析消息chan
var ProducerAnalisysChans []chan *MyInfo

// InputAnalisysKafka ...
func InputAnalisysKafka(topicName string, producer sarama.AsyncProducer, i uint32) {

	var count int
	for sendKafkaMassage := range ProducerAnalisysChans[i] {

		producerMessage := &sarama.ProducerMessage{}
		producerMessage.Topic = topicName
		count++
		producerMessage.Key = sarama.StringEncoder("ABCD999" + strconv.Itoa(count))

		encodemsg, errEncode := json.Marshal(sendKafkaMassage)
		if errEncode != nil {
			fmt.Println("Producer_ana[" + strconv.Itoa(int(i)) + "] Encode err: " + errEncode.Error())
		}
		producerMessage.Value = sarama.ByteEncoder(encodemsg)

		select {
		case producer.Input() <- producerMessage:
			fmt.Println("Producer_ana[" + strconv.Itoa(int(i)) + "] message send success")
		case err := <-producer.Errors():
			fmt.Println("Failed to producer_ana[" + strconv.Itoa(int(i)) + "] message " + err.Error())
		}
	}

}
