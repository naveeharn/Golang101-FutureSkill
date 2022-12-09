package main

import (
	"fmt"
	"math"
)

func main() {
	show_sum(12.4, 4.2)

	fmt.Printf("type: %T\n", add)
	fmt.Printf("add(4,1) = %#v\n", add(4, 1))

	fmt.Println(sum_diff_mutiple_declaration(4, 1))
	sum, diff := sum_diff_mutiple_declaration(4, 1)
	fmt.Printf("sum = %#v, diff = %#v\n", sum, diff)

	var first, last string = swap("hello", "world")
	fmt.Printf("%#v %#v\n", first, last)

	start, end := between_value_naked_function(0, 1)
	fmt.Printf("[start, end] = [%#v, %#v]\n", start, end)

	fmt.Printf("(3, 4, %#v)\n", pythagorus(3, 4))

	fmt.Printf("compute type: %T\n", compute)
	fmt.Println(compute(sum_diff_mutiple_declaration, 7, 5))

	fmt.Printf("adder type: %T\n", adder)
	fmt.Println(adder)
	fmt.Println(adder())
	// fmt.Printf("adder type: %#v\n", adder)
	inc, double := adder()

	fmt.Println(inc, double)
	fmt.Println(inc(), double())

	// #
	fmt.Printf("multiply type and addr:  %#v\n", multiply)
	fmt.Println(multiply)
	fmt.Printf("multiply 4 * 3 = %#v\n", multiply(4, 3))
}

func show_sum(x float64, y float64) {
	fmt.Printf("%#v + %#v = %#v\n", x, y, x+y)

}

func add(x float64, y float64) float64 {
	return x + y
}

func sum_diff_mutiple_declaration(x float64, y float64) (float64, float64) {
	return x + y, x - y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func between_value_naked_function(center_point float64, distance float64) (start float64, end float64) {
	start = center_point - distance
	end = center_point + distance
	return
}

func pythagorus(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func compute(fn func(first_param, second_param float64) (float64, float64), a, b float64) (x float64, y float64) {
	x, y = fn(a, b)
	return
}

func inc() int {
	return 1
}

func double() int {
	return 2
}

func adder() (func() int, func() int) {
	return inc, double
}

var multiply = func(x float64, y float64) float64 {
	return x * y
}
