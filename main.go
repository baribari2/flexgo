package main

import (
	"log"

	"flexgo/mempool"
)

func main() {
	log.Println("Starting mempool listener")

	ml := &mempool.MempoolListener{}
	ml.Start()
}
