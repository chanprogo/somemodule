package ethrpcclient

import (
	"encoding/json"
	"fmt"
)

type GetBlockReply struct {
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

	Uncles []string `json:"uncles"`

	Transactions []Tx `json:"transactions"`

	// https://github.com/ethereum/EIPs/issues/95
	SealFields []string `json:"sealFields"`
}

type Tx struct {
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

func (r *RPCClient) getBlockBy(method string, params []interface{}) (*GetBlockReply, error) {
	rpcResp, err := r.doPost(r.Url, method, params)
	if err != nil {
		return nil, err
	}
	if rpcResp.Result != nil {
		var reply *GetBlockReply
		err = json.Unmarshal(*rpcResp.Result, &reply)
		return reply, err
	}
	return nil, nil
}

func (r *RPCClient) GetBlockByHash(hash string) (*GetBlockReply, error) {
	params := []interface{}{hash, true}
	return r.getBlockBy("eth_getBlockByHash", params)
}

func (r *RPCClient) GetBlockByHeight(height int64) (*GetBlockReply, error) {
	params := []interface{}{fmt.Sprintf("0x%x", height), true}
	return r.getBlockBy("eth_getBlockByNumber", params)
}

func (r *RPCClient) GetUncleByBlockNumberAndIndex(height int64, index int) (*GetBlockReply, error) {
	params := []interface{}{fmt.Sprintf("0x%x", height), fmt.Sprintf("0x%x", index)}
	return r.getBlockBy("eth_getUncleByBlockNumberAndIndex", params)
}

type GetBlockReplyPart struct {
	Number     string `json:"number"`
	Difficulty string `json:"difficulty"`
}

func (r *RPCClient) GetPendingBlock() (*GetBlockReplyPart, error) {
	rpcResp, err := r.doPost(r.Url, "eth_getBlockByNumber", []interface{}{"pending", false})
	if err != nil {
		return nil, err
	}
	if rpcResp.Result != nil {
		var reply *GetBlockReplyPart
		err = json.Unmarshal(*rpcResp.Result, &reply)
		return reply, err
	}
	return nil, nil
}
