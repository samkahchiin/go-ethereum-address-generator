package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Converting the private key to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// Converting the private key to hexadecimal string using the Encode method
	// hexadecimal = making it shorter by including a-f in the number
	fmt.Println("DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKeyBytes))

	// Generate the public key from the private key
	publicKey := privateKey.Public()
	// ECDSA = elliptic Curve Digital Signature Algorithm
	// Checking the type of Public Key
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key:", hexutil.Encode(publicKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address:", address)
}
