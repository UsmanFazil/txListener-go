package service

import (
	"context"
	"fmt"
	"math/big"
	"time"

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
	fmt.Println("timeNow------------start", time.Now())

	for i := startBlock; i < endBlock; i++ {

		go FindTx(ctx, blockNumber, true, contractAddress, chainId, client)
		blockNumber.Add(blockNumber, big.NewInt(int64(1)))
	}

	_, err = mysql.SharedStore().UpdateSyncStatus(1)
	if err != nil {
		fmt.Println("error in syncBlocks", err)
		return
	}

	fmt.Println("Back up sync completed for :", chainId)

}
