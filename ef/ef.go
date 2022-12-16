package ef

import (
	"context"
	"errors"
	fga "flexgo/abi"
	"flexgo/fsm"
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EF struct {
	Sent bool
}

func New() *EF {
	return &EF{
		Sent: false,
	}
}

func (e *EF) Start(wg *sync.WaitGroup, ec *ethclient.Client, senderPK, senderA string) {
	defer wg.Done()

	// -------------------- Variables -------------------- //
	var (
		pk       = crypto.ToECDSAUnsafe(common.FromHex(senderPK))
		to       = common.HexToAddress("0xD4A0E3EC2A937E7CCa4A192756a8439A8BF4bA91")
		from     = senderA
		gasLimit = uint64(300000)
		fsm      = fsm.New()
	)

	cid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Printf("EF: failed to get chain id: %v", err.Error())
		return
	}

	// Pending nonce for given account
	n, err := ec.PendingNonceAt(context.Background(), common.HexToAddress(from))
	if err != nil {
		log.Printf("EF: failed to get nonce: %v", err.Error())
		return
	}

	// Gas variables
	tc, _ := ec.SuggestGasTipCap(context.Background())
	gc, _ := ec.SuggestGasPrice(context.Background())

	tc = tc.Mul(tc, big.NewInt(115))
	gc = gc.Mul(gc, big.NewInt(4))

	// -------------------- Variables -------------------- //

	ch := make(chan *types.Header)

	// Subscribe to new heads from the client
	_, err = ec.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		log.Printf("EF: failed to sub to new head: %v", err.Error())
		return
	}

	// On new head ...
	for h := range ch {
		log.Printf("\x1b[33m%s\x1b[32m%s\x1b[0m%v", "FC:", " Block number ", h.Number.Int64())

		// -------------------- ABI Encoding -------------------- //
		a := fga.ExternallyFundedABI

		babi, err := abi.JSON(strings.NewReader(a))
		if err != nil {
			log.Printf("EF: failed to parse abi: %v", err.Error())
			return
		}

		d, err := babi.Pack("updateResult")
		if err != nil {
			log.Printf("EF: failed to pack data: %v", err.Error())
			return
		}
		// -------------------- ABI Encoding -------------------- //

		txn := types.NewTx(&types.DynamicFeeTx{
			ChainID:   cid,
			Nonce:     n,
			GasTipCap: tc,
			GasFeeCap: gc,
			Gas:       gasLimit,
			To:        &to,
			Value:     big.NewInt(0),
			Data:      d,
		})

		passed, err := fsm.CheckDelay(ec)
		if err != nil {
			switch err.Error() {
			case errors.New("FSM: delay not passed").Error():
				log.Printf("\x1b[31m%s\x1b[0m", "EF: Delay not passed")
				continue
			default:
				log.Printf("EF: failed to check delay: %v", err.Error())
				return
			}
		}

		if passed {
			log.Printf("\x1b[32m%s\x1b[0m", "Delay passed")
			log.Printf("\x1b[32m%s\x1b[0m%s", "Signing tx: ", txn.Hash().String())

			st, err := types.SignTx(txn, types.NewLondonSigner(cid), pk)
			if err != nil {
				log.Printf("EF: failed to sign tx: %v", err.Error())
				return
			}

			// TODO: Handle possible high gas price
			if can := st.Cost().IsInt64(); !can {
				log.Printf("\x1b[31m%s\x1b[0m", "Tx cost may be too high")
				log.Printf("Tx cost: %v", st.Cost().String())
				log.Printf("Gas: %v", st.Gas())
				log.Printf("Gas price: %v", st.GasPrice())
			} else {
				log.Printf("\x1b[33m%s\x1b[0m%v", "Tx cost:", st.Cost().Int64()*(10/8))
				log.Printf("\x1b[33m%s\x1b[0m%v", "Gas: ", (st.Gas()))
				log.Printf("\x1b[33m%s\x1b[0m%v", "Gas price: ", st.GasPrice().Int64()*(10/8))
			}

			// res, err := ec.CallContract(
			// 	context.Background(), geth.CallMsg{
			// 		From:     common.HexToAddress(from),
			// 		To:       &to,
			// 		Gas:      gasLimit,
			// 		GasPrice: gc,
			// 		Value:    big.NewInt(0),
			// 		Data:     []byte("0x80ebb08e")}, nil)

			log.Printf("\x1b[32m%s\x1b[0m%s", "Sending tx: ", st.Hash().String())
			if e.Sent != true {
				err = ec.SendTransaction(context.Background(), st)

				// Parse error
				if err != nil {
					switch err.Error() {
					case errors.New("execution reverted: ExternallyFundedOSM/not-passed").Error():
						log.Printf(errors.New("EF: OSM not passed").Error())
						e.Sent = false
						continue
					case errors.New("already known").Error():
						log.Printf(errors.New("EF: already known").Error())
						e.Sent = false
						continue
					case errors.New("execution reverted").Error():
						log.Printf(errors.New("EF: execution reverted").Error())
						e.Sent = false
						continue
					default:
						c := "err: max fee per gas less than"
						e.Sent = false
						if len(err.Error()) > 30 && err.Error()[0:len(c)] == errors.New(c).Error() {
							log.Printf(errors.New("EF: max fee per gas less than block base fee").Error())
							continue
						}

						log.Printf("EF: failed to send tx: %v", err.Error())
						return
					}
				}

				e.Sent = false
			}

			// log.Printf("\x1b[32m%s\x1b[0m%s", "Sending tx: ", st.Hash().String())

			// err = ec.SendTransaction(context.Background(), st)
			// if err != nil {
			// 	log.Printf("failed to send tx: %v", err.Error())
			// 	return
			// }
			// }
		}
	}
}
