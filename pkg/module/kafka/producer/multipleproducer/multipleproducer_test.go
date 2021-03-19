package multipleproducer

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

func TestMultipleProducer(t *testing.T) {

	anaProducers := make([]sarama.AsyncProducer, 10) // 一个数组

	ProducerAnalisysChans = make([]chan *MyInfo, 10) // 一个数组，其中每个元素都是通道

	var i uint32
	for i = 0; i < 10; i++ {
		ProducerAnalisysChans[i] = make(chan *MyInfo, 10000)

		var pErr error
		anaProducers[i], pErr = sarama.NewAsyncProducer(strings.Split("39.104.81.59:9092,39.104.153.163:9092,39.104.145.49:9092", ","), nil)
		if pErr != nil {
			fmt.Println("producer_ana[" + strconv.Itoa(int(i)) + "]" + "NewAsyncProducer err: " + pErr.Error())
			panic(pErr)
		}
	}

	for i = 0; i < 10; i++ {
		go InputAnalisysKafka("analysis-msg-topic", anaProducers[i], i)
	}

	defer func() {
		for i = 0; i < 10; i++ {
			if err := anaProducers[i].Close(); err != nil {
				fmt.Println("producer_ana[", i, "] Error closing the producer", err)
			}

			close(ProducerAnalisysChans[i])

		}

	}()

	for {
		time.Sleep(time.Second * 20)
	}

}
