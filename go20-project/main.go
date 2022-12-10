package main

import (
	"fmt"

	"github.com/naveeharn/Golang101_FutureSkill/movie"
	"github.com/naveeharn/Golang101_FutureSkill/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 7.7)
	ticket.Buy(movieName)

	fmt.Printf("\n")
	fmt.Println()
}
