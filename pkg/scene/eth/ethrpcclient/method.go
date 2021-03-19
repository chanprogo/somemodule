package ethrpcclient

import (
	"encoding/json"
	"errors"
	"regexp"
)

func (r *RPCClient) GetWork() ([]string, error) {
	rpcResp, err := r.doPost(r.Url, "eth_getWork", []string{})
	if err != nil {
		return nil, err
	}
	var reply []string
	err = json.Unmarshal(*rpcResp.Result, &reply)
	return reply, err
}

func (r *RPCClient) VerifyShare(params []string) (bool, error) {
	rpcResp, err := r.doPost(r.Url, "eth_posShareCheck", params)
	if err != nil {
		return false, err
	}
	var reply bool
	err = json.Unmarshal(*rpcResp.Result, &reply)
	return reply, err
}

func (r *RPCClient) SubmitBlock(params []string) (bool, error) {
	rpcResp, err := r.doPost(r.Url, "eth_submitWork", params)
	if err != nil {
		return false, err
	}
	var reply bool
	err = json.Unmarshal(*rpcResp.Result, &reply)
	return reply, err
}

func isZeroHash(s string) bool {
	var zeroHash = regexp.MustCompile("^0?x?0+$")
	return zeroHash.MatchString(s)
}

func (r *RPCClient) SendTransaction(from, to, gas, gasPrice, value string, autoGas bool) (string, error) {
	params := map[string]string{
		"from":  from,
		"to":    to,
		"value": value,
	}
	if !autoGas {
		params["gas"] = gas
		params["gasPrice"] = gasPrice
	}
	rpcResp, err := r.doPost(r.Url, "eth_sendTransaction", []interface{}{params})
	var reply string
	if err != nil {
		return reply, err
	}
	err = json.Unmarshal(*rpcResp.Result, &reply)
	if err != nil {
		return reply, err
	}
	/* There is an inconsistence in a "standard". Geth returns error if it can't unlock signer account,
	 * but Parity returns zero hash 0x000... if it can't send tx, so we must handle this case.
	 * https://github.com/ethereum/wiki/wiki/JSON-RPC#returns-22
	 */
	if isZeroHash(reply) {
		err = errors.New("transaction is not yet available")
	}
	return reply, err
}
