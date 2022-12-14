package ef

import (
	"context"
	"errors"
	"log"
	"math/big"
	"sync"

	geth "github.com/ethereum/go-ethereum"
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
		gasLimit = uint64(100000)
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
		log.Printf("\x1b[32m%s\x1b[0m%v", "New head: ", h)

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

		log.Printf("\x1b[32m%s\x1b[0m%s", "Signing tx: ", txn.Hash().String())

		if can := txn.Cost().IsInt64(); !can {
			// TODO: Handle possible high gas price
			log.Printf("\x1b[31m%s\x1b[0m", "Tx cost may be too high")
			log.Printf("Tx cost: %v", txn.Cost().String())
			log.Printf("Gas: %v", txn.Gas())
			log.Printf("Gas price: %v", txn.GasPrice())
		} else {
			log.Printf("\x1b[33m%s\x1b[0m%v", "Tx cost:", txn.Cost().Int64()*(10/8))
			log.Printf("\x1b[33m%s\x1b[0m%v", "Gas: ", (txn.Gas()))
			log.Printf("\x1b[33m%s\x1b[0m%v", "Gas price: ", txn.GasPrice().Int64()*(10/8))
		}

		st, err := types.SignTx(txn, types.NewLondonSigner(cid), pk)
		if err != nil {
			log.Printf("Failed to sign tx: %v", err.Error())
			return
		}

		res, err := ec.CallContract(
			context.Background(), geth.CallMsg{
				From:     common.HexToAddress(from),
				To:       &to,
				Gas:      gasLimit,
				GasPrice: gc,
				Value:    big.NewInt(0),
				Data:     []byte("0x80ebb08e")}, nil)

		if res != nil {
			log.Printf("Result: %v", string(res))
		}

		// Parse error
		if err != nil {
			switch err.Error() {
			case errors.New("execution reverted: ExternallyFundedOSM/not-passed").Error():
				log.Printf(errors.New("EF: OSM not passed").Error())
				continue
			case errors.New("already known").Error():
				log.Printf(errors.New("EF: Already known").Error())
				continue
			case errors.New("execution reverted").Error():
				log.Printf(errors.New("EF: Execution reverted").Error())
				continue
			default:
				c := "err: max fee per gas less than"
				if len(err.Error()) > 30 && err.Error()[0:len(c)] == errors.New(c).Error() {
					log.Printf(errors.New("EF: Max fee per gas less than block base fee").Error())
					continue
				}

				log.Printf("Failed to send tx: %v", err.Error())
				return
			}
		}

		log.Printf("\x1b[32m%s\x1b[0m%s", "Sending tx: ", st.Hash().String())
		err = ec.SendTransaction(context.Background(), st)
	}
}
