package main

import (
	"bufio"
	"career_focus_hw/app/constants"
	"career_focus_hw/app/domain/repository"
	"career_focus_hw/app/domain/service"
	"career_focus_hw/app/infra/db"
	"career_focus_hw/app/infra/rpc"
	"encoding/json"
	"log"
	"os"
	"strings"
)

// put this into config file in the future
//const address = "0x0000000000000000000000000000000000000000"

func main() {

	// init container
	storage := db.NewLocalMem()
	repo := repository.NewEthRepo(storage)
	rpcClient := rpc.NewEthClient(constants.URL)
	daemon := service.NewParserDaemon(&rpcClient, repo)
	daemon.Init()
	parser := service.NewParser(constants.URL, repo, daemon)
	// init

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {

		cmd := in.Text()

		if cmd == "exit" {
			log.Println("exiting...")
			return
		} else if strings.Index(cmd, "subscribe ") == 0 {
			splits := strings.Split(cmd, " ")
			if len(splits) != 2 {
				log.Printf("invalid command:%s", cmd)
				continue
			}
			address := splits[1]
			log.Printf("address to subscribe:%s", address)
			if parser.Subscribe(address) {
				log.Printf("successfully subscribed to address:%s", address)
			} else {
				log.Printf("failed to subscribe to address:%s", address)
			}
		} else if strings.Index(cmd, "transaction ") == 0 {
			splits := strings.Split(cmd, " ")
			if len(splits) != 2 {
				log.Printf("invalid command:%s", cmd)
				continue
			}
			address := splits[1]
			transactions := parser.GetTransactions(address)
			if len(transactions) == 0 {
				log.Printf("no transactions found for address:%s", address)
			} else {
				marshal, _ := json.Marshal(transactions)
				log.Printf("transactions found for address:%s as:", address)
				log.Printf("%s", string(marshal))
			}
		} else if "get_current_block" == cmd {
			currentBLock := parser.GetCurrentBlock()
			log.Printf("current block:%d", currentBLock)
		} else {
			log.Printf("unknown command:%s", cmd)
		}
	}
}
