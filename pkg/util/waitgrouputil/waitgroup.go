package waitgrouputil

import (
	"sync"
)

var WaitGroup waitGroupWrapper

var Quit bool

func init() {
	WaitGroup.waitGroup = new(sync.WaitGroup)
}

type waitGroupWrapper struct {
	waitGroup *sync.WaitGroup
}

// 同步执行
func (w waitGroupWrapper) Wrap(handler func(params ...interface{}), params ...interface{}) {
	w.waitGroup.Add(1)
	defer w.waitGroup.Done()
	defer func() {
		if err := recover(); err != nil {
			// log  err
		}
	}()
	handler(params...)
}

// 启动协程执行
func (w waitGroupWrapper) WrapGoroutine(handler func(params ...interface{}), params ...interface{}) {
	w.waitGroup.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// log err
			}
		}()
		defer w.waitGroup.Done()
		handler(params...)
	}()
}

var waitGroupDone = make(chan struct{})

func (w waitGroupWrapper) Wait() {
	go func() {
		w.waitGroup.Wait()
		waitGroupDone <- struct{}{}
	}()
}

func MainWait() {

	WaitGroup.Wait()

	select {
	case <-waitGroupDone:
		// case <-time.After(10 * time.Second):
		// log.Logger.Error("wait group 超时")
	}
}
