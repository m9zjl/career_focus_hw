package service

import (
	"career_focus_hw/app/api"
	"career_focus_hw/app/domain/repository"
	"log"
	"sync"
)

type EthParser struct {
	Url          string
	repo         repository.EthRepo
	lock         sync.RWMutex
	currentBlock int64
	daemon       *ParserDaemon
}

func NewParser(
	url string,
	repo repository.EthRepo,
	daemon *ParserDaemon,
) api.Parser {
	return &EthParser{
		Url:          url,
		repo:         repo,
		lock:         sync.RWMutex{},
		currentBlock: -1,
		daemon:       daemon,
	}
}

func (e *EthParser) GetCurrentBlock() int64 {
	return e.daemon.latestBlock
}

func (e *EthParser) Subscribe(address string) bool {
	return e.daemon.subscribe(address)
}

func (e *EthParser) GetTransactions(address string) []api.Transaction {
	transactions, err := e.repo.ById(address)
	if err != nil {
		log.Fatalf("Error while getting transactions from repo: %v", err)
	}
	return transactions
}
