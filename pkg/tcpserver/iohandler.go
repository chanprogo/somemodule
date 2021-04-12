package tcpserver

import "net"

type IoHandler interface {
	OnClosed()
	OnError(error)
	OnReadFinished(**string, net.Conn, []byte) (bool, int)
}

type IoHandlerFactory interface {
	CreateIoHandler() IoHandler
}