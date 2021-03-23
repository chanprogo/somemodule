package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/router"
	"github.com/chanprogo/somemodule/pkg/conf/iconf"
	"github.com/chanprogo/somemodule/pkg/log/logging"
	"github.com/chanprogo/somemodule/pkg/module/database/gormmodel"
	"github.com/chanprogo/somemodule/pkg/module/redis/gredis"
)

func main() {
	iconf.Setup()

	logging.Setup(iconf.AppSetting.RuntimeRootPath, iconf.AppSetting.LogSavePath, iconf.AppSetting.LogSaveName,
		iconf.AppSetting.TimeFormat, iconf.AppSetting.LogFileExt)

	gormmodel.Setup(iconf.DatabaseSetting.Type, iconf.DatabaseSetting.User, iconf.DatabaseSetting.Password,
		iconf.DatabaseSetting.Host, iconf.DatabaseSetting.Name, iconf.DatabaseSetting.TablePrefix)

	gredis.Setup(iconf.RedisSetting.MaxIdle, iconf.RedisSetting.MaxActive, iconf.RedisSetting.IdleTimeout,
		iconf.RedisSetting.Host, iconf.RedisSetting.Password)

	app.NewAPISvr(iconf.ServerSetting.RunMode, iconf.ServerSetting.HttpPort, iconf.ServerSetting.ReadTimeout, iconf.ServerSetting.WriteTimeout)
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
