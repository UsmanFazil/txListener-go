package service

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

func FindTx(block *types.Block, backupSync bool, contractAddress string, chainId int) {

	for _, tx := range block.Transactions() {

		if tx.To() == nil {
			continue
		}

		if tx.To().String() == contractAddress {
			fmt.Println("matched------------------")
			SaveTx(tx.Hash().String(), contractAddress, uint(block.Number().Uint64()), chainId)

		}
	}

	SaveLastConfirmed(int(block.Number().Int64()), chainId, backupSync)

	fmt.Println("Block parsed : ", block.Number().Uint64())

	// go OpenTx()
}
