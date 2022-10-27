package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

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

// LogTransfer ..
type LogTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}
