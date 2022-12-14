package main

import (
	"log"
	"flexgo/mempool"
    fga "flexgo/abi"

    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
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

    pk, err := crypto.HexToECDSA(PK)
    if err != nil {
        log.Printf("Failed to creake private key: %v", err.Error())
        return
    }

    pubK := pk.Public()
    pke, ok := pubK.(ecdsa.*publicKeyECDSA)
    if !ok {
        log.Printf("Failed to create public key: %v", err.Error())
        return 
    }

    auth := bind.NewKeyedTransactor(pk)

    fga.NewExternallyFunded(common.HexToAddress("0xD4A0E3EC2A937E7CCa4A192756a8439A8BF4bA91"), auth)
}
