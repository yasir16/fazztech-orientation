package main

import "github.com/mrp130/fazztech-orientation/coupling/data"

func main() {
	wallet := data.NewWallet(1, 1000)

	repo := data.NewService()
	repo.DeleteByUserID(wallet.UserID())
}
