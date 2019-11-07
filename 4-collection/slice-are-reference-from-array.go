package main

import (
	"fmt"
	"reflect"
)

func main() {
	original := [4]int{1, 3, 5, 3}
	fmt.Println("original:", original, "is a", reflect.TypeOf(original))

	// generated slices are references to original array
	a := original[0:2]
	b := original[2:4]
	fmt.Println("'a' is a", reflect.TypeOf(a))
	fmt.Println("'b' is a", reflect.TypeOf(a))

	fmt.Println(a, b)

	b[0] = 13
	fmt.Println(a, b)
	fmt.Println(original)

	// slice can be created directly by a literal, not specifing a capacity type
	originalSlice := []int{1, 3, 5, 3}
	fmt.Println("'originalSlice' is a", reflect.TypeOf(originalSlice))
}
