package tcpsocket

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// Server ...
type Server struct {
	ioHandlerFactory  IoHandlerFactory
	stopListenChannel chan bool
}

var obdServer *Server

func GetServer() *Server {
	if obdServer == nil {
		obdServer = new(Server)
	}
	return obdServer
}

var connMapMutex sync.Mutex
var clientConnMap map[string]net.Conn

// Start ...
func (thisServer *Server) Start(factory IoHandlerFactory) error {

	thisServer.ioHandlerFactory = factory
	clientConnMap = make(map[string]net.Conn)

	errorChannel := make(chan error)
	defer close(errorChannel)
	go thisServer.startTCP(errorChannel)

	err := <-errorChannel
	if err == nil {
		thisServer.stopListenChannel = make(chan bool)
	} else {
		fmt.Println(err.Error())
	}
	return err
}

// Stop ...
func (thisServer *Server) Stop() {
	fmt.Println("Server stop")
	thisServer.stopListenChannel <- true
	// 清理客户端
	//thisServer.closeAllConn()
}

// func (thisServer *Server) closeAllConn() {
// }

// OnIoDisCon ...
func (thisServer *Server) OnIoDisCon(myKey *string) {
	fmt.Println("[invalid]: remove one conn from map..")
	if myKey != nil {
		{
			connMapMutex.Lock()
			delete(clientConnMap, *myKey)
			connMapMutex.Unlock()
		}
	}
}

func (thisServer *Server) startTCP(errorChannel chan error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:11987")
	if nil != err {
		fmt.Println("TCP address error.")
		errorChannel <- err
		return
	}
	fmt.Println("Resolved TCP addr: " + tcpAddr.String())

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if nil != err {
		fmt.Println("TCP error.")
		errorChannel <- err
		return
	}

	fmt.Println("Start listening")

	errorChannel <- nil

	for {
		select {

		case <-thisServer.stopListenChannel:
			break

		default:
			err := listener.SetDeadline(time.Now().Add(time.Second))
			if err != nil {
				fmt.Println("set deadline fail")
			}

			conn, err := listener.AcceptTCP()
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
					fmt.Print(".")
					continue
				}

				fmt.Println(err.Error() + ", client conn count:" + strconv.Itoa(len(clientConnMap)))

				time.Sleep(1 * time.Second)
				continue
			}

			fmt.Println("Accept a connection, addr:" + conn.RemoteAddr().String())

			thisServer.onConn(conn)

		} //End of 'select'
	} //End of 'for'
}

func (thisServer *Server) onConn(conn net.Conn) {
	ioConn := new(ObdTCPConn)
	ioConn.Start(thisServer.ioHandlerFactory.CreateIoHandler(), conn, 10000)
}

// SendRespMsg ...
func SendRespMsg(obdsnOm string, dataMem []byte) int {

	var conn net.Conn
	var ok bool

	{
		connMapMutex.Lock()
		if conn, ok = clientConnMap[obdsnOm]; ok {
			connMapMutex.Unlock()
		} else {
			connMapMutex.Unlock()
			return 1
		}
	}

	if conn == nil {
		fmt.Println("remote is offline..")
		return 2
	}

	if err := conn.SetWriteDeadline(time.Now().Add(time.Second * 5)); err != nil {
		fmt.Println("set write timeout fail")
	}

	sendlen, err := conn.Write(dataMem)
	if sendlen == 0 || err != nil {
		var strErr string
		if nil == err {
			strErr = ""
		} else {
			strErr = err.Error()
		}
		fmt.Println("send data fail, sendlen:" + strconv.Itoa(sendlen) + ", err:" + strErr)
		conn.Close()
		return 3
	}

	fmt.Println("send out msg, sendlen:" + strconv.Itoa(sendlen))
	return 0
}
