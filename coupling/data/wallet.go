package data

//Wallet dummy wallet
type Wallet struct {
	userID uint64
	money  float64
}

//NewWallet create new wallet object
func NewWallet(userID uint64, money float64) Wallet {
	return Wallet{
		userID: userID,
		money:  money,
	}
}

//UserID return wallet user ID
func (w Wallet) UserID() uint64 {
	return w.userID
}

//Money return wallet money
func (w Wallet) Money() float64 {
	return w.money
}
