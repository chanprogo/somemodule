package id

import "fmt"

type idWorker interface {
	Generate() int64
}

var IdWorker idWorker

func NewIdWorker(node int64) {
	if node < 0 || node > nodeMax {
		panic(fmt.Sprintf("snowflake 节点必须在 0-%d 之间", nodeMax))
	}
	snowflakeIns := &snowflake{
		timestamp: 0,
		node:      node,
		step:      0,
	}
	IdWorker = snowflakeIns
}
