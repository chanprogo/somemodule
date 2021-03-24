package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/pkg/conf/bconf"
	"github.com/chanprogo/somemodule/pkg/log"
	// "github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/rpcclient"
)

func main() {

	bconf.NewConfiger("./app.conf")

	log.NewLogger(bconf.Configer.ApiConf.LogPath, bconf.Configer.ApiConf.LogLevel)

	app.NewApp()
	app.NewRPCSvr(bconf.Configer.ApiConf.RpcPort)

	// rpcclient.RpcClient() // 测试 rpc 服务

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-quitChan
	app.StopRPCSvr()
}
