package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}
