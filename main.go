package main

import (
	"flexgo/fc"
	"flexgo/mempool"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	wg := &sync.WaitGroup{}
	rc, err := rpc.Dial("ws://localhost:8546")
	if err != nil {
		log.Printf("failed to dial rpc: %v", err.Error())
		return
	}

	c, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/f0e8a1d30f2f4b66bee88b508d52708f")
	if err != nil {
		log.Printf("failed to dial websocket: %s", err.Error())
		return
	}

	wg.Add(2)

	log.Println("Starting mempool listener")
	ml := mempool.New()
	go ml.Start(wg, rc, c)

	log.Println("Starting FC listener")
	fc := fc.New()
	go fc.Start(wg, c, PK, ME)

	// log.Println("Starting EF listener")
	// exf := ef.New()
	// go exf.Start(wg, c, PK, ME)

	wg.Wait()
}
