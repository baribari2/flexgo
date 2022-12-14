package mempool

import (
    "log"
    "context"

    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/ethclient/gethclient"
    "github.com/ethereum/go-ethereum/common"
)

type Mempool struct {}

func (m *Mempool) Start(rc *rpc.Client, ec *ethclient.Client) {
    gc := gethclient.New(rc)

    ch := make(chan common.Hash)
    _, err := gc.SubscribePendingTransactions(context.Background(), ch)
    if err != nil {
        log.Printf("Failed to sub to pending txs: %v", err.Error())
        return
    }

    for h := range ch {
        tx, p, err := ec.TransactionByHash(context.Background(), h)
        if err != nil {
            log.Printf("Failed to fetch tx (%s) by hash: %v", h, err.Error())
        }

        if p != true {
            log.Printf("Transaaction hash %v is not pending", h)
            continue
        }

        h := common.HexToAddress("0xd4a0e3ec2a937e7cca4a192756a8439a8bf4ba91") 
        if tx.To() == &h {

        log.Printf("%s\x1b[31m%s\x1b[0m%s", "---------- ", "EF Transaction Found", " ----------")
        log.Printf("Hash: %v", h)
        log.Printf("To: %v", tx.To())
        log.Printf("Cost: %v", tx.Cost())
        log.Printf("Gas Limit: %v", tx.Gas())
        log.Printf("Nonce: %v", tx.Nonce())
        log.Printf("Data: %v", tx.Data())

        break
    }

        log.Printf("%s\x1b[32m%s\x1b[0m%s", "---------- ", "Transaction Found", " ----------")
        log.Printf("Hash: %v", h)
        log.Printf("To: %v", tx.To())
        log.Printf("Cost: %v", tx.Cost())
        log.Printf("Gas Limit: %v", tx.Gas())
        log.Printf("Nonce: %v", tx.Nonce())
        log.Printf("Data: %v", tx.Data())
    }

    return
}
