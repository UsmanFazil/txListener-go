package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/block-listener/models/mysql"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SyncBlocks(startBlock, endBlock uint64, contractAddress, wsRpc string, chainId int, ethClient *ethclient.Client) {

	_, err := mysql.SharedStore().UpdateSyncStatus(0)
	if err != nil {
		fmt.Println("error in syncBlocks", err)
		return
	}

	client, err := ethclient.Dial(wsRpc)
	ctx := context.Background()
	if err != nil {
		fmt.Println("error in syncBlocks", err)
		return
	}

	blockNumber := big.NewInt(int64(startBlock))
	fmt.Println("startBlock:", startBlock, "endBlock", endBlock)

	for i := startBlock; i < endBlock; i++ {

		block, err := client.BlockByNumber(ctx, blockNumber)
		if err != nil {
			fmt.Println("error in syncBlocks", err)
			return
		}

		go FindTx(block, true, contractAddress, chainId, ethClient)
		blockNumber.Add(blockNumber, big.NewInt(int64(1)))
	}

	_, err = mysql.SharedStore().UpdateSyncStatus(1)
	if err != nil {
		fmt.Println("error in syncBlocks", err)
		return
	}

	fmt.Println("Back up sync completed for :", chainId)

}
