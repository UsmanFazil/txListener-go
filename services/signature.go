package service

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/block-listener/utils"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

var pk = keyGen(utils.GetPK())

func Sign(originChainId, toChainId int64, contractAddress, refId, msgSender, amount string, key *ecdsa.PrivateKey) []byte {
	// Turn the message into a 32-byte hash
	hash := solsha3.SoliditySHA3(
		[]string{"uint256", "uint256", "address", "bytes32", "address", "uint256"},
		[]interface{}{
			originChainId, toChainId, contractAddress, refId, msgSender, amount,
		},
	)
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

func getSignature(originChainId, toChainId int64, contractAddress, refId, msgSender, amount string) string {

	refId = "0x" + refId
	sig := Sign(originChainId, toChainId, contractAddress, refId, msgSender, amount, pk)

	// fmt.Println("address:", hex.EncodeToString(crypto.PubkeyToAddress(pk.PublicKey).Bytes()))
	// fmt.Println("signature:", hex.EncodeToString(sig))

	return hex.EncodeToString(sig)
}

func keyGen(pkString string) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(pkString)

	if err != nil {
		fmt.Println("private key type casting failed")
		panic(err)
	}

	return privateKey
}
