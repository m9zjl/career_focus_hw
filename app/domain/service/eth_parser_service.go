package service

import (
	"career_focus_hw/app/api"
	"career_focus_hw/app/domain/repository"
	"career_focus_hw/app/infra/rpc"
	"log"
	"strconv"
	"sync"
	"time"
)

type EthParser struct {
	Url          string
	repo         repository.EthRepo
	lock         sync.RWMutex
	client       rpc.EthClient
	subscribers  map[string]bool
	currentBlock int
}

func NewParser(
	url string,
	repo repository.EthRepo,
	client rpc.EthClient,
) api.Parser {
	return &EthParser{
		Url:          url,
		repo:         repo,
		lock:         sync.RWMutex{},
		client:       client,
		subscribers:  map[string]bool{},
		currentBlock: -1,
	}
}

func (e *EthParser) run() {
	for {
		e.fetchEthTransactions()
		time.Sleep(1 * time.Second)
	}
}

func (e *EthParser) fetchEthTransactions() {
	resp, err := e.client.GetEthBlockNumber()
	if err != nil {
		log.Fatalf("Error getting current block: %v", err)
		return
	}
	blockNumStr := resp.Result
	blockNum, err := strconv.Atoi(blockNumStr)
	if err != nil {
		log.Fatalf("Error converting block number to int: %v", err)
		return
	}
	if e.currentBlock == -1 {
		e.fetchBlockByBlockNum(blockNum)
	} else {
		for curBlock := e.currentBlock; curBlock < blockNum; curBlock++ {
			e.fetchBlockByBlockNum(curBlock)
		}
	}
	e.currentBlock = blockNum
}

func (e *EthParser) fetchBlockByBlockNum(block int) {
	resp, err := e.client.GetBlcokByNum(block)
	if err != nil {
		log.Fatalf("Error getting block by number: %v", err)
	}
	if len(resp.Transactions) == 0 {
		return
	}
	for _, transaction := range resp.Transactions {
		if _, ok := e.subscribers[transaction.To]; ok {
			save, err := e.repo.Save(transaction.To, transaction)
			if err != nil || !save {
				log.Fatalf("Error saving transaction: %v", err)
				return
			}
			save, err = e.repo.Save(transaction.From, transaction)
			if err != nil || !save {
				log.Fatalf("Error saving transaction: %v", err)
				return
			}

		}
	}
}

func (e *EthParser) GetCurrentBlock() int {
	return e.currentBlock
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
