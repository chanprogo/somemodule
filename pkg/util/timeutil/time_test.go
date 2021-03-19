package timeutil

import (
	"fmt"
	"testing"
	"time"
)

func TestCalculateTimeDelay(t *testing.T) {
	delay, ok := CalculateTimeDelay("2019-10-24 17:17:00", time.Now().Format("2006-01-02 15:04:05"))
	if ok == nil && delay > 18 {
		fmt.Printf("CalculateTimeDelay: %+v\n", delay)
	}
}

func TestGetMonthStartAndEnd(t *testing.T) {
	result := GetMonthStartAndEnd("2020", "6")
	fmt.Print(result)
}

/*
	for {
		now := time.Now()
		next := now.Add(time.Second * 1)
		t := time.NewTimer(next.Sub(now))
		<-t.C

		fmt.Println("timer")
	}
*/
