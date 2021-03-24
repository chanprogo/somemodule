package bconf

type ApiConfig struct {
	HttpPort    int
	RpcPort     string
	RunMode     string
	LogPath     string
	LogLevel    string
	WorkerID    int64
	ParamSecret string
}
