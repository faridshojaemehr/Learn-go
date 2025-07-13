package main

import "fmt"

func main() {
	var x int = 1
	var y int = 2

	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	result := (x / y) / 2.0
	fmt.Printf("result = %v, type of %T\n", result, result)
}
