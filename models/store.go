package models

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error

	//GetConfigs() ([]*Config, error)

	GetTxByHash(txHash string) (*Txhash, error)
	AddTx(tx *Txhash) error
	GetBlockNum() (*Lastconfirmed, error)
	AddBlockNum(lastconfirmedNum *Lastconfirmed) error
	UpdateBlockNum(newConfimed, oldConfirmed int) (*Lastconfirmed, error)
}
