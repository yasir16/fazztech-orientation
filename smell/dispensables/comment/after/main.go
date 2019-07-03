package after

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

func getFirstName(user User) string {
	return strings.Split(user.Name, " ")[0]
}

func printAllFemaleFirstName(users []User) {
	for _, user := range users {
		if user.Gender != "female" {
			continue
		}
		fmt.Println(getFirstName(user))
	}
}

func main() {
	users := []User{}
	printAllFemaleFirstName(users)
}
