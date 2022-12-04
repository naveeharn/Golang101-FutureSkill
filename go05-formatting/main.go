package main

import (
	"fmt"
)

func main() {
	var msg, age, price, check = "Hello Gopher", 55, 22.52, true
	var r rune = ''
	// var r int32 = 'A' // rune alias int32
	fmt.Println(msg, age, price, check)
	fmt.Println(r)
	fmt.Printf("r: %c\n", r)
	fmt.Println()

	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %f\n", price)
	fmt.Printf("check: %t\n", check)
	fmt.Printf("r: %c\n", r)
	fmt.Println()

	fmt.Printf("msg: %#v\n", msg)
	fmt.Printf("age: %#v\n", age)
	fmt.Printf("price: %#v\n", price)
	fmt.Printf("check: %#v\n", check)
	fmt.Printf("r: %c\n", r)
	fmt.Printf("r: %#v\n", r)
	fmt.Println("")

	fmt.Printf("type: %T\tmsg: %#v\n", msg, msg)
	fmt.Printf("type: %T\tage: %#v\n", age, age)
	fmt.Printf("type: %T\tprice: %#v\n", price, price)
	fmt.Printf("type: %T\tcheck: %#v\n", check, check)
	fmt.Printf("type: %T\tr: %c\n", r, r)
	fmt.Printf("type: %T\tr: %#v\n", r, r)
	fmt.Printf("")
}
