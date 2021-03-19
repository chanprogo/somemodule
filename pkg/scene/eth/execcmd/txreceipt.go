package execcmd

import (
	"encoding/json"

	"log"
)

type ReceiptResult struct {
	ID      int     `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  Receipt `json:"result"`
}

type Receipt struct {
	BlockNumber string `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`

	TransactionHash  string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`

	LogsBloom string `json:"logsBloom"`
	Logs      []Log  `json:"logs"`

	FromAddress string `json:"from"`
	ToAddress   string `json:"to"`

	GasUsed string `json:"gasUsed"`
}

func TransactionReceipt(address string) Receipt {

	method := "eth_getTransactionReceipt"
	params := `["` + address + `"]`

	jsrp, err := CallGeth(method, params)
	if err != nil {
		log.Println("fail to get infomation")
		panic("fail to get infomation")
	}

	r := ReceiptResult{}
	err = json.Unmarshal([]byte(jsrp), &r)
	if err != nil {
		log.Println("chang json to Receipt failed", err.Error())
		panic(err)
	}

	return r.Result
}
