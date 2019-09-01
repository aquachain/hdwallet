package hdwallet

import (
	"testing"

	"gitlab.com/aquachain/aquachain/aqua/accounts"
)

const testAddress = "0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947"

func testWallet(t *testing.T) (*Wallet, accounts.DerivationPath, accounts.Account) {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := NewFromMnemonic(mnemonic, "")
	if err != nil {
		t.Error(err)
	}

	path, err := ParseDerivationPath("m/44'/60'/0'/0/0")
	if err != nil {
		t.Error(err)
	}

	account, err := wallet.Derive(path, false)
	if err != nil {
		t.Error(err)
	}
	if testAddress != account.Address.Hex() {
		t.Error("wrong address")
	}
	return wallet, path, account
}

func TestIncrement(t *testing.T) {
	w, p, _ := testWallet(t)
	// Output:
	if len(p) == 0 {
		t.Error("len p == 0")
	}
	field := len(p) - 1
	uniques := map[string]bool{}
	testLen := MaxPath // way too slow
	testLen = 1048576  // really slow
	testLen = 8192     // not really good test but w/e
	for i := uint32(0); i < testLen; i++ {
		p[field] = i
		acct, err := w.Derive(p, false)
		if err != nil {
			t.Error(err)
		}
		if uniques[acct.Address.Hex()] {
			t.Error("not unique key")
		}
		uniques[acct.Address.Hex()] = true
		//log.Printf("%d: %s\n", i, acct.Address.Hex())
	}

}
