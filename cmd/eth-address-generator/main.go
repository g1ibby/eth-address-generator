// Generator ethereum addresses
package main

import (
	"crypto/ecdsa"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/powerman/structlog"
)

//nolint:gochecknoglobals // Main.
var (
	log = structlog.New(structlog.KeyUnit, "main")
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}
	privateKeyString := hexutil.Encode(crypto.FromECDSA(privateKey))[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("invalid public key: %s", err)
	}
	publicKeyString := hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))[4:]
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("Private key: ", privateKeyString)
	fmt.Println("Public key:  ", publicKeyString)
	fmt.Println("Address:     ", address.String())

	path := filepath.Join(".", "addr")
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("failed to create directory: %s", err)
	}

	if err := writeToFile(filepath.Join(path, "privateKey"), privateKeyString); err != nil {
		log.Fatalf("failed to write private key: %s", err)
	}
	if err := writeToFile(filepath.Join(path, "publicKey"), publicKeyString); err != nil {
		log.Fatalf("failed to write public key: %s", err)
	}
	if err := writeToFile(filepath.Join(path, "address"), address.String()); err != nil {
		log.Fatalf("failed to write address: %s", err)
	}
}

func writeToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
