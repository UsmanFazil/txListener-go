package service

import (
	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
	"github.com/pkg/errors"
)

func SaveTx(txhash, contractAdd string, blockNumber uint) (*models.Txhash, error) {

	tx, err := GetTxByHash(txhash)
	if err != nil {
		return nil, err
	}
	if tx != nil {
		return nil, errors.New("tx hash already saved")
	}

	tx = &models.Txhash{
		TxHash:      txhash,
		Blocknum:    int(blockNumber),
		Contractadd: contractAdd,
	}
	return tx, mysql.SharedStore().AddTx(tx)
}

func GetTxByHash(hash string) (*models.Txhash, error) {
	return mysql.SharedStore().GetTxByHash(hash)
}

func SaveLastConfirmed(blockNumber int) (*models.Lastconfirmed, error) {

	blockNum, err := GetLastConfirmedNum()
	if err != nil {
		return nil, err
	}

	if blockNum == nil {
		blockNum = &models.Lastconfirmed{
			Blocknum: int(blockNumber),
		}
		return blockNum, mysql.SharedStore().AddBlockNum(blockNum)
	}

	return mysql.SharedStore().UpdateBlockNum(blockNumber, blockNum.Blocknum)
}

func GetLastConfirmedNum() (*models.Lastconfirmed, error) {
	return mysql.SharedStore().GetBlockNum()
}
