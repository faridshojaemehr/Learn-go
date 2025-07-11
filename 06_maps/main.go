package main

import "fmt"

func main() {
	arr := map[string]int{
		"amazon":    200,
		"google":    500,
		"microsoft": 1000,
	}

	fmt.Println(arr["microsoft"])
	fmt.Println(arr["amazon"])
	fmt.Println(arr["google"])

	value, ok := arr["apple"]

	if !ok {
		fmt.Println("No apple")
	} else {
		fmt.Println(value)
	}

	arr["apple"] = 350
	fmt.Println(arr)

	delete(arr, "amazon")
	fmt.Println(arr)

	for key := range arr {
		fmt.Println(key)
	}

	for key, value := range arr {
		fmt.Printf("%s => %d\n", key, value)
	}

}
