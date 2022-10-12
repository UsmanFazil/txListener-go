package models

type Txhash struct {
	TxHash      string `gorm:"column:txhash;primary_key"`
	Blocknum    int
	Contractadd string
}

type Lastconfirmed struct {
	Blocknum int `gorm:"column:blocknum;primary_key"`
}
