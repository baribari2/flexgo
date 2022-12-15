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
		log.Printf("failed to instantiate FSM: %v", err.Error())
		return false, err
	}

	// FSM last reimburse time
	rt, err := fsm.LastReimburseTime(nil)
	if err != nil {
		log.Printf("failed to get last reimburse time: %v", err.Error())
		return false, err
	}

	// Latest block
	b, err := ec.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Printf("failed to get latest block: %v", err.Error())
		return false, err
	}

	// Now (block.timestamp)
	n := b.Time()
	// --------------------- Variables --------------------- //

	if can := rt.IsInt64(); !can {
		return false, errors.New("rt is not int64")
	}

	if ok := int64(n)-rt.Int64() >= int64(3600); !ok {
		log.Printf("FSM Check: %v", strconv.FormatInt(int64(3600)-(int64(n)-rt.Int64()), 10)+" seconds until 1hr has passed")
		return false, errors.New("FSM: delay not passed")
	}

	return true, nil
}
