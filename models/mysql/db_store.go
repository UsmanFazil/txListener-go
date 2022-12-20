package mysql

import (
	"fmt"

	"github.com/block-listener/models"

	"github.com/jinzhu/gorm"
)

func (s *Store) GetTxByHash(txHash string) (*models.Txhash, error) {
	var tx models.Txhash
	err := s.db.Raw("SELECT * FROM g_txhash WHERE txhash=?", txHash).Scan(&tx).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, err
}

func (s *Store) GetTxBurnbyUserAddr(userAddr string) ([]*models.Txburninfo, error) {
	var tx []*models.Txburninfo
	err := s.db.Raw("SELECT * FROM g_txburninfo WHERE address=?", userAddr).Scan(&tx).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return tx, err
}

func (s *Store) GetTxHash(chainId int) (*[]models.Txhash, error) {
	var tx []models.Txhash
	err := s.db.Raw("SELECT * FROM g_txhash WHERE chainid=? and completed=?", chainId, false).Scan(&tx).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &tx, err
}

func (s *Store) GetTxBurnInfo(txHash string) (*models.Txburninfo, error) {
	var tx models.Txburninfo
	err := s.db.Raw("SELECT * FROM g_txburninfo WHERE txhash=?", txHash).Scan(&tx).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &tx, err
}

func (s *Store) GetPendingTx(blocknum, chainId int) ([]models.Txpendinginfo, error) {
	var tx []models.Txpendinginfo
	err := s.db.Raw("SELECT * FROM g_txburninfo AS a, g_txhash AS b WHERE a.status='pending' AND b.id = a.txhashid AND b.blocknum<=? and b.chainid=?", blocknum, chainId).Scan(&tx).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return tx, err
}

func (s *Store) AddTx(tx *models.Txhash) error {
	return s.db.Create(tx).Error
}

func (s *Store) GetBlockSyncInfo() (*models.Blocksyncinfo, error) {
	var lastconfirmedNum models.Blocksyncinfo
	err := s.db.Raw("SELECT * FROM g_blocksyncinfo").Scan(&lastconfirmedNum).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &lastconfirmedNum, err
}

func (s *Store) GetBlockInfobyChainId(chainId int) (*models.Blocksyncinfo, error) {
	var lastconfirmedNum models.Blocksyncinfo
	err := s.db.Raw("SELECT * FROM g_blocksyncinfo WHERE chainid=?", chainId).Scan(&lastconfirmedNum).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &lastconfirmedNum, err
}

func (s *Store) UpdateTxHash(TxHash string, status bool) error {
	var tx models.Txhash

	err := s.db.Raw("UPDATE g_txhash SET completed=? WHERE txHash=?",
		status, TxHash).Scan(&tx).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	}
	fmt.Println("Updated Transaction:", tx)

	return err
}

func (s *Store) UpdateTxBurnInfo(TxHash, signature string) error {
	var tx models.Txburninfo

	err := s.db.Raw("UPDATE g_txburninfo SET status=?,  signature=? WHERE txhash=?", "signed", signature, TxHash).Scan(&tx).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	}
	fmt.Println("Updated Transaction:", tx)

	return err
}

func (s *Store) AddBlockSyncInfo(lastconfirmedNum *models.Blocksyncinfo) error {
	return s.db.Create(lastconfirmedNum).Error
}

func (s *Store) AddTxMintInfo(dataInfo *models.Txmintinfo) error {
	return s.db.Create(dataInfo).Error
}

func (s *Store) AddTxBurnInfo(dataInfo *models.Txburninfo) error {
	return s.db.Create(dataInfo).Error
}

func (s *Store) UpdateSyncInfo(syncInfo *models.Blocksyncinfo) (*models.Blocksyncinfo, error) {
	err := s.db.Raw("UPDATE g_blocksyncinfo SET blocksyncnum=?,syncstatus=?,backupsyncnum=? WHERE chainid=?",
		syncInfo.Blocksyncnum, syncInfo.Syncstatus, syncInfo.Backupsyncnum, syncInfo.Chainid).Scan(&syncInfo).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return syncInfo, err
}

func (s *Store) UpdateSyncStatus(status int) (*models.Blocksyncinfo, error) {
	var syncInfo models.Blocksyncinfo
	err := s.db.Raw("UPDATE g_blocksyncinfo SET syncstatus = ?", status).Scan(&syncInfo).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &syncInfo, err
}
