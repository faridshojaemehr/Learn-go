package main

import "fmt"

func main() {
	words := []string{"Saba", "hello", "Go", "farid"}
	for word := range words {
		fmt.Println("index => ", word, "-", words[word])
	}

	// single value range
	for i := range words {
		fmt.Println(i)
	}

	// Double value range
	for index, word := range words {
		//fmt.Println(index, word)
		fmt.Printf("%d ==> %s\n", index, word)
	}

	// Double value range ignore index by using
	for _, word := range words {
		fmt.Println(word)
	}

	amir := append(words, "Espoo")
	words = append(words, "Helsinki")
	fmt.Println(words)
	fmt.Println(amir)
}
