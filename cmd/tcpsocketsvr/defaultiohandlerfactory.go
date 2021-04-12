package main

import "github.com/chanprogo/somemodule/pkg/tcpserver"

type DefaultIoHandlerFactory struct {
}

func (defaultFactory *DefaultIoHandlerFactory) CreateIoHandler() tcpserver.IoHandler {
	return new(DefaultIoHandler)
}
