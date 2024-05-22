package rpc

import (
	"career_focus_hw/app/api"
	"encoding/json"
	"io"
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
	req := api.RpcRequest{JsonRpc: "2.0", Method: method, Params: params, Id: client.seq}
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

func closeResponse(resp *http.Response) {
	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
}

// GetRecentBlockNumber curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"seq":83}'
func (client *EthClient) GetRecentBlockNumber() (*api.GetRecentBlockResp, error) {
	resp, err := client.doRequest("eth_blockNumber", []interface{}{})
	if err != nil {
		return nil, err
	}
	defer closeResponse(resp)
	var ans = new(api.GetRecentBlockResp)
	if err := json.NewDecoder(resp.Body).Decode(ans); err != nil {
		return nil, err
	}
	return ans, nil
}

func (client *EthClient) GetBlockByNum(block string) (*api.GetBlockByNumberResp, error) {
	resp, err := client.doRequest("eth_getBlockByNumber", []interface{}{block, true})
	if err != nil {
		return nil, err
	}
	defer closeResponse(resp)
	var ans = new(api.GetBlockByNumberResp)
	if err := json.NewDecoder(resp.Body).Decode(ans); err != nil {
		return nil, err
	}
	return ans, nil
}
