package tcpserver

import "net"

type IoConn interface {
	Start(IoHandler, net.Conn, uint32)
	Close()

	Read(uint32)
	Write([]byte)

	RemoteName() string
	RemoteAddr() string

	OnClose(**string)
}
