package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Txhash struct {
	TxHash      string `gorm:"column:txhash;primary_key"`
	Blocknum    int
	Contractadd string
	Chainid     int
	completed   bool
}

type Blocksyncinfo struct {
	Blocksyncnum  int
	Syncstatus    int
	Backupsyncnum int
	Chainid       int
}

type Txmintinfo struct {
	Txhash        string
	Address       string
	Amount        string
	Burnid        string
	Originchainid int64
	Tochainid     int64
	Status        string
}

type Txburninfo struct {
	Txhash        string
	Address       string
	Amount        string
	Signature     string
	Originchainid int64
	Tochainid     int64
	Status        string
	Burnid        string
}

// LogBurn ..
type LogBurn struct {
	Owner         common.Address
	Amount        *big.Int
	OriginChainId *big.Int
	ToChainId     *big.Int
	BurnId        [32]byte
}

// LogMint ..
type LogMint struct {
	Owner         common.Address
	Amount        *big.Int
	OriginChainId *big.Int
	ToChainId     *big.Int
	BurnId        [32]byte
	RefId         [32]byte
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}
