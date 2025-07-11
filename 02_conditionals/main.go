package main

import (
	"fmt"
)

func main() {
	Helper(20)
}

func Helper(i int) {
	if i > 10 {
		fmt.Println("x > 10")
	} else {
		fmt.Println("x < 10")
	}
}
