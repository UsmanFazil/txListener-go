package service

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/block-listener/conf"
	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SyncBlocks(startBlock, endBlock uint64, blockInfo *models.Blocksyncinfo) {
	mysql.SharedStore().UpdateSyncInfo(blockInfo)
	_, err := mysql.SharedStore().UpdateSyncStatus(0)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(conf.GetConfig().WsRpc)
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(int64(startBlock))
	bigOne := big.NewInt(int64(1))

	for i := startBlock; i <= endBlock; i++ {

		block, err := client.BlockByNumber(ctx, blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		go FindTx(block, true)

		blockNumber.Add(blockNumber, bigOne)
	}

	fmt.Println("All sync threads initiated")
	blockInfo.Syncstatus = 1
	mysql.SharedStore().UpdateSyncInfo(blockInfo)
}
