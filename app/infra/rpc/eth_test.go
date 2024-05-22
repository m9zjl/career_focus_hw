package rpc

import (
	"testing"
)

const URL = "https://cloudflare-eth.com"

func getEthRpClient() EthClient {
	return NewEthClient(URL)
}

func TestEthClient_GetRecentBlockNumber(t *testing.T) {
	rpcClient := getEthRpClient()
	resp, err := rpcClient.GetRecentBlockNumber()
	if err != nil {
		t.Errorf("GetRecentBlockNumber() error = %v", err)
	}
	if len(resp.Result) <= 0 {
		t.Errorf("GetRecentBlockNumber() number = %v", resp.Result)
	}
}

func TestEthClient_GetBlockByNum(t *testing.T) {
	rpcClient := getEthRpClient()
	resp, err := rpcClient.GetRecentBlockNumber()
	if err != nil {
		t.Errorf("GetRecentBlockNumber() error = %v", err)
	}

	blockByNumResp, err := rpcClient.GetBlockByNum(resp.Result)
	if err != nil {
		t.Errorf("GetRecentBlockNumber() error = %v", err)
	}
	if len(blockByNumResp.Result.Transactions) <= 0 {
		t.Errorf("GetRecentBlockNumber() error = %v", err)
	}
}
