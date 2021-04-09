package main

import (
	"fmt"
	"time"

	"github.com/chanprogo/somemodule/pkg/tcpsocket"
)

func main() {

	server := tcpsocket.GetServer()

	ioHandlerFactory := new(tcpsocket.DefaultIoHandlerFactory)

	sErr := server.Start(ioHandlerFactory)
	if sErr != nil {
		fmt.Println("Server Start err: " + sErr.Error())
		return
	}

	for {
		time.Sleep(time.Second * 20)
	}
}
