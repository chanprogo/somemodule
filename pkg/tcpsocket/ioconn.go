package tcpsocket

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// IoConn ...
type IoConn interface {
	Start(IoHandler, net.Conn, uint32)
	Close()
	Read(uint32)
	Write([]byte)
	RemoteName() string
	RemoteAddr() string
	OnClose(**string)
}

// ObdTCPConn ...
type ObdTCPConn struct {
	conn          net.Conn
	readedData    []byte
	readedDataLen int
	handler       IoHandler
}

// Start ...
func (oTConn *ObdTCPConn) Start(handler IoHandler, conn net.Conn, maxSendQueue uint32) {
	go func() {
		oTConn.conn = conn
		oTConn.handler = handler
		oTConn.readedData = make([]byte, 7000)
		oTConn.readedDataLen = 0

		var myKey *string

		var stopLoop bool
		for !stopLoop {
			fmt.Println("Read data begin. address:" + conn.RemoteAddr().String() + "\n")

			if err := oTConn.conn.SetReadDeadline(time.Now().Add(time.Minute * (time.Duration(5)))); err != nil {
				fmt.Println("set read timeout fail, addr:" + conn.RemoteAddr().String())
				stopLoop = true
			} else {

				readBytes, err := oTConn.conn.Read(oTConn.readedData[oTConn.readedDataLen:])
				if err != nil {
					fmt.Println("Read data fail. err:" + err.Error() + ", addr:" + conn.RemoteAddr().String())
					stopLoop = true
				} else if readBytes == 0 {
					fmt.Println("Read nothing.., addr:" + conn.RemoteAddr().String())
					stopLoop = true
				} else {
					fmt.Println("Receive data len:" + strconv.Itoa(readBytes))
					oTConn.readedDataLen += readBytes

					for oTConn.readedDataLen > 0 {

						suc, msgLen := oTConn.handler.OnReadFinished(&myKey, conn, oTConn.readedData[:oTConn.readedDataLen])
						if !suc {
							fmt.Println("OnReadFinished fail.")
							stopLoop = true
							break

						} else if msgLen > 0 {

							{
								leftLen := oTConn.readedDataLen - msgLen
								for i := 0; i < leftLen; i++ {
									oTConn.readedData[i] = oTConn.readedData[i+msgLen]
								}
							}

							oTConn.readedDataLen -= msgLen

						} else {
							fmt.Println("No complete msg left.")
							break
						}
					}
				}
			}
		} //'for !stopLoop'

		oTConn.OnClose(&myKey)

	}() //End of 'go func'
}

// Close ...
func (oTConn *ObdTCPConn) Close() {
}

func (oTConn *ObdTCPConn) Read(bytes uint32) {
}

func (oTConn *ObdTCPConn) Write(data []byte) {
}

// RemoteName ...
func (oTConn *ObdTCPConn) RemoteName() string {
	return oTConn.conn.RemoteAddr().Network()
}

// RemoteAddr ...
func (oTConn *ObdTCPConn) RemoteAddr() string {
	return oTConn.conn.RemoteAddr().String()
}

// OnClose ...
func (oTConn *ObdTCPConn) OnClose(myKey **string) {
	fmt.Println("ObdTcpConn OnClose begin. addr:" + oTConn.conn.RemoteAddr().String())

	GetServer().OnIoDisCon(*myKey)
	oTConn.conn.Close()
	oTConn.handler.OnClosed()

	fmt.Println("ObdTcpConn OnClose finish..")
}
