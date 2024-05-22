package service

import (
	"career_focus_hw/app/domain/repository"
	"career_focus_hw/app/infra/rpc"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ParserDaemon struct {
	lock        sync.RWMutex
	latestBlock int64
	client      *rpc.EthClient
	subscribers map[string]bool
	ethRepo     repository.EthRepo
}

func NewParserDaemon(
	client *rpc.EthClient,
	ethRepo repository.EthRepo,
) *ParserDaemon {
	return &ParserDaemon{
		lock:        sync.RWMutex{},
		latestBlock: -1,
		client:      client,
		subscribers: make(map[string]bool),
		ethRepo:     ethRepo,
	}
}

func (pd *ParserDaemon) Init() {
	go pd.start()
}

func (pd *ParserDaemon) start() {
	for {
		pd.updateFromRemoteServer()
		time.Sleep(1 * time.Second)
	}
}

func (pd *ParserDaemon) updateFromRemoteServer() {
	blockNumResp, err := pd.client.GetRecentBlockNumber()
	if err != nil {
		// handle error
		log.Fatalln(fmt.Sprint("Error: ", err))
		return
	}
	blockNum := hexToInt(blockNumResp.Result)
	if pd.latestBlock == -1 {
		pd.parseBlockByBlockNum(blockNum)
	} else {
		for i := pd.latestBlock; i <= blockNum; i++ {
			pd.parseBlockByBlockNum(i)
		}
	}
	pd.latestBlock = blockNum

}

var forTest = true

func (pd *ParserDaemon) parseBlockByBlockNum(blockNum int64) {
	resp, err := pd.client.GetBlockByNum(intToHex(blockNum))
	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to get block %d", blockNum))
	}
	if resp == nil || resp.Result.Transactions == nil {
		log.Println(fmt.Sprintf("no transactions in block %d", blockNum))
		return
	}

	for _, tx := range resp.Result.Transactions {
		if forTest {
			log.Println(fmt.Sprintf("test address:%s", tx.From))
			forTest = false
		}
		if pd.subscribers[tx.To] {
			_, err = pd.ethRepo.Save(tx.To, tx)
			if err != nil {
				log.Fatalln(fmt.Sprintf("failed to save block %d", blockNum))
			}
		}
		if pd.subscribers[tx.From] {
			_, err = pd.ethRepo.Save(tx.From, tx)
			if err != nil {
				log.Fatalln(fmt.Sprintf("failed to save block %d", blockNum))
			}
		}
	}
}

func (pd *ParserDaemon) subscribe(address string) bool {
	pd.lock.Lock()
	defer pd.lock.Unlock()
	if ok := pd.subscribers[address]; !ok {
		pd.subscribers[address] = true
	}
	return true
}

func hexToInt(str string) int64 {
	parsed, _ := strconv.ParseInt(strings.Replace(str, "0x", "", -1), 16, 32)
	return parsed
}

func intToHex(num int64) string {
	return fmt.Sprintf("0x%x", num)
}
