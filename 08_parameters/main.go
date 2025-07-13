package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 4
	fmt.Println(n)
}

func doublePtr(n *int) {
	*n *= 4
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10

	double(val)
	fmt.Println("passed by value ==>", val)

	doublePtr(&val)
	fmt.Println("passed by pointer ==>", val)
}
