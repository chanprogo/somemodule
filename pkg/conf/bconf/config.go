package bconf

import (
	"strings"

	beeConfig "github.com/astaxie/beego/config"
	"github.com/gin-gonic/gin"
)

var Configer *config

type config struct {
	BeeConfiger beeConfig.Configer

	ApiConf   ApiConfig
	DbConf    DbConfig
	RedisConf RedisConfig

	KafkaConfig     KafkaConfig
	ZookeeperConfig ZookeeperConfig

	MQTTConf MQTTConfig
	ETHConf  ETHConf
}

func NewConfiger(filename string) {
	Configer = new(config)
	var err error
	Configer.BeeConfiger, err = beeConfig.NewConfig("ini", filename)
	if err != nil {
		panic("读取配置文件出错")
	}
	Configer.load()
}

func (c *config) load() {

	var err error

	c.ApiConf.HttpPort = c.BeeConfiger.String("system::http_port")
	c.ApiConf.RpcPort = c.BeeConfiger.String("system::rpc_port")
	c.ApiConf.RunMode = c.BeeConfiger.String("system::run_mode")
	if c.ApiConf.RunMode != gin.DebugMode && c.ApiConf.RunMode != gin.TestMode && c.ApiConf.RunMode != gin.ReleaseMode {
		panic("run_mode 配置错误")
	}
	c.ApiConf.ParamSecret = c.BeeConfiger.String("system::param_secret")

	if c.BeeConfiger.String("system::worker_id") != "" {
		c.ApiConf.WorkerID, err = c.BeeConfiger.Int64("system::worker_id")
		if err != nil {
			panic("读取system::worker_id配置出错")
		}
	}

	c.ApiConf.LogPath = c.BeeConfiger.String("log::path")
	c.ApiConf.LogLevel = c.BeeConfiger.String("log::level")
	if c.ApiConf.LogLevel != "debug" && c.ApiConf.LogLevel != "info" && c.ApiConf.LogLevel != "error" {
		panic("log_level配置错误")
	}

	if c.BeeConfiger.String("db::driver_name") != "" {
		c.DbConf.DriverName = c.BeeConfiger.String("db::driver_name")
		c.DbConf.Host = c.BeeConfiger.String("db::host")
		c.DbConf.Port = c.BeeConfiger.String("db::port")
		c.DbConf.User = c.BeeConfiger.String("db::user")
		c.DbConf.Password = c.BeeConfiger.String("db::password")
		c.DbConf.Database = c.BeeConfiger.String("db::database")
		c.DbConf.MaxOpen, err = c.BeeConfiger.Int("db::max_open")
		if err != nil {
			panic("读取db:max_open配置出错")
		}

		c.DbConf.MaxIdle, err = c.BeeConfiger.Int("db::max_idle")
		if err != nil {
			panic("读取db:max_idle配置出错")
		}
	}

	if c.BeeConfiger.String("redis::type") != "" {
		c.RedisConf.Type, err = c.BeeConfiger.Int("redis::type")
		if err != nil {
			panic("读取redis::type配置出错")
		}
		if c.RedisConf.Type != 1 && c.RedisConf.Type != 2 {
			panic("读取redis::type配置出错，只能为1、2")
		}

		switch c.RedisConf.Type {
		case 2: //集群
			c.RedisConf.Cluster.Addrs = c.BeeConfiger.Strings("redis::addrs")
		default: //单点
			c.RedisConf.Single.Host = c.BeeConfiger.String("redis::host")
			c.RedisConf.Single.Port = c.BeeConfiger.String("redis::port")
			c.RedisConf.Single.DB, err = c.BeeConfiger.Int("redis::db")
			if err != nil {
				panic("读取redis::db配置出错")
			}
		}

		c.RedisConf.Password = c.BeeConfiger.String("redis::password")

		if c.BeeConfiger.String("redis::pool_size") != "" {
			c.RedisConf.PoolSize, err = c.BeeConfiger.Int("redis::pool_size")
			if err != nil {
				panic("读取redis::pool_size配置出错")
			}
		}
	}

	if c.BeeConfiger.String("kafka::addresses") != "" {
		addresses := c.BeeConfiger.String("kafka::addresses")
		c.KafkaConfig.Addresses = strings.Split(addresses, ",")
	}

	if c.BeeConfiger.String("zookeeper::addresses") != "" {
		addresses := c.BeeConfiger.String("zookeeper::addresses")
		c.ZookeeperConfig.Addresses = strings.Split(addresses, ",")
	}

	// MQTT 配置
	c.MQTTConf.Broker = c.BeeConfiger.String("mqtt::broker")
	c.MQTTConf.ClientID = c.BeeConfiger.String("mqtt::client_id")

	// ETH 配置
	c.ETHConf.Connstr = c.BeeConfiger.String("common::connstr")
	c.ETHConf.ContractAddr = c.BeeConfiger.String("common::contractAddr")
	c.ETHConf.CallerAddr = c.BeeConfiger.String("common::callerAddr")
	c.ETHConf.CallerPass = c.BeeConfiger.String("common::callerPass")
	c.ETHConf.Keydir = c.BeeConfiger.String("common::keyDir")

}
