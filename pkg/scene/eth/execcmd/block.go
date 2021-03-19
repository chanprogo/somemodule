package execcmd

import (
	"encoding/json"
	"log"
	"strconv"
)

type BlcokResult struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  Block  `json:"result"`
}

type Block struct {
	Number     string `json:"number"`
	Timestamp  string `json:"timestamp"`
	ParentHash string `json:"parentHash"`
	Hash       string `json:"hash"`

	Nonce string `json:"nonce"`
	Size  string `json:"size"`
	Miner string `json:"miner"`

	Difficulty      string `json:"difficulty"`
	TotalDifficulty string `json:"totalDifficulty"`

	GasLimit string `json:"gasLimit"`
	GasUsed  string `json:"gasUsed"`

	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	BlockNumber      string `json:"blockNumber"`
	BlockHash        string `json:"blockHash"`
	TransactionIndex string `json:"transactionIndex"`
	TransactionHash  string `json:"hash"`

	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`

	Input string `json:"input"`

	Gas      string `json:"gas"`
	GasPrice string `json:"gasprice"`
	Nonce    string `json:"nonce"`
}

func GetBlockByNumber(number int64) Block {

	method := "eth_getBlockByNumber"
	params := `["` + "0x" + strconv.FormatInt(number, 16) + `",true]`

	jsrp, err := CallGeth(method, params)
	if err != nil {
		log.Println("fail to get infomation")
		panic(err)
	}

	r := BlcokResult{}
	err = json.Unmarshal([]byte(jsrp), &r)
	if err != nil {
		log.Println("change to Block failed", err.Error())
		panic(err)
	}

	return r.Result
}
