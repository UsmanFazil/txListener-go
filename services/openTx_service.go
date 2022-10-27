package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/block-listener/models"
	"github.com/block-listener/models/mysql"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// data layer
// logic layer

func OpenTx(client *ethclient.Client) {
	tx, err := GetTxHash()
	if err != nil {
		return
	}
	txHash := common.HexToHash((*tx)[0].TxHash)

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		fmt.Println(err)
	}

	File, err := ioutil.ReadFile("ABI/contractABI.json")

	contractAbi, err := abi.JSON(strings.NewReader(string(File)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range receipt.Logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent models.LogTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			decimal := strconv.FormatInt(vLog.Topics[3].Big().Int64(), 16)
			fmt.Println("decimal:", decimal)
		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent models.LogApproval

			err = contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())

			// call sign service for tx

			// db save tx info + sign + update the status of tx_hash table
		}

		fmt.Printf("\n\n")
	}

}

func GetTxHash() (*[]models.Txhash, error) {
	return mysql.SharedStore().GetTxHash()
}

// functions to listen
// 1- burn tx
// 2- mint tx

// burn(address, amount);
// mint(address, amount);
