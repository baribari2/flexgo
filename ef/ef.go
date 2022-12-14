package ef

import (
    "log"
    fga "flexgo/abi"
    "sync"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
)

type EF struct {}

func (e *EF) Start(wg *sync.WaitGroup, ec *ethclient.Client) {
    defer wg.Done()


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


    exf, err := fga.NewExternallyFunded(common.HexToAddress("0xD4A0E3EC2A937E7CCa4A192756a8439A8BF4bA91"), auth)
    if err != nil {
        log.Printf("Failed to make EF instance: %v", err.Error())
        return
    }

    ch := make(chan *types.Header)

    _, err := ec.SubscribeNewHead(context.Background(), ch)
    if err != nil {
        log.Printf("Failed tp sub to new head: %v", err.Error()
        return
    }

    log.Printf("EF: %v", exf)

    // for h := ch {

}

