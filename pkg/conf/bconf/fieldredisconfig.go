package bconf

type RedisConfig struct {
	Type     int // 类型，1-单点  2-集群
	Password string
	PoolSize int
	Single   struct { //单点配置
		Host string
		Port string
		DB   int
	}
	Cluster struct { //集群配置
		Addrs []string
	}
}
