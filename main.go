package main

import (
	"career_focus_hw/app/constants"
	"career_focus_hw/app/domain/repository"
	"career_focus_hw/app/domain/service"
	"career_focus_hw/app/infra/db"
	"career_focus_hw/app/infra/rpc"
	"log"
)

func main() {
	address := "0x0000000000000000000000000000000000000000"

	// new
	db := db.NewLocalMem()
	repo := repository.NewEthRepo(db)
	rpcClient := rpc.NewEthClient(constants.URL)

	// init
	server := service.NewParser(constants.URL, repo, rpcClient)

	block := server.GetCurrentBlock()
	log.Printf("current block:%d", block)
	subscribeResult := server.Subscribe("address")
	log.Printf("subscribe to addresss:%s result:%s", address, subscribeResult)
	transactions := server.GetTransactions("address")
	log.Printf("get transactions from address:%s, size:%d", address, len(transactions))
}
