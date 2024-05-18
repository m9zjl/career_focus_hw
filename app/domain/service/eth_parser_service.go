package service

import (
	"career_focus_hw/app/api"
	"career_focus_hw/app/domain/repository"
	"career_focus_hw/app/infra/rpc"
	"log"
	"sync"
)

type EthParser struct {
	Url         string
	repo        repository.EthRepo
	lock        sync.RWMutex
	client      rpc.EthClient
	subscribers map[string]bool
}

func NewParser(
	url string,
	repo repository.EthRepo,
	client rpc.EthClient,
) api.Parser {
	return &EthParser{
		Url:         url,
		repo:        repo,
		lock:        sync.RWMutex{},
		client:      client,
		subscribers: map[string]bool{},
	}
}

func (e *EthParser) GetCurrentBlock() int {
	resp, err := e.client.GetRecentBlockNumber()
	if err != nil {
		log.Fatalf("Error getting current block: %v", err)
		return -1
	}
	return resp.Id
}

func (e *EthParser) Subscribe(address string) bool {
	if subscribed, ok := e.subscribers[address]; ok && subscribed {
		return true
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	if subscribed, ok := e.subscribers[address]; !ok || !subscribed {
		e.subscribers[address] = true
	}
	return true
}

func (e *EthParser) GetTransactions(address string) []api.Transaction {
	transactions, err := e.repo.ById(address)
	if err != nil {
		log.Fatalf("Error getting transactions: %v", err)
		return nil
	}
	return transactions
}
