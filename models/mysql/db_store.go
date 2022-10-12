package mysql

import (
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

func (s *Store) AddTx(tx *models.Txhash) error {
	return s.db.Create(tx).Error
}

func (s *Store) GetBlockNum() (*models.Lastconfirmed, error) {
	var lastconfirmedNum models.Lastconfirmed
	err := s.db.Raw("SELECT * FROM g_lastconfirmed").Scan(&lastconfirmedNum).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &lastconfirmedNum, err
}

func (s *Store) AddBlockNum(lastconfirmedNum *models.Lastconfirmed) error {
	return s.db.Create(lastconfirmedNum).Error
}

func (s *Store) UpdateBlockNum(newConfimed, oldConfirmed int) (*models.Lastconfirmed, error) {
	var lastconfirmedNum models.Lastconfirmed
	err := s.db.Raw("UPDATE g_lastconfirmed SET blocknum=? WHERE blocknum=?", newConfimed, oldConfirmed).Scan(&lastconfirmedNum).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &lastconfirmedNum, err
}
