package ethrpcclient

import (
	"net/http"
	"sync"
	"time"
)

type RPCClient struct {
	sync.RWMutex

	Url         string
	Name        string
	sick        bool
	sickRate    int
	successRate int

	client *http.Client
}

func NewRPCClient(name, url, timeout string) *RPCClient {

	rpcClient := &RPCClient{Name: name, Url: url}

	timeoutIntv, err := time.ParseDuration(timeout)
	if err != nil {
		panic("util: Can't parse duration `" + timeout + "`: " + err.Error())
	}

	rpcClient.client = &http.Client{
		Timeout: timeoutIntv,
	}

	return rpcClient
}
