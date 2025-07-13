package main

import "fmt"

func add(a int, b int) {
	i := a + b
	println(i)
}

func rdivmod(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	add(10, 30)

	div, mode := rdivmod(10, 30)
	fmt.Printf("div = %d, mode = %d\n", div, mode)
}
