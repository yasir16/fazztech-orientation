package communicational

import "errors"

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

//Deduct funct to deduct balance from wallet
func (w Wallet) Deduct(amount float64) (Wallet, error) {
	newAmount := w.balance - amount
	if newAmount < 0 {
		return w, errors.New("insufficient balance")
	}
	w.balance = newAmount
	return w, nil
}

//Add funct to add balanca to wallet
func (w Wallet) Add(amount float64) (Wallet, error) {
	if amount < 0 {
		return w, errors.New("cannot add negative amount. Use deduct function")
	}
	w.balance += amount
	return w, nil
}

//UserID return wallet user ID
func (w Wallet) UserID() uint64 {
	return w.userID
}

//Balance return wallet balance
func (w Wallet) Balance() float64 {
	return w.balance
}
