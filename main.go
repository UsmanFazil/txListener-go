package main

import (
	"context"
	"fmt"
	"log"

	"github.com/block-listener/conf"
	service "github.com/block-listener/services"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(conf.GetConfig().WsRpc)
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	lastConfimedBlock, err := service.GetLastConfirmedNum()
	if err != nil {
		log.Fatal(err)
	}

	syncFlag := true

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("New block :", block.Number().Uint64())

			if syncFlag && lastConfimedBlock != nil {
				go service.SyncBlocks(uint64(lastConfimedBlock.Blocknum), block.Number().Uint64())
				syncFlag = false
			}
			go service.FindTx(block)
		}
	}

}
