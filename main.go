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
	go chainService(conf.GetConfig().EthData.WsRpc, conf.GetConfig().EthData.ContractAddress, conf.GetConfig().EthData.ChainId)
	go chainService(conf.GetConfig().BscData.WsRpc, conf.GetConfig().BscData.ContractAddress, conf.GetConfig().BscData.ChainId)
	chainService(conf.GetConfig().CronosData.WsRpc, conf.GetConfig().CronosData.ContractAddress, conf.GetConfig().CronosData.ChainId)
}

func chainService(wsRpc, contractAddress string, chainId int) {
	client, err := ethclient.Dial(wsRpc)
	if err != nil {
		log.Fatal(err)
	}
	service.OpenTx(client)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	lastConfimedBlock, err := service.GetBlockInfobyChainId(chainId)
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

				go service.SyncBlocks(syncStartNum, block.Number().Uint64(), lastConfimedBlock, contractAddress, wsRpc, chainId)
			}
			go service.FindTx(block, false, contractAddress, chainId)
			firstRun = false
		}
	}
}
