package service

import (
	"fmt"

	"github.com/block-listener/conf"
	"github.com/ethereum/go-ethereum/core/types"
)

func FindTx(block *types.Block, backupSync bool) {

	contractAddress := conf.GetConfig().ContractAddress

	for _, tx := range block.Transactions() {

		if tx.To() == nil {
			continue
		}

		if tx.To().String() == contractAddress {
			SaveTx(tx.Hash().String(), contractAddress, uint(block.Number().Uint64()))
		}
	}
	SaveLastConfirmed(int(block.Number().Int64()), backupSync)

	fmt.Println("Block parsed : ", block.Number().Uint64())
}
