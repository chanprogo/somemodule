package ethrpcclient

import (
	"encoding/json"
)

type TxReceipt struct {
	BlockNumber string `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`

	TransactionIndex string `json:"transactionIndex"`
	TransactionHash  string `json:"transactionHash"`

	FromAddress string `json:"from"`
	ToAddress   string `json:"to"`

	GasUsed string `json:"gasUsed"`

	LogsBloom string `json:"logsBloom"`
	// Logs             []Log  `json:"logs"`

	Status string `json:"status"`
}

func (r *RPCClient) GetTxReceipt(hash string) (*TxReceipt, error) {
	rpcResp, err := r.doPost(r.Url, "eth_getTransactionReceipt", []string{hash})
	if err != nil {
		return nil, err
	}
	if rpcResp.Result != nil {
		var reply *TxReceipt
		err = json.Unmarshal(*rpcResp.Result, &reply)
		return reply, err
	}
	return nil, nil
}

func (r *TxReceipt) Confirmed() bool {
	return len(r.BlockHash) > 0
}

const receiptStatusSuccessful = "0x1"

// Use with previous method
func (r *TxReceipt) Successful() bool {
	if len(r.Status) > 0 {
		return r.Status == receiptStatusSuccessful
	}
	return true
}
