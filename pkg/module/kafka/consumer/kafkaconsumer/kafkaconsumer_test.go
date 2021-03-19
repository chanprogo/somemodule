package kafkaconsumer

import "testing"

func TestKafkaConsumer(t *testing.T) {

	MyCacheData = make(chan MyInfo, 10000)
	MainChannel = make(chan MyInfo, 1000)
	go GoCacheData(MainChannel, 5)

	go goAnalysis()

	consume()

}
