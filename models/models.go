package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Txhash struct {
	Id          int `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Txhash      string
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

type GetInfoVo struct {
	Originchainid int64
	Tochainid     int64
	Amount        string
	Transaction   string
	Useraddress   string
	Txtime        uint64
	Status        string
}

type Txburninfo struct {
	Txhash        string `gorm:"column:txhash;primary_key"`
	Txhashid      int
	Address       string
	Amount        string
	Signature     string
	Originchainid int64
	Tochainid     int64
	Status        string
	Txtime        uint64
	Burnid        string
}

type Txpendinginfo struct {
	Txhash        string `gorm:"column:txhash;primary_key"`
	Txhashid      int
	Address       string
	Amount        string
	Signature     string
	Originchainid int64
	Tochainid     int64
	Status        string
	Burnid        string
	Blocknum      int
	Contractadd   string
}

func UserInfoVo(product *Txburninfo) *GetInfoVo {
	return &GetInfoVo{
		Originchainid: product.Originchainid,
		Tochainid:     product.Tochainid,
		Amount:        product.Amount,
		Transaction:   product.Txhash,
		Useraddress:   product.Address,
		Status:        product.Status,
		Txtime:        product.Txtime,
	}
}
