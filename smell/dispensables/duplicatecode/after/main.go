package after

import "fmt"

func isEmpty(s *string) bool {
	return s == nil || *s == ""
}

func foo(s *string) {
	if isEmpty(s) {
		fmt.Println("s empty")
	}
	//...
}

func bar(s *string) {
	if isEmpty(s) {
		fmt.Println("s empty")
	}
	//...
}
