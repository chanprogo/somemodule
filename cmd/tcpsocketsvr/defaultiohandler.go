package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/chanprogo/somemodule/pkg/tcpserver"
)

type DefaultIoHandler struct {
	// conn IoConn
}

func (dH *DefaultIoHandler) OnClosed() {
	//	fmt.Println("io close, addr:%s, name:%s", dH.conn.RemoteAddr(), dH.conn.RemoteName())
}

func (dH *DefaultIoHandler) OnError(err error) {
	//	fmt.Println("io error form, addr:%s, name:%s", dH.conn.RemoteAddr(), dH.conn.RemoteName())
	//	dH.conn.Close()
}

// 返回值：    bool: 是否正确处理        int: 消息处理 用掉的数据长度
func (dH *DefaultIoHandler) OnReadFinished(myKey **string, conn net.Conn, data []byte) (bool, int) {

	msgLen := len(data)

	var cpdata = make([]byte, msgLen)
	copy(cpdata, data[:msgLen])

	if *myKey == nil {
		temp := strconv.FormatInt(time.Now().Unix(), 10)
		*myKey = &temp
		{
			tcpserver.ConnMapMutex.Lock()
			tcpserver.ClientConnMap[**myKey] = conn
			tcpserver.ConnMapMutex.Unlock()
		}
	}

	rsp := []byte{1, 2, 3}
	if rsp != nil {

		if err := conn.SetWriteDeadline(time.Now().Add(time.Second * 5)); err != nil {
			fmt.Println("set write timeout fail")
		}

		sendlen, err := conn.Write(rsp)

		if sendlen == 0 || err != nil {
			strErr := ""
			if err != nil {
				strErr = err.Error()
			}
			fmt.Println("send data fail, sendlen:" + strconv.Itoa(sendlen) + ", err:" + strErr)
			return false, msgLen
		}

		fmt.Println("send out msg, sendlen:" + strconv.Itoa(sendlen))
	}

	return true, msgLen

}
