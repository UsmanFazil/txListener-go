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
