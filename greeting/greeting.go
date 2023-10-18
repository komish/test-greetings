package greeting

import "fmt"

func SayHello(s string) {
	if s == "" {
		s = "World"
	}
	fmt.Printf("Hello, %s!\n", s)
}
