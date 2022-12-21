package service

import (
	"context"
	"fmt"

	"github.com/block-listener/models/mysql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FindTx(block *types.Block, backupSync bool, contractAddress string, chainId int, client *ethclient.Client) {
	openFlag := false
	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			continue
		}

		if tx.To().String() == contractAddress {
			fmt.Println("matched------------------")

			openFlag = true
			SaveTx(tx.Hash().String(), contractAddress, uint(block.Number().Uint64()), chainId)
		}
	}

	SaveLastConfirmed(int(block.Number().Int64()), chainId, backupSync)

	fmt.Println("Block parsed : ", block.Number().Uint64())
	if openFlag {
		go OpenTx(client, chainId)
	}

	go blockConf(block.Number().Int64(), client, chainId)
}

func blockConf(blocknum int64, client *ethclient.Client, chainId int) {

	pendingtx, err := mysql.SharedStore().GetPendingTx(int(blocknum-5), (chainId))
	if pendingtx == nil || err != nil {
		return
	}

	for i := 0; i < len(pendingtx); i++ {
		txHash := common.HexToHash(pendingtx[i].Txhash)

		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		fmt.Println("status reciept ---------------", receipt.Status)
		if receipt.Status == 0 {
			return
		}

		signature := getSignature(pendingtx[i].Originchainid, pendingtx[i].Tochainid, pendingtx[i].Contractadd, pendingtx[i].Burnid, pendingtx[i].Address, pendingtx[i].Amount)
		err = mysql.SharedStore().UpdateTxBurnInfo(pendingtx[i].Txhash, signature)
		if err != nil {
			fmt.Println("error in updating Burn Info:", err)
		}

		if err != nil {
			fmt.Println("error is:", err)
		}
		fmt.Println("receipt:", receipt)
	}
}
