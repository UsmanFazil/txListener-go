package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	logBurnSig         = []byte("Burn(address,uint256,uint256,uint256,bytes32)")
	LogApprovalSig     = []byte("Approval(address,address,uint256)")
	logBurnSigHash     = crypto.Keccak256Hash(logBurnSig)
	logApprovalSigHash = crypto.Keccak256Hash(LogApprovalSig)
)

func OpenTx(client *ethclient.Client) {
	tx, err := GetTxHash()
	if err != nil || len((*tx)) == 0 {
		return
	}

	txHash := common.HexToHash((*tx)[0].TxHash)

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		fmt.Println(err)
	}

	File, err := ioutil.ReadFile("ABI/bridgeABI.json")

	contractAbi, err := abi.JSON(strings.NewReader(string(File)))
	if err != nil {
		log.Fatal(err)
	}

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
				Txhash:      string(vLog.TxHash.String()),
				Address:     burnEvent.Owner.String(),
				Amount:      burnEvent.Amount.String(),
				Tochainid:   vLog.Topics[3].Big().Int64(),
				Fromchainid: vLog.Topics[2].Big().Int64(),
				Status:      "pending",
				Burnid:      hex.EncodeToString(burnEvent.BurnId[:]),
			}

			mysql.SharedStore().AddTxBurnInfo(tx)
		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

		}

	}

}

func GetTxHash() (*[]models.Txhash, error) {
	return mysql.SharedStore().GetTxHash()
}
