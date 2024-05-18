package api

type Transaction struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	Signature            string `json:"signature"`
	Nonce                uint64 `json:"nonce"`
	Value                uint64 `json:"value"`
	Input                string `json:"input"`
	GasLimit             uint64 `json:"gasLimit"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas         string `json:"MaxFeePerGas"`
}
