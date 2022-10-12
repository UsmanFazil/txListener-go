package service

import (
	"context"
	"log"
	"math/big"

	"github.com/block-listener/conf"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SyncBlocks(startBlock, endBlock uint64) {
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

		go FindTx(block)

		blockNumber.Add(blockNumber, bigOne)
	}

}
