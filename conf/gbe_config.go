package conf

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type GbeConfig struct {
	DataSource      DataSourceConfig `json:"dataSource"`
	Rpc             string           `json:"ws-rpc"`
	ContractAddress string           `json:"contract-Address"`
}

type DataSourceConfig struct {
	DriverName        string `json:"driverName"`
	Addr              string `json:"addr"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
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
