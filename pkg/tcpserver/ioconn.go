package tcpserver

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type MyTCPConn struct {
	conn net.Conn

	readedData    []byte
	readedDataLen int

	handler IoHandler
}

func (oTConn *MyTCPConn) Start(handler IoHandler, conn net.Conn, maxSendQueue uint32) {

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
func (oTConn *MyTCPConn) Close() {
}

func (oTConn *MyTCPConn) Read(bytes uint32) {
}

func (oTConn *MyTCPConn) Write(data []byte) {
}

// RemoteName ...
func (oTConn *MyTCPConn) RemoteName() string {
	return oTConn.conn.RemoteAddr().Network()
}

// RemoteAddr ...
func (oTConn *MyTCPConn) RemoteAddr() string {
	return oTConn.conn.RemoteAddr().String()
}

// OnClose ...
func (oTConn *MyTCPConn) OnClose(myKey **string) {
	fmt.Println("ObdTcpConn OnClose begin. addr:" + oTConn.conn.RemoteAddr().String())

	GetServer().OnIoDisCon(*myKey)
	oTConn.conn.Close()
	oTConn.handler.OnClosed()

	fmt.Println("ObdTcpConn OnClose finish..")
}
