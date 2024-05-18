package service

import (
	"career_focus_hw/app/api"
	"career_focus_hw/app/domain/repository"
	"career_focus_hw/app/infra/rpc"
	"sync"
)

type EthParser struct {
	Url    string
	repo   repository.EthRepo
	lock   sync.RWMutex
	client *rpc.EthClient
}

func NewParser(
	url string,
	repo repository.EthRepo,
	client *rpc.EthClient,
) api.Parser {
	return &EthParser{
		Url:    url,
		repo:   repo,
		lock:   sync.RWMutex{},
		client: client,
	}
}

func (e *EthParser) GetCurrentBlock() int {
	panic(nil)
}

func (e *EthParser) Subscribe(address string) bool {
	return false
}

func (e *EthParser) GetTransactions(address string) []api.Transaction {
	return nil
}
