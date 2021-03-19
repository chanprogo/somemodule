package main

import (
	"fmt"
	"time"
)

func main() {

	server := GetServer()
	sErr := server.Start(new(DefaultIoHandlerFactory))
	if sErr != nil {
		fmt.Println("Server Start err: " + sErr.Error())
		return
	}

	for {
		time.Sleep(time.Second * 20)
	}
}
