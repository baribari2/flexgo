package ef

import (
    "log"
    fga "flexgo/abi"
    "sync"
    // "crypto/ecdsa"
    "crypto"
    "context"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts"
    "github.com/ethereum/go-ethereum/accounts/keystore"
)

type EF struct {}

func (e *EF) Start(wg *sync.WaitGroup, ec *ethclient.Client) {
    defer wg.Done()

    ks := keystore.NewKeyStore("~/ethereum/keystore", keystore.StandarnScryptN, keystore.StandardScryptP)
    am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)

    exf, err := fga.NewExternallyFunded(common.HexToAddress("0xD4A0E3EC2A937E7CCa4A192756a8439A8BF4bA91"), auth)
    if err != nil {
        log.Printf("Failed to make EF instance: %v", err.Error())
        return
    }

    ch := make(chan *types.Header)

    _, err = ec.SubscribeNewHead(context.Background(), ch)
    if err != nil {
        log.Printf("Failed tp sub to new head: %v", err.Error())
        return
    }

    log.Printf("EF: %v", exf)

    // for h := ch {

    return
}

