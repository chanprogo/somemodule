package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/protodatasvr"
	"github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/rpcclient"
	"github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/rpcserver"
	"github.com/chanprogo/somemodule/pkg/conf/bconf"
	"github.com/chanprogo/somemodule/pkg/log"
)

func main() {

	bconf.NewConfiger("./app.conf")

	log.NewLogger(bconf.Configer.ApiConf.LogPath, bconf.Configer.ApiConf.LogLevel)

	app.NewRPCSvr(bconf.Configer.ApiConf.RpcPort)

	protodatasvr.RegisterEmailServiceServer(app.RPCSvr, &rpcserver.SendEmailServer{})

	rpcclient.RpcClient()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-quitChan
	app.StopRPCSvr()
}
