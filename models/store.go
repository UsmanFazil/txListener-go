package models

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error

	//GetConfigs() ([]*Config, error)

	GetTxByHash(txHash string) (*Txhash, error)
	GetTxHash() (*[]Txhash, error)
	AddTx(*Txhash) error

	GetBlockSyncInfo() (*Blocksyncinfo, error)
	GetBlockInfobyChainId(int) (*Blocksyncinfo, error)
	AddBlockSyncInfo(*Blocksyncinfo) error

	UpdateSyncInfo(*Blocksyncinfo) (*Blocksyncinfo, error)
	UpdateSyncStatus(status int) (*Blocksyncinfo, error)

	AddTxMintInfo(*Txmintinfo) error
	AddTxBurnInfo(*Txburninfo) error
}
