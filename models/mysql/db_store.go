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
		return nil, nil
	}
	return &tx, err
}

func (s *Store) GetTxHash() (*[]models.Txhash, error) {
	var tx []models.Txhash
	err := s.db.Raw("SELECT * FROM g_txhash").Scan(&tx).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &tx, err
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

func (s *Store) AddBlockSyncInfo(lastconfirmedNum *models.Blocksyncinfo) error {
	return s.db.Create(lastconfirmedNum).Error
}

func (s *Store) AddTxMintInfo(dataInfo *models.Txmintinfo) error {
	return s.db.Create(dataInfo).Error
}

func (s *Store) AddTxBurnInfo(dataInfo *models.Txburninfo) error {
	fmt.Println("dataInfo:", dataInfo)
	return s.db.Create(dataInfo).Error
}

func (s *Store) UpdateSyncInfo(syncInfo *models.Blocksyncinfo) (*models.Blocksyncinfo, error) {
	err := s.db.Model(&syncInfo).Update(&syncInfo).Error
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
