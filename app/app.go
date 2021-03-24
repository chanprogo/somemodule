package app

import (
	"github.com/chanprogo/somemodule/pkg/conf/bconf"
	"github.com/chanprogo/somemodule/pkg/module/database/model"
	"github.com/chanprogo/somemodule/pkg/module/redis/cache"

	"github.com/chanprogo/somemodule/pkg/util/encryption/aesutil/cbc"
	"github.com/chanprogo/somemodule/pkg/util/id"
)

func NewApp() {

	id.NewIdWorker(bconf.Configer.ApiConf.WorkerID)

	redisPassword := bconf.Configer.RedisConf.Password
	encrypt := bconf.Configer.BeeConfiger.DefaultBool("system::pswd_encrypt", false)
	if encrypt {
		pswd := cbc.AesDecryptByKey("abcdikejigitcraa", redisPassword)
		redisPassword = string(pswd)
	}
	cache.NewRedisClient(redisPassword, bconf.Configer.RedisConf.Type, bconf.Configer.RedisConf.Cluster.Addrs,
		bconf.Configer.RedisConf.Single.Host, bconf.Configer.RedisConf.Single.Port,
		bconf.Configer.RedisConf.Single.DB, bconf.Configer.RedisConf.PoolSize)

	password := bconf.Configer.DbConf.Password
	// encrypt := bconf.Configer.BeeConfiger.DefaultBool("system::pswd_encrypt", false)
	if encrypt {
		pswd := cbc.AesDecryptByKey("abcdikejigitcraa", password)
		password = string(pswd)
	}
	model.NewMysqlOrm(bconf.Configer.DbConf.User, password, bconf.Configer.DbConf.Host, bconf.Configer.DbConf.Port,
		bconf.Configer.DbConf.Database, bconf.Configer.DbConf.MaxOpen, bconf.Configer.DbConf.MaxIdle)

}
