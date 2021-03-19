package ethrpcclient

import (
	"encoding/json"
	"strconv"
)

type Log struct {
	Address string `json:"address"`

	Removed bool `json:"removed"`

	BlockNumber string `json:"blockNumber"`

	TransactionIndex string `json:"transactionIndex"`
	TransactionHash  string `json:"transactionHash"`

	LogIndex string `json:"logIndex"`

	Topics []string `json:"topics"`
	Data   string   `json:"data"`
}

func (r *RPCClient) GetLogs(fromBlock int64, toBlock int64, contractAddress string) ([]Log, error) {

	strSlice := []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}
	params := map[string]interface{}{
		"fromBlock": "0x" + strconv.FormatInt(fromBlock, 16),
		"toBlock":   "0x" + strconv.FormatInt(toBlock, 16),
		"address":   contractAddress,
		"topics":    strSlice,
	}

	rpcResp, err := r.doPost(r.Url, "eth_getLogs", []interface{}{params})
	if err != nil {
		return nil, err
	}

	var reply []Log
	err = json.Unmarshal(*rpcResp.Result, &reply)
	return reply, err
}
