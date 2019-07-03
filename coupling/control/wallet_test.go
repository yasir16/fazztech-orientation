package control

import "testing"

func TestFilterWallet(t *testing.T) {
	wallets := Wallets{
		Wallet{1, 1000},
		Wallet{2, 2000},
		Wallet{3, 500},
	}

	wallets = wallets.Filter(func(w Wallet) bool {
		return w.balance > 500
	})

	if len(wallets) != 2 {
		t.Errorf("Total wallets after filter must be 2")
	}
}
