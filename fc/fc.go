package fc

import (
	"context"
	fga "flexgo/abi"
	"log"
	"math/big"
	"strings"
	"sync"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type FC struct {
	Sent bool
}

func New() *FC {
	return &FC{
		Sent: false,
	}
}

func (f *FC) Start(wg *sync.WaitGroup, ec *ethclient.Client, senderPK, senderA string) {
	defer wg.Done()

	// -------------------- Variables -------------------- //
	var (
		pk       = crypto.ToECDSAUnsafe(common.FromHex(senderPK))
		to       = common.HexToAddress("0x34515A81a6bFf50925295a363cE0b1d1bddE6620")
		from     = senderA
		gasLimit = uint64(300000)
	)

	cid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Printf("FC: failed to get chain id: %v", err.Error())
		return
	}

	// Pending nonce for given account
	n, err := ec.PendingNonceAt(context.Background(), common.HexToAddress(from))
	if err != nil {
		log.Printf("FC: failed to get nonce: %v", err.Error())
		return
	}

	// Gas variables
	tc, _ := ec.SuggestGasTipCap(context.Background())
	gc, _ := ec.SuggestGasPrice(context.Background())

	tc = tc.Mul(tc, big.NewInt(115))
	gc = gc.Mul(gc, big.NewInt(4))

	// -------------------- Variables -------------------- //

	ch := make(chan *types.Header)

	_, err = ec.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		log.Printf("FC: failed to subscribe to new head: %v", err.Error())
		return
	}

	for h := range ch {
		log.Printf("\x1b[33m%s\x1b[32m%s\x1b[0m%v", "FC:", " Block number ", h.Number.Int64())

		// -------------------- ABI Encoding -------------------- //
		f := fga.FlexCallerABI

		babi, err := abi.JSON(strings.NewReader(f))
		if err != nil {
			log.Printf("FC: failed to parse abi: %v", err.Error())
			return
		}

		d, err := babi.Pack("flexCall")
		if err != nil {
			log.Printf("FC: failed to pack data: %v", err.Error())
			return
		}
		// -------------------- ABI Encoding -------------------- //

		res, err := ec.PendingCallContract(context.Background(), geth.CallMsg{
			From:     common.HexToAddress(from),
			To:       &to,
			Gas:      gasLimit,
			GasPrice: gc,
			Value:    big.NewInt(0),
			Data:     d,
		})

		if res != nil {
			log.Printf("\x1b[33m%s\x1b[0m%v", "Result ", string(res))
		}

		if err != nil {
			log.Printf("FC: failed to call contract: %v", err.Error())
			continue
		}

		// -------------------- Transaction -------------------- //

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

		_, err = types.SignTx(txn, types.NewLondonSigner(cid), pk)
		if err != nil {
			log.Printf("FC: failed to sign transaction: %v", err.Error())
			return
		}
		// -------------------- Transaction -------------------- //
	}

}
