package data

import "fmt"

//WalletService dummy wallet repo
type WalletService struct{}

//NewService create new wallet repository
func NewService() WalletService {
	return WalletService{}
}

//DeleteByUserID dummy delete
func (r *WalletService) DeleteByUserID(userID uint64) *Wallet {
	fmt.Printf("delete wallet with user ID: %d\n", userID)
	return nil
}
