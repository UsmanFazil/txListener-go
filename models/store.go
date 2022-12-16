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

	GetBlockSyncInfo() (*Blocksyncinfo, error)
	GetBlockInfobyChainId(int) (*Blocksyncinfo, error)
	AddBlockSyncInfo(*Blocksyncinfo) error

	UpdateSyncInfo(*Blocksyncinfo) (*Blocksyncinfo, error)
	UpdateSyncStatus(status int) (*Blocksyncinfo, error)

	AddTxMintInfo(*Txmintinfo) error
	AddTxBurnInfo(*Txburninfo) error
}
