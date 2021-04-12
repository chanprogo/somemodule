package tcpserver

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var ConnMapMutex sync.Mutex
var ClientConnMap map[string]net.Conn

// Server ...
type Server struct {
	ioHandlerFactory  IoHandlerFactory
	stopListenChannel chan bool
}

var myServer *Server

func GetServer() *Server {
	if myServer == nil {
		myServer = new(Server)
	}
	return myServer
}

func (thisServer *Server) Start(factory IoHandlerFactory) error {

	thisServer.ioHandlerFactory = factory

	ClientConnMap = make(map[string]net.Conn)

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
				fmt.Println(err.Error() + ", client conn count:" + strconv.Itoa(len(ClientConnMap)))
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Println("Accept a connection, addr:" + conn.RemoteAddr().String())
			thisServer.onConn(conn)
		} //End of 'select'
	} //End of 'for'
}

func (thisServer *Server) onConn(conn net.Conn) {
	ioConn := new(MyTCPConn)
	ioConn.Start(thisServer.ioHandlerFactory.CreateIoHandler(), conn, 10000)
}

func SendRespMsg(obdsnOm string, dataMem []byte) int {
	var conn net.Conn
	var ok bool

	{
		ConnMapMutex.Lock()
		if conn, ok = ClientConnMap[obdsnOm]; ok {
			ConnMapMutex.Unlock()
		} else {
			ConnMapMutex.Unlock()
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

func (thisServer *Server) OnIoDisCon(myKey *string) {
	fmt.Println("[invalid]: remove one conn from map..")
	if myKey != nil {
		{
			ConnMapMutex.Lock()
			delete(ClientConnMap, *myKey)
			ConnMapMutex.Unlock()
		}
	}
}

func (thisServer *Server) Stop() {
	fmt.Println("Server stop")
	thisServer.stopListenChannel <- true
	//thisServer.closeAllConn() // 清理客户端
}

// func (thisServer *Server) closeAllConn() {
// }
