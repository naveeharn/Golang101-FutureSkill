package ticket

import "fmt"

func init() {
	fmt.Println("init: ticket")
}

// package ticket
func Buy(movie string) {
	fmt.Printf("I bought tickets to %v\n", movie)
}
