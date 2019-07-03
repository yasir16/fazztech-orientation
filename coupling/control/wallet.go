package control

//Wallet dummy wallet
type Wallet struct {
	userID uint64
	money  float64
}

//Wallets is array of wallet
type Wallets []Wallet

//FilterFunc used on Filter method
type FilterFunc func(w Wallet) bool

//Filter only take wallet thats true based on filter function
func (wallets Wallets) Filter(f FilterFunc) Wallets {
	var result Wallets
	for _, w := range wallets {
		if f(w) {
			result = append(result, w)
		}
	}
	return result
}
