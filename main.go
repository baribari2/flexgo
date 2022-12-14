package main

import (
	"flexgo/ef"
	"flexgo/mempool"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	log.Println("Starting mempool listener")

	wg := &sync.WaitGroup{}
	rc, err := rpc.Dial("ws://localhost:8546")
	if err != nil {
		log.Printf("Failed to dial rpc: %v", err.Error())
		return
	}

	c, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Printf("Failed to dial websocket: %s", err.Error())
		return
	}

	wg.Add(2)

	log.Println("Starting mempool listener")
	ml := &mempool.Mempool{}
	go ml.Start(wg, rc, c)

	log.Println("Starting EF listener")
	exf := &ef.EF{}
	go exf.Start(wg, c, PK, ME)

	wg.Wait()
}
