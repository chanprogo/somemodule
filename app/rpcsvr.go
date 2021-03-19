package app

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var RPCSvr *grpc.Server
var RPCAddr string

func NewRPCSvr(rpcport string) {

	RPCAddr = ":" + rpcport
	RPCSvr = grpc.NewServer()

	listen, err := net.Listen("tcp", RPCAddr)
	if err != nil {
		panic("监听 rpc 端口失败，err: " + err.Error())
	}

	go func() {
		err := RPCSvr.Serve(listen)
		if err != nil {
			panic(fmt.Sprintf("启动 rpc 服务失败，%v", err))
		}
	}()
}

func StopRPCSvr() {
	RPCSvr.GracefulStop()
}
