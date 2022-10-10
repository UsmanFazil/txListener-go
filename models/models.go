package models

type Txhash struct {
	TxHash      string `gorm:"column:txhash;primary_key"`
	Blocknum    int
	Contractadd string
}
