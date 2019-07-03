package before

import "time"

type User struct {
	ID        uint64
	Username  string
	Email     string
	Name      string
	Gender    string
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

func sendReceiptToEmail(email string, userName string, itemName string, price float64, qty int, updatedAt time.Time) {
	//
}

func main() {
	user := User{}
	order := Order{}
	sendReceiptToEmail(user.Email, user.Name, order.ItemName, order.Price, order.Qty, order.UpdatedAt)
}
