package hdwallet_test

import (
	"fmt"
	"log"
	"math/big"

	"github.com/aquachain/hdwallet"

	"gitlab.com/aquachain/aquachain/common"
	"gitlab.com/aquachain/aquachain/core/types"
	"gitlab.com/aquachain/aquachain/params"
)

func ExampleNewFromMnemonic() {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic, "")
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account2, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
	fmt.Println(account2.Address.Hex())
	// Output:
	// 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
	// 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
}

func ExampleWallet_Derive() {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"

	wallet, err := hdwallet.NewFromMnemonic(mnemonic, "")
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account address: %s\n", account.Address.Hex())

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private key in hex: %s\n", privateKey)

	publicKey, _ := wallet.PublicKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Public key in hex: %s\n", publicKey)
	// Output:
	// Account address: 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
	// Private key in hex: 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
	// Public key in hex: 6005c86a6718f66221713a77073c41291cc3abbfcd03aa4955e9b2b50dbf7f9b6672dad0d46ade61e382f79888a73ea7899d9419becf1d6c9ec2087c1188fa18

}

func ExampleNewSeed() {
	seed, _ := hdwallet.NewSeed()
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

func ExampleNewFromMnemonic_sign() {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic, "")
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}

	nonce := uint64(0)
	value := big.NewInt(1000000000000000000)
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := wallet.SignTx(account, tx, nil)
	if err != nil {
		log.Fatal(err)
	}

	chainid := params.MainnetChainConfig.ChainId
	tx2, err := signedTx.AsMessage(types.NewEIP155Signer(chainid))
	if err != nil {
		log.Println(err)
		return
	}
	from := tx2.From()

	fmt.Println(from.Hex())
	fmt.Println(account.Address.Hex())
	// Output:
	// 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
	// 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

}
