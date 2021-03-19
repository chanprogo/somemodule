package execcmd

import (
	"encoding/json"
	"log"
	"strconv"
)

type BlockNumberResult struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
}

func BlockNumber() int64 {

	method := "eth_blockNumber"
	params := `[]`

	jsrp, err := CallGeth(method, params)
	if err != nil {
		log.Println("fail to get infomation")
		panic("fail to get infomation")
	}

	r := BlockNumberResult{}

	err = json.Unmarshal([]byte(jsrp), &r)
	if err != nil {
		log.Println("change to blockNumber failed", err.Error())
		panic(err)
	}

	number, err := strconv.ParseInt(r.Result, 0, 64)
	if err != nil {
		log.Println("string to int failed", err.Error())
		panic(err.Error())
	}

	return number
}
