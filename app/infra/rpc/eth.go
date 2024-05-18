package rpc

import (
	"career_focus_hw/app/api"
	"encoding/json"
	"net/http"
	"strings"
)

type EthClient struct {
	url string
	seq int
}

func NewEthClient(url string) EthClient {
	return EthClient{
		url: url,
	}
}

func (client *EthClient) doRequest(method string, params []interface{}) (*http.Response, error) {
	defer func() { client.seq++ }()
	req := api.RpcRequest{"2.0", method, params, client.seq}
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(client.url, "application/json", strings.NewReader(string(marshal)))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetRecentBlockNumber curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"seq":83}'
func (client *EthClient) GetRecentBlockNumber() (*api.RpcResponse, error) {
	resp, err := client.doRequest("eth_blockNumber", []interface{}{})
	if err != nil {
		return nil, err
	}
	var ans = new(api.RpcResponse)
	if err := json.NewDecoder(resp.Body).Decode(ans); err != nil {
		return nil, err
	}
	return ans, nil
}
