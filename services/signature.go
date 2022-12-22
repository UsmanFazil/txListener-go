package service

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/block-listener/utils"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

var pk = keyGen(utils.GetPK())

func Sign(originChainId, toChainId int64, contractAddress, refId, msgSender, amount string, key *ecdsa.PrivateKey) []byte {
	num := new(big.Int)
	num, ok := num.SetString(amount, 10)
	if !ok {
		fmt.Println("SetString: error")
	}

	// Turn the message into a 32-byte hash
	hash := solsha3.SoliditySHA3(solsha3.Uint256(big.NewInt(originChainId)),
		solsha3.Uint256(big.NewInt(toChainId)),
		solsha3.Address(contractAddress),
		solsha3.Bytes32(refId),
		solsha3.Address(msgSender),
		solsha3.Uint256(num))

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
func KeyGen() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("7cd308f8b5f80233b279079ccf61dfc8ad4b2a7728ed971c2321794621f1f15c")

	if err != nil {
		panic(err)
	}
	return privateKey
}
