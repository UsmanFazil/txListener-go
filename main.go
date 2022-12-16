package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/block-listener/conf"
	service "github.com/block-listener/services"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	go chainService(conf.GetConfig().EthData)
	go chainService(conf.GetConfig().BscData)
	go chainService(conf.GetConfig().CronosData)

	http.HandleFunc("/getUserTx", service.GetUserTx)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func chainService(chainInfo conf.ChainData) {

	client, err := ethclient.Dial(chainInfo.WsRpc)
	if err != nil {
		log.Fatal(err)
	}
	// service.OpenTx(client, chainInfo.ChainId)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ChainId:", chainInfo.ChainId)
	lastConfimedBlock, err := service.GetBlockInfobyChainId(chainInfo.ChainId)
	if err != nil {
		log.Fatal(err)
	}

	firstRun := true

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("restarting service of chain id", chainInfo.ChainId, "and error is:", err)
		case header := <-headers:
			block, err := client.BlockByNumber(context.Background(), header.Number)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("New block :", block.Number().Uint64())

			if firstRun && lastConfimedBlock != nil {

				syncStartNum := uint64(lastConfimedBlock.Backupsyncnum)
				if lastConfimedBlock.Syncstatus == 1 {
					syncStartNum = uint64(lastConfimedBlock.Blocksyncnum)
				}

				go service.SyncBlocks(syncStartNum, block.Number().Uint64(), lastConfimedBlock, chainInfo.ContractAddress, chainInfo.WsRpc, chainInfo.ChainId, client)
			}

			go service.FindTx(block, false, chainInfo.ContractAddress, chainInfo.ChainId, client)

			firstRun = false
		}
	}
}
