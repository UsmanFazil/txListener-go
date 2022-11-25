package conf

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type GbeConfig struct {
	DataSource DataSourceConfig `json:"dataSource"`
	EthData    ChainData        `json:"eth"`
	BscData    ChainData        `json:"bsc"`
	CronosData ChainData        `json:"cronos"`
}

type DataSourceConfig struct {
	DriverName        string `json:"driverName"`
	Addr              string `json:"addr"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
}

type ChainData struct {
	WsRpc           string `json:"ws-rpc"`
	ContractAddress string `json:"bridge-contract"`
	ChainId         int    `json:"chain-id"`
}

var config GbeConfig
var configOnce sync.Once

func GetConfig() *GbeConfig {
	configOnce.Do(func() {
		bytes, err := ioutil.ReadFile("conf.json")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
