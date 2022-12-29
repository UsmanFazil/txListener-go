package main

import (
	"context"
	"fmt"

	"github.com/block-listener/conf"
	service "github.com/block-listener/services"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	go chainService(conf.GetConfig().EthData)
	go chainService(conf.GetConfig().BscData)
	chainService(conf.GetConfig().CronosData)

}

func chainService(chainInfo conf.ChainData) {
	fmt.Println("service started, chain ID: ", chainInfo.ChainId)
	client, err := ethclient.Dial(chainInfo.WsRpc)
	if err != nil {
		chainService(chainInfo)
		return
	}
	service.OpenTx(client, chainInfo.ChainId)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		chainService(chainInfo)
		return
	}
	fmt.Println("ChainId:", chainInfo.ChainId)
	lastConfimedBlock, err := service.GetBlockInfobyChainId(chainInfo.ChainId)
	if err != nil {
		chainService(chainInfo)
		return
	}

	firstRun := true

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("restarting service of chain id", chainInfo.ChainId, "and error is:", err)
			chainService(chainInfo)
		case header := <-headers:
			block, err := client.BlockByNumber(context.Background(), header.Number)
			if err != nil {
				fmt.Println("error in GetBlockByNumber:", err)
				continue
			}
			fmt.Println("New block :", block.Number().Uint64())

			if firstRun && lastConfimedBlock != nil {

				syncStartNum := uint64(lastConfimedBlock.Backupsyncnum)
				if lastConfimedBlock.Syncstatus == 1 {
					syncStartNum = uint64(lastConfimedBlock.Blocksyncnum)
				}

				go service.SyncBlocks(syncStartNum, block.Number().Uint64(), chainInfo.ContractAddress, chainInfo.WsRpc, chainInfo.ChainId, client)
			}

			go service.FindTx(block, false, chainInfo.ContractAddress, chainInfo.ChainId, client)

			firstRun = false
		}
	}
}
