package kafkaproducer

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// MyInfo ...
type MyInfo struct {
	MyID     int    `json:"myID"`
	MyName   string `json:"myName"`
	MyTime   string `json:"myTime"`
	MyEnable bool   `json:"myEnable"`
}

// PushStrings ...
var PushStrings chan string

func TestKafkaProducer(t *testing.T) {

	PushStrings = make(chan string, 1000)

	for i := 400; i < 500; i++ {
		var temp MyInfo
		j := i
		temp.MyID = j
		temp.MyName = "chan" + strconv.Itoa(j)
		temp.MyTime = strconv.FormatInt(time.Now().Unix(), 10)
		temp.MyEnable = true

		tjson, err := json.Marshal(temp)
		if err != nil {
			panic(err)
		}
		PushStrings <- string(tjson)
		fmt.Printf("%v - ", j)
	}
	fmt.Println()

	go PushStringDataToKafka(PushStrings, "topic001", "120.79.56.54:9092,112.74.173.118:9092,47.106.188.16:9092")

	quit := make(chan bool)
	<-quit
}
