package fsm

import (
	"context"
	"errors"
	"flexgo/abi"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type FSMWatcher struct{}

func New() *FSMWatcher {
	return &FSMWatcher{}
}

func (f *FSMWatcher) CheckDelay(ec *ethclient.Client) (bool, error) {
	// --------------------- Variables --------------------- //
	fsm, err := abi.NewFSM(common.HexToAddress("0x105b857583346E250FBD04a57ce0E491EB204BA3"), ec)
	if err != nil {
		return false, err
	}

	// FSM last reimburse time
	rt, err := fsm.LastReimburseTime(nil)
	if err != nil {
		return false, err
	}

	// Latest block
	b, err := ec.BlockByNumber(context.Background(), nil)
	if err != nil {
		return false, err
	}

	// Now (block.timestamp)
	n := b.Time()
	// --------------------- Variables --------------------- //

	if can := rt.IsInt64(); !can {
		return false, errors.New("FSM: rt is not int64")
	}

	if ok := int64(n)-rt.Int64() >= int64(3600) || int64(3600)-(int64(n)-rt.Int64()) <= int64(15); !ok {
		log.Printf("FSM Check: %v", strconv.FormatInt(int64(3600)-(int64(n)-rt.Int64()), 10)+" seconds until 1 hr has passed")
		return false, errors.New("FSM: delay not passed")
	}

	return true, nil
}
