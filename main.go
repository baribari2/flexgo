package main

import (
	"log"
	"flexgo/mempool"
    // fga "flexgo/abi"

    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/ethclient"
    // "github.com/ethereum/go-ethereum/accounts/abi/bind"
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
