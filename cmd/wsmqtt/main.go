package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chanprogo/somemodule/pkg/log"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/internal/wsmqttpkg/router"
	"github.com/chanprogo/somemodule/pkg/conf/bconf"
)

func main() {
	bconf.NewConfiger("./app.conf")

	log.NewLogger(bconf.Configer.ApiConf.LogPath, bconf.Configer.ApiConf.LogLevel)

	app.NewAPISvr(bconf.Configer.ApiConf.RunMode, bconf.Configer.ApiConf.HttpPort, 60*time.Second, 60*time.Second)
	router.InitRouter(app.GinEngine)
	app.RunAPISvr()

	// services.SubscribeData()
	// mqtt.Start(bconf.Configer.MQTTConf.Broker, bconf.Configer.MQTTConf.ClientID)

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-quitChan
	app.StopAPISvr()
}
