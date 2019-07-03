package after

import "time"

type User struct {
	ID        uint64
	Username  string
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Order struct {
	ID        uint64
	UserID    uint64
	ItemID    uint64
	ItemName  string
	Price     float64
	Qty       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func sendReceiptToEmail(user User, order Order) {
	//
}

func main() {
	user := User{}
	order := Order{}
	sendReceiptToEmail(user, order)
}
