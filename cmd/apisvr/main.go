package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/router"
	"github.com/chanprogo/somemodule/pkg/conf/iconf"
	"github.com/chanprogo/somemodule/pkg/log/logging"
)

func main() {
	iconf.Setup()
	logging.Setup(iconf.AppSetting.RuntimeRootPath, iconf.AppSetting.LogSavePath, iconf.AppSetting.LogSaveName,
		iconf.AppSetting.TimeFormat, iconf.AppSetting.LogFileExt)

	app.NewAPISvr(iconf.ServerSetting.RunMode, iconf.ServerSetting.HttpPort)
	router.InitRouter(app.GinEngine)
	app.RunAPISvr()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-quitChan
	app.StopAPISvr()
}
