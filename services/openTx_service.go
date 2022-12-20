package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/block-listener/conf"
	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	logBurnSig     = []byte("Burn(address,uint256,uint256,uint256,bytes32)")
	LogMintSig     = []byte("Mint(address,uint256,uint256,uint256,bytes32,bytes32)")
	logBurnSigHash = crypto.Keccak256Hash(logBurnSig)
	logMintSigHash = crypto.Keccak256Hash(LogMintSig)
)

const path = "ABI/bridgeABI.json"

func OpenTx(client *ethclient.Client, chainId int) {
	tx, err := GetTxHash(chainId)
	if err != nil || len((*tx)) == 0 {
		return
	}

	for i := 0; i < len(*tx); i++ {
		OpenLogs(client, (*tx)[i].Txhash, (*tx)[i].Id)
	}
}

func OpenLogs(client *ethclient.Client, singletxHash string, txhashid int) {

	txHash := common.HexToHash(singletxHash)
	receipt, err := client.TransactionReceipt(context.Background(), txHash)

	if err != nil {
		fmt.Println(err)
	}

	File, err := ioutil.ReadFile(path)

	contractAbi, err := abi.JSON(strings.NewReader(string(File)))
	if err != nil {
		fmt.Println("error in Contract ABI:", err)
		return
	}
	fmt.Println("receipt.Logs:", receipt.Logs)
	for _, vLog := range receipt.Logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		fmt.Printf("Log Value: %s\n", vLog.Topics[0].String())

		switch vLog.Topics[0].Hex() {
		case logBurnSigHash.Hex():
			fmt.Printf("Log Name: Burn\n")

			var burnEvent models.LogBurn
			err := contractAbi.UnpackIntoInterface(&burnEvent, "Burn", vLog.Data)
			if err != nil {
				fmt.Println("error", err)
			}
			burnEvent.Owner = common.HexToAddress(vLog.Topics[1].Hex())

			tx := &models.Txburninfo{
				Txhash:        string(vLog.TxHash.String()),
				Address:       burnEvent.Owner.String(),
				Amount:        burnEvent.Amount.String(),
				Tochainid:     vLog.Topics[3].Big().Int64(),
				Originchainid: vLog.Topics[2].Big().Int64(),
				Status:        "pending",
				Txhashid:      txhashid,
				Burnid:        hex.EncodeToString(burnEvent.BurnId[:]),
			}

			// contractAddress := getContractAddress(tx.Originchainid)
			// tx.Signature = getSignature(tx.Originchainid, tx.Tochainid, contractAddress, tx.Burnid, tx.Address, burnEvent.Amount.String())
			mysql.SharedStore().AddTxBurnInfo(tx)

		case logMintSigHash.Hex():
			fmt.Printf("Log Name: Mint\n")

			var mintEvent models.LogMint
			err := contractAbi.UnpackIntoInterface(&mintEvent, "Mint", vLog.Data)
			if err != nil {
				fmt.Println("error", err)
			}
			mintEvent.Owner = common.HexToAddress(vLog.Topics[1].Hex())
			fmt.Println("burnEvent:", mintEvent)

			tx := &models.Txmintinfo{
				Txhash:        string(vLog.TxHash.String()),
				Address:       mintEvent.Owner.String(),
				Amount:        mintEvent.Amount.String(),
				Tochainid:     vLog.Topics[3].Big().Int64(),
				Originchainid: vLog.Topics[2].Big().Int64(),
				Status:        "pending",
				Burnid:        hex.EncodeToString(mintEvent.BurnId[:]),
			}

			mysql.SharedStore().AddTxMintInfo(tx)
		}
	}

	//mysql.SharedStore().UpdateTxHash(singletxHash, true)

}

func GetTxHash(chainId int) (*[]models.Txhash, error) {
	return mysql.SharedStore().GetTxHash(chainId)
}

func getContractAddress(chainId int64) string {
	var contractAddress string
	if int64(conf.GetConfig().BscData.ChainId) == chainId {
		contractAddress = conf.GetConfig().BscData.ContractAddress
	} else if int64(conf.GetConfig().CronosData.ChainId) == chainId {
		contractAddress = conf.GetConfig().CronosData.ContractAddress
	} else if int64(conf.GetConfig().EthData.ChainId) == chainId {
		contractAddress = conf.GetConfig().EthData.ContractAddress
	}

	return contractAddress
}
