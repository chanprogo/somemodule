package main

import (
	"fmt"
	"time"

	"github.com/chanprogo/somemodule/pkg/tcpserver"
)

func main() {

	server := tcpserver.GetServer()

	

	sErr := server.Start(new(DefaultIoHandlerFactory))
	
	if sErr != nil {
		fmt.Println("Server Start err: " + sErr.Error())
		return
	}

	for {
		time.Sleep(time.Second * 20)
	}
}
