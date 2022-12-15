package mempool

import (
	"context"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Mempool struct{}

func New() *Mempool {
	return &Mempool{}
}

func (m *Mempool) Start(wg *sync.WaitGroup, rc *rpc.Client, ec *ethclient.Client) {
	defer wg.Done()

	gc := gethclient.New(rc)

	ch := make(chan common.Hash)
	_, err := gc.SubscribePendingTransactions(context.Background(), ch)
	if err != nil {
		log.Printf("failed to sub to pending txs: %v", err.Error())
		return
	}

	for h := range ch {
		tx, _, _ := ec.TransactionByHash(context.Background(), h)

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

		// log.Printf("%s\x1b[32m%s\x1b[0m%s", "---------- ", "Transaction Found", " ----------")
		// log.Printf("Hash: %v", h)
		// log.Printf("To: %v", tx.To())
		// log.Printf("Cost: %v", tx.Cost())
		// log.Printf("Gas Limit: %v", tx.Gas())
		// log.Printf("Nonce: %v", tx.Nonce())
		// log.Printf("Data: %v", tx.Data())
	}

	return
}
