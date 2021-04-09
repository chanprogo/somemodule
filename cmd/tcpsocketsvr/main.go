package main

import (
	"fmt"
	"time"

	"github.com/chanprogo/somemodule/pkg/iohandler"
)

func main() {

	server := GetServer()

	ioHandlerFactory := new(iohandler.DefaultIoHandlerFactory)

	sErr := server.Start(ioHandlerFactory)
	if sErr != nil {
		fmt.Println("Server Start err: " + sErr.Error())
		return
	}

	for {
		time.Sleep(time.Second * 20)
	}
}
