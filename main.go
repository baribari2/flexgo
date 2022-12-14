package main

import (
	"log"
	"flexgo/mempool"

    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	log.Println("Starting mempool listener")

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

	ml := &mempool.Mempool{}
	ml.Start(rc, c)
}
