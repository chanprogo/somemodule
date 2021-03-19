package ethrpcclient

import (
	"encoding/json"
	"math/big"
	"strconv"
	"strings"
)

func (r *RPCClient) GetPeerCount() (int64, error) {
	rpcResp, err := r.doPost(r.Url, "net_peerCount", nil)
	if err != nil {
		return 0, err
	}
	var reply string
	err = json.Unmarshal(*rpcResp.Result, &reply)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strings.Replace(reply, "0x", "", -1), 16, 64)
}

func (r *RPCClient) GetBalance(address string, num ...int64) (*big.Int, error) {

	params := []string{address, "latest"}
	if len(num) > 0 {
		params = []string{address, "0x" + strconv.FormatInt(num[0], 16)}
	}

	rpcResp, err := r.doPost(r.Url, "eth_getBalance", params)
	if err != nil {
		return nil, err
	}
	var reply string
	err = json.Unmarshal(*rpcResp.Result, &reply)
	if err != nil {
		return nil, err
	}

	bignumber := big.NewInt(0)
	bignumber.SetString(reply, 0)

	return bignumber, err
}

// curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' 8.129.172.186:8545
func (r *RPCClient) GetBlockNumber() (int64, error) {
	rpcResp, err := r.doPost(r.Url, "eth_blockNumber", nil)
	if err != nil {
		return 0, err
	}
	var reply string
	err = json.Unmarshal(*rpcResp.Result, &reply)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strings.Replace(reply, "0x", "", -1), 16, 64)
}
