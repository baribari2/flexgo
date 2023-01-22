package main

import (
	"flexgo/discord"
	"flexgo/ef"
	"flexgo/fc"
	"flexgo/mempool"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	wg := &sync.WaitGroup{}
	rc, err := rpc.Dial("ws://localhost:8546")
	if err != nil {
		log.Printf("failed to dial rpc: %v", err.Error())
		return
	}

	c, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Printf("failed to dial websocket: %s", err.Error())
		return
	}

	d, err := discord.New(DTOKEN, CHANNEL)
	if err != nil {
		log.Printf("failed to create discord: %v", err.Error())
		return
	}

	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	wg.Add(4)

	log.Println("Starting mempool listener")
	ml := mempool.New()
	go ml.Start(wg, rc, c)

	log.Println("Starting FC listener")
	fc := fc.New()
	go fc.Start(wg, c, PK2, ADD2)

	log.Println("Starting EF listener")
	exf := ef.New()
	go exf.Start(wg, c, d, PKMAIN, ADDMAIN)

	wg.Wait()
}
