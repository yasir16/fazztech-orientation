package before

import "fmt"

func foo(s *string) {
	if s == nil || *s == "" {
		fmt.Println("s empty")
	}
	//...
}

func bar(s *string) {
	if s == nil || *s == "" {
		fmt.Println("s empty")
	}
	//...
}
