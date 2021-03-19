package app

import (
	"fmt"

	"strconv"
	"sync"
	"time"

	"github.com/chanprogo/somemodule/pkg/util/waitgrouputil"
)

var IsQuit bool

type CronObj struct {
	lastTime int64
	lock     sync.RWMutex
}

func (cron *CronObj) Init(intervalS string, f func()) {

	// cron.lastTime = time.Now().UnixNano()

	interval, err := strconv.ParseInt(intervalS, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("解析 intervalS 错误:%s", intervalS))
	}
	interval *= int64(time.Second)

	waitgrouputil.WaitGroup.WrapGoroutine(func(...interface{}) {
		for {
			if IsQuit {
				break
			}
			nowTime := time.Now().UnixNano()
			if cron.lastTime+interval < nowTime {
				cron.lastTime = nowTime
				f()
			}
			time.Sleep(time.Millisecond * 100)
		}
	})
}
