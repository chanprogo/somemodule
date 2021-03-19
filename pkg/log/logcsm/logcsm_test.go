package logcsm

import (
	"fmt"
	"testing"
	"time"
)

func TestLogCSM(t *testing.T) {

	mylog := GetLog()
	err := mylog.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		GetLog().Stop()
	}()

	LogDebug(">>>> >>> Start running! >>>> >>>")

	// for {
	// 	now := time.Now()
	// 	next := now.Add(time.Minute * 1)
	// 	t := time.NewTimer(next.Sub(now))
	// 	<-t.C
	// }
	time.Sleep(time.Duration(1) * time.Second)
}
