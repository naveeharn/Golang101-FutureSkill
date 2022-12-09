package movie

import "fmt"

func init() {
	fmt.Println("init: movie")
}

// package movie
func Review(name string, rating float64) {
	fmt.Printf("T reviewed %sv and it's rating is %v\n", name, rating)
}
