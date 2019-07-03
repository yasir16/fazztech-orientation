package main

import (
	"github.com/mrp130/fazztech-orientation/coupling/stamp"
)

func main() {
	wallet := stamp.NewWallet(1, 1000)

	repo := stamp.NewService()
	repo.DeleteByUserID(wallet)
}
