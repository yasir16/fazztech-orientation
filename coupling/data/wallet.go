package data

//Wallet dummy wallet
type Wallet struct {
	userID  uint64
	balance float64
}

//NewWallet create new wallet object
func NewWallet(userID uint64, balance float64) Wallet {
	return Wallet{
		userID:  userID,
		balance: balance,
	}
}

//UserID return wallet user ID
func (w Wallet) UserID() uint64 {
	return w.userID
}

//Balance return wallet balance
func (w Wallet) Balance() float64 {
	return w.balance
}
