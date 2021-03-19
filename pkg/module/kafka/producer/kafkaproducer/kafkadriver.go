package kafkaproducer

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	jsoniter "github.com/json-iterator/go"

	"github.com/golang/glog"
)

// PushStringDataToKafka ...
func PushStringDataToKafka(c chan string, topic string, connect string) {

	producer, err := sarama.NewAsyncProducer(strings.Split(connect, ","), nil)
	if err != nil {
		fmt.Println("producer NewAsyncProducer err: ", err.Error())
		panic(err)
	}

	fmt.Println("PushStringDataToKafka begin...")
	count := 0
	for {
		fmt.Print(" -  ")
		dataSave := <-c

		count++
		fmt.Printf("count:%v \n", count)

		producerMessage := &sarama.ProducerMessage{}
		producerMessage.Topic = topic
		producerMessage.Value = sarama.ByteEncoder([]byte(dataSave))

		select {
		case producer.Input() <- producerMessage:
		case produceSendErr := <-producer.Errors():
			fmt.Println("Failed to produce message", produceSendErr)
		}
	}
}

var newProducer sarama.AsyncProducer

func InitAsynProducer(addrs []string) error {

	var err error

	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.ChannelBufferSize = 1 << 12 // config.ChannelBufferSize = 1 << 18
	config.Net.MaxOpenRequests = 32
	config.Producer.RequiredAcks = sarama.WaitForLocal // high level ack, Must ensure that data is not lost
	// config.Producer.Retry.Backoff = time.Duration(math.MaxInt32)
	config.Metadata.Full = true
	config.Producer.Retry.Max = 1 << 10

	newProducer, err = sarama.NewAsyncProducer(addrs, config)
	if err != nil {
		glog.Errorln(err)
		// panic(err)
		return err
	}

	go chkAsyncProducer(newProducer)
	return nil
}

func chkAsyncProducer(p sarama.AsyncProducer) {
	errors := p.Errors()
	success := p.Successes()
	for {
		select {
		case err := <-errors:
			if err != nil {
				glog.Error("producer errror", err)
			}
		case <-success:
		}
	}
}

func SendOrderToMw(v interface{}, groupID string) error {
	buf, err := jsoniter.Marshal(v)
	if err != nil {
		glog.Error(err)
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: fmt.Sprintf("market_group_%s", groupID),
		Value: sarama.ByteEncoder(buf),
	}
	newProducer.Input() <- msg
	return nil
}

func SendMakeVolToMw(v interface{}, groupID string) error {
	buf, err := jsoniter.Marshal(v)
	if err != nil {
		glog.Error(err)
		return err
	}
	msg := &sarama.ProducerMessage{
		Key:   sarama.ByteEncoder([]byte("make_volume")),
		Topic: fmt.Sprintf("make_vol_%s", groupID),
		Value: sarama.ByteEncoder(buf),
	}
	newProducer.Input() <- msg
	return nil
}
