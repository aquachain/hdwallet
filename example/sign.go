//+build ignore

package main

import (
	"log"
	"math/big"

	"github.com/aquachain/hdwallet"
	"github.com/davecgh/go-spew/spew"
	"gitlab.com/aquachain/aquachain/common"
	"gitlab.com/aquachain/aquachain/core/types"
)

func main() {
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
	signedTx, err := wallet.SignTx(account, tx, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(signedTx)
}
