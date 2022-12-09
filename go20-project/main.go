package main

import (
	"fmt"

	"github.com/naveeharn/golang101-futureskill/tree/master/go20-project/movie"
	"github.com/naveeharn/golang101-futureskill/tree/master/go20-project/ticket"
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
