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

type Txmintinfo struct {
	Txhash      string
	Address     string
	Amount      *big.Int
	Signature   string
	Fromchainid *big.Int
	Tochainid   *big.Int
	Status      string
}

type Txburninfo struct {
	Txhash      string
	Address     string
	Amount      string
	Signature   string
	Fromchainid int64
	Tochainid   int64
	Status      string
	Burnid      string
}

// LogBurn ..
type LogBurn struct {
	Owner         common.Address
	Amount        *big.Int
	OriginChainId *big.Int
	ToChainId     *big.Int
	BurnId        [32]byte
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}
