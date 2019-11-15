package main

import "fmt"

func main() {
	// in array
	ints := [3]int{1, 2}

	for _, value := range ints {
		fmt.Println(value)
	}

	// in slice
	intSlice := []int{1, 2}
	for _, value := range intSlice {
		fmt.Println(value)
	}

	// in map
	map1 := map[string]string{"a": "123", "b": "456"}

	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}

	// in string
	str := "test"
	for _, letter := range str {
		fmt.Println("letter unicode is: ", letter)
	}
}
