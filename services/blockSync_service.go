package service

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SyncBlocks(startBlock, endBlock uint64, blockInfo *models.Blocksyncinfo, contractAddress, wsRpc string, chainId int, ethClient *ethclient.Client) {
	mysql.SharedStore().UpdateSyncInfo(blockInfo)
	_, err := mysql.SharedStore().UpdateSyncStatus(0)
	if err != nil {
		fmt.Println("18 syncBlocks--------------", ethClient)
		log.Fatal(err)
	}

	client, err := ethclient.Dial(wsRpc)
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(int64(startBlock))
	fmt.Println("startBlock:", startBlock, "endBlock", endBlock)

	for i := startBlock; i <= endBlock; i++ {

		block, err := client.BlockByNumber(ctx, blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		go FindTx(block, true, contractAddress, chainId, ethClient)
		blockNumber.Add(blockNumber, big.NewInt(int64(1)))
	}

	fmt.Println("All sync threads initiated")
	blockInfo.Syncstatus = 1
	mysql.SharedStore().UpdateSyncInfo(blockInfo)
}
