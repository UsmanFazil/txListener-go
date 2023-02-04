package models

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error

	//GetConfigs() ([]*Config, error)

	GetTxByHash(txHash string) (*Txhash, error)
	GetTxHash(chainId int) (*[]Txhash, error)
	AddTx(*Txhash) error
	UpdateTxHash(TxHash string, status bool) error
	GetPendingTx(blocknum, chainid int) ([]Txpendinginfo, error)
	GetBlockSyncInfo() (*Blocksyncinfo, error)
	GetBlockInfobyChainId(int) (*Blocksyncinfo, error)
	AddBlockSyncInfo(*Blocksyncinfo) error

	UpdateSyncInfo(*Blocksyncinfo) (*Blocksyncinfo, error)
	UpdateSyncStatus(status int) (*Blocksyncinfo, error)

	AddTxMintInfo(*Txmintinfo) error
	AddTxBurnInfo(*Txburninfo) error
	UpdateTxBurnInfo(TxHash, signature string) error
	UpdateTxBurnInfoMinted(burnid string, chainId int) error
	GetTxBurnInfo(txHash string) (*Txburninfo, error)
	GetTxBurnbyUserAddr(userAddr string) ([]*Txburninfo, error)

	//GetTxMintInfo(txHash string) (*Txmintinfo, error)

	GetTxMintInfoByBurnId(burnId string) (*Txmintinfo, error)
}
