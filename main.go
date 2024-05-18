package main

import (
	"career_focus_hw/app/constants"
	"career_focus_hw/app/domain/service"
	"log"
)

func main() {
	address := "0x0000000000000000000000000000000000000000"
	server := service.NewParser(constants.URL)
	block := server.GetCurrentBlock()
	log.Printf("current block:%d", block)
	subscribeResult := server.Subscribe("address")
	log.Printf("subscribe to addresss:%s result:%s", address, subscribeResult)
	transactions := server.GetTransactions("address")
	log.Printf("get transactions from address:%s, size:%d", address, len(transactions))
}
