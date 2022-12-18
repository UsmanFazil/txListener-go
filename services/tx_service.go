package service

import (
	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
	"github.com/pkg/errors"
)

func SaveTx(txhash, contractAdd string, blockNumber uint, chainId int) (*models.Txhash, error) {

	tx, _ := GetTxByHash(txhash)

	if tx != nil {
		return nil, errors.New("tx hash already saved")
	}

	tx = &models.Txhash{
		TxHash:      txhash,
		Blocknum:    int(blockNumber),
		Contractadd: contractAdd,
		Chainid:     chainId,
	}

	return tx, mysql.SharedStore().AddTx(tx)
}

func GetTxByHash(hash string) (*models.Txhash, error) {
	return mysql.SharedStore().GetTxByHash(hash)
}

func SaveLastConfirmed(blockNumber, chainId int, backUpSync bool) (*models.Blocksyncinfo, error) {

	blockNum, err := GetBlockInfobyChainId(chainId)
	if err != nil {
		return nil, err
	}

	if blockNum == nil {
		blockNum = &models.Blocksyncinfo{
			Blocksyncnum:  int(blockNumber),
			Syncstatus:    1,
			Backupsyncnum: 0,
			Chainid:       chainId,
		}
		return blockNum, mysql.SharedStore().AddBlockSyncInfo(blockNum)
	}

	if backUpSync {
		blockNum.Backupsyncnum = blockNumber
	} else {
		blockNum.Blocksyncnum = blockNumber
	}

	// blockNum.Syncstatus = 0
	return mysql.SharedStore().UpdateSyncInfo(blockNum)
}

func GetBlockInfobyChainId(chainId int) (*models.Blocksyncinfo, error) {
	return mysql.SharedStore().GetBlockInfobyChainId(chainId)
}

// func GetBlockSyncInfo() (*models.Blocksyncinfo, error) {
// 	return mysql.SharedStore().GetBlockSyncInfo()
// }
