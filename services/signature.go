package service

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/gin-gonic/gin"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func KeyGen() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("b4f17b38aacf6f4f529e20195438cbdcf360823b9322475ed023e26206fc49f6")

	if err != nil {
		panic(err)
	}

	return privateKey
}

func Sign(message string, key *ecdsa.PrivateKey) []byte {
	// Turn the message into a 32-byte hash
	hash := solsha3.SoliditySHA3(solsha3.String(message))
	// Prefix and then hash to mimic behavior of eth_sign
	prefixed := solsha3.SoliditySHA3(solsha3.String("\x19Ethereum Signed Message:\n32"), solsha3.Bytes32(hash))
	signature, err := secp256k1.Sign(prefixed, math.PaddedBigBytes(key.D, 32))

	if signature[64] == 0 || signature[64] == 1 {
		signature[64] += 27
	}

	if err != nil {
		panic(err)
	}

	return signature
}

func main() {
	router := gin.Default()
	router.GET("/signature/:value", getSignature)
	router.Run("localhost:8081")
}

func getSignature(c *gin.Context) {

	key := KeyGen()
	message := c.Param("value")

	sig := Sign(message, key)

	fmt.Println("address:", hex.EncodeToString(crypto.PubkeyToAddress(key.PublicKey).Bytes()))
	fmt.Println("signature:", hex.EncodeToString(sig))
	// c.IndentedJSON(http.StatusOK, hex.EncodeToString(sig))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "0x" + hex.EncodeToString(sig),
	})
	return
}
