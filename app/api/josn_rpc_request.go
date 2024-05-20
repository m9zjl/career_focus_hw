package api

type RpcRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type RpcResponse struct {
	JsonRpc      string        `json:"jsonrpc"`
	Result       string        `json:"result"`
	Transactions []Transaction `json:"transactions"`
	Id           int           `json:"id"`
}
