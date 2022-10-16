package models

type Txhash struct {
	TxHash      string `gorm:"column:txhash;primary_key"`
	Blocknum    int
	Contractadd string
}

type Blocksyncinfo struct {
	Blocksyncnum  int
	Syncstatus    int
	Backupsyncnum int
}
