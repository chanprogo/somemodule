package execcmd

import (
	"encoding/json"
	"log"
	"strconv"
)

type LogsResult struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  []Log  `json:"result"`
}

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

func Logs(fromBlock int64, toBlock int64, contractAddress string) []Log {

	method := "eth_getLogs"
	params := `[{"fromBlock":"0x` + strconv.FormatInt(fromBlock, 16) + `","toBlock":"0x` + strconv.FormatInt(toBlock, 16) +
		`","address":"` + contractAddress +
		`", "topics": ["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"]}]`

	jsrp, err := CallGeth(method, params)
	if err != nil {
		log.Println("fail to get infomation")
		panic("fail to get infomation")
	}

	r := LogsResult{}
	err = json.Unmarshal([]byte(jsrp), &r)
	if err != nil {
		log.Println("change to Block failed", err.Error())
		panic(err)
	}

	return r.Result
}
