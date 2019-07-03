package before

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID        uint64
	Username  string
	Email     string
	Name      string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	users := []User{}

	//Print all female users' first name
	for _, user := range users {
		if user.Gender != "female" {
			continue
		}
		//get first name
		a := strings.Split(user.Name, " ")[0]
		fmt.Println(a)
	}
}
