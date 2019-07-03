package after

import (
	"errors"
	"fmt"
)

type Phone string

func isPhoneNumber(s string) bool {
	return true
}

func NewPhone(s string) (*Phone, error) {
	if !isPhoneNumber(s) {
		return nil, errors.New("invalid phone number")
	}
	p := Phone(s)
	return &p, nil
}

func main() {
	var phone *Phone

	phone, _ = NewPhone("08123123123")

	fmt.Printf("phone number: %s\n", *phone)
}
