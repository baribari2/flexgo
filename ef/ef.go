package ef

import (
	"context"
	"errors"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EF struct{}

func (e *EF) Start(wg *sync.WaitGroup, ec *ethclient.Client, senderPK, senderA string) {
	defer wg.Done()

	// -------------------- Variables -------------------- //
	var (
		pk       = crypto.ToECDSAUnsafe(common.FromHex(senderPK))
		to       = common.HexToAddress("0xD4A0E3EC2A937E7CCa4A192756a8439A8BF4bA91")
		from     = senderA
		gasLimit = uint64(3000000)
	)
	cid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain id: %v", err.Error())
		return
	}

	n, err := ec.PendingNonceAt(context.Background(), common.HexToAddress(from))
	if err != nil {
		log.Printf("Failed to get nonce: %v", err.Error())
		return
	}

	tc, _ := ec.SuggestGasTipCap(context.Background())
	gc, _ := ec.SuggestGasPrice(context.Background())

	// -------------------- Variables -------------------- //

	ch := make(chan *types.Header)

	// Subscribe to new heads from the client
	_, err = ec.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		log.Printf("Failed to sub to new head: %v", err.Error())
		return
	}

	// // PassedDelay tx
	// passed, err := exf.PassedDelay(&bind.CallOpts{Context: context.Background()})
	// if err != nil {
	// 	log.Printf("Failed to call PassedDelay: %v", err.Error())
	// 	return
	// }

	// On new head ...
	for h := range ch {
		log.Printf("New head: %v", h)

		txn := types.NewTx(&types.DynamicFeeTx{
			ChainID:   cid,
			Nonce:     n,
			GasTipCap: tc,
			GasFeeCap: gc,
			Gas:       gasLimit,
			To:        &to,
			Value:     big.NewInt(0),
			Data:      []byte("0x80ebb08e"),
		})

		log.Printf("Signing tx: %v", txn.Hash().String())

		st, err := types.SignTx(txn, types.NewLondonSigner(cid), pk)
		if err != nil {
			log.Printf("Failed to sign tx: %v", err.Error())
			return
		}

		err = ec.SendTransaction(context.Background(), st)

		// Parse error
		if err != nil {
			if errors.New("execution reverted: ExternallyFundedOSM/not-passed").Error() == err.Error() {
				log.Printf(errors.New("EF: OSM not passed").Error())
				return
			} else {
				log.Printf("Failed to call UpdateResult: %v", err.Error())
				return
			}
		}
	}
}
