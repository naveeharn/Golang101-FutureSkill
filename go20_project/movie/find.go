package movie

import "fmt"

func init() {
	fmt.Println("init: find")
}

// package movie
func FindName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avenger: Endgame"
	case "tt3852683":
		return "Black Panther"
	}
	return "not found"
}
