package service

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://polygon-mainnet.g.alchemy.com/v2/hONSC7upVBQ-kZ7yXYmudiQV50rsIoCg")
	if err != nil {
		log.Fatal(err)
	}

	// Get the balance of an account
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account balance:", balance) // 25893180161173005034

	// Get the latest known block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest block:", block.Number().Uint64())
}

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func syncBlocks() {
// 	client, err := ethclient.Dial("https://polygon-rpc.com")
// 	if err != nil {
// 		log.Fatal("Whoops something went wrong!", err)
// 	}

// 	ctx := context.Background()
// 	// tx, pending, _ := client.TransactionByHash(ctx, common.HexToHash("0xb561039340d0dc505b69debd2b5d3500b42b55b368a1b081df338b14fec08cae"))
// 	// if !pending {
// 	// 	fmt.Println(tx)
// 	// }

// 	blockNumber := big.NewInt(int64(19317182))
// 	one := big.NewInt(int64(1))

// 	for i := 19317183; i < 31231927; i++ {

// 		blockNumber.Add(blockNumber, one)
// 		block, err := client.BlockByNumber(ctx, blockNumber)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		contractAddress := "0x7d33c606736408D62Bd8e166b7571668d43F0c21"
// 		for _, tx := range block.Transactions() {

// 			if tx.To() == nil {
// 				continue
// 			}

// 			fmt.Println(tx.To())

// 			if tx.To().String() == contractAddress {
// 				signer := types.NewEIP155Signer(tx.ChainId())
// 				sender, err := signer.Sender(tx)
// 				if err != nil {
// 					fmt.Printf("sender: %v", sender.String())
// 				}
// 			}

// 		}

// 		fmt.Println("block Number : ", i)
// 	}

// }
