package main

import "fmt"

func paincRoutine() {
	fmt.Println("-> start paincRoutine")

	error := fmt.Errorf("custom")
	// panic will break a application, is like throw Error in other lang
	panic(error)
}

func rescueRoutine() {
	fmt.Println("-> start rescueRoutine")

	rescued := recover()
	if rescued != nil {
		fmt.Println("rescued", rescued)
	}

	fmt.Println("-> end rescueRoutine")

}

func main() {
	// defer schedule functin to run after caller
	defer rescueRoutine()

	fmt.Println("-> start main")
	paincRoutine()
	fmt.Println("-> end main")
}
