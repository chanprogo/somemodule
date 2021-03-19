package kafkaconsumer

import (
	"fmt"
	"time"
)

// MainChannel ...
var MainChannel chan MyInfo

func goAnalysis() {

	timeout := make(chan bool, 1)
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(2))
			timeout <- true
		}
	}()

	var num int

	for {
		var TotalMsg MyInfo
		select {

		case TotalMsg = <-MainChannel:
			num++
			fmt.Printf("%v goAnalysis --------- %+v \n", num, TotalMsg)
		case <-timeout:
			fmt.Println("-------")

		}
	}
}
