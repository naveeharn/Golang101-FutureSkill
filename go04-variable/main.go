package main

import (
	"fmt"
)

// var msg string = "hello"
// var msg = "hello"
// msg := "hello" # can not shortcut for declaring variable

func main() {
	// var msg string = "hello"
	// var msg = "hello"
	msg := "hello"
	fmt.Println(msg)
	msg = "world" + "!"
	fmt.Println(msg)

	// var age int = 14
	age := 14
	fmt.Println("age =", age)

	// var price float32 = 22.14
	var price float64 = 22.14
	fmt.Println("price =", price)

	var checked bool = false
	fmt.Println("checked =", checked)

	// var a, b, c = "a", 1, true
	a, b, c := "a", 1, true
	fmt.Println(a, b, c)
}
