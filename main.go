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
	service.OpenTx(client)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	lastConfimedBlock, err := service.GetBlockSyncInfo()
	if err != nil {
		log.Fatal(err)
	}

	firstRun := true

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

			if firstRun && lastConfimedBlock != nil {

				syncStartNum := uint64(lastConfimedBlock.Backupsyncnum)
				if lastConfimedBlock.Syncstatus == 1 {
					syncStartNum = uint64(lastConfimedBlock.Blocksyncnum)
				}

				go service.SyncBlocks(syncStartNum, block.Number().Uint64(), lastConfimedBlock)
			}
			go service.FindTx(block, false)
			firstRun = false
		}
	}

}
