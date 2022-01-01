package fundamentals

import "fmt"

func double(n *int) {
	*n *= 2
}

func appendGreeting(s *string) {
	*s = fmt.Sprintf("Hi, %s", *s)
}
