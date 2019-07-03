package main

import (
	"fmt"

	"github.com/mrp130/fazztech-orientation/coupling/content"
)

//Wallet aliasing
type Wallet content.Wallet

func deductWallet(w *Wallet, amount float64) {
	w.Balance -= amount
}

func main() {
	wallet := &Wallet{
		UserID:  123,
		Balance: 10000.00,
	}

	deductWallet(wallet, 1000.00)
	fmt.Println(wallet)
}
