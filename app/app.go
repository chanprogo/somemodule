package app

import (
	"github.com/chanprogo/somemodule/pkg/conf/bconf"
	"github.com/chanprogo/somemodule/pkg/module/redis/cache"

	"github.com/chanprogo/somemodule/pkg/log"
	"github.com/chanprogo/somemodule/pkg/module/database/model"
	"github.com/chanprogo/somemodule/pkg/util/encryption/aesutil/cbc"
	"github.com/chanprogo/somemodule/pkg/util/id"
)

func NewApp() {

	log.NewLogger(bconf.Configer.ApiConf.LogPath, bconf.Configer.ApiConf.LogLevel)

	if bconf.Configer.BeeConfiger.String("system::worker_id") != "" {
		id.NewIdWorker(bconf.Configer.ApiConf.WorkerID)
	}

	if bconf.Configer.BeeConfiger.String("redis::host") != "" {
		redisPassword := bconf.Configer.RedisConf.Password
		encrypt := bconf.Configer.BeeConfiger.DefaultBool("system::pswd_encrypt", false)
		if encrypt {
			pswd := cbc.AesDecryptByKey("yashikejigitcrpy", redisPassword)
			redisPassword = string(pswd)
		}
		cache.NewRedisClient(redisPassword, bconf.Configer.RedisConf.Type, bconf.Configer.RedisConf.Cluster.Addrs,
			bconf.Configer.RedisConf.Single.Host, bconf.Configer.RedisConf.Single.Port,
			bconf.Configer.RedisConf.Single.DB, bconf.Configer.RedisConf.PoolSize)
	}

	password := bconf.Configer.DbConf.Password
	encrypt := bconf.Configer.BeeConfiger.DefaultBool("system::pswd_encrypt", false)
	if encrypt {
		pswd := cbc.AesDecryptByKey("yashikejigitcrpy", password)
		password = string(pswd)
	}

	switch bconf.Configer.DbConf.DriverName {
	case "mysql":
		model.NewMysqlOrm(bconf.Configer.DbConf.User, password, bconf.Configer.DbConf.Host, bconf.Configer.DbConf.Port,
			bconf.Configer.DbConf.Database, bconf.Configer.DbConf.MaxOpen, bconf.Configer.DbConf.MaxIdle)
	case "postgres":
		model.NewPostgresOrm(bconf.Configer.DbConf.Host, bconf.Configer.DbConf.Port, bconf.Configer.DbConf.User, password,
			bconf.Configer.DbConf.Database, bconf.Configer.DbConf.MaxOpen, bconf.Configer.DbConf.MaxIdle)
	}

}
