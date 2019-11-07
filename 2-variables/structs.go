package main

import "fmt"

type Address struct {
	Street string
	Number int
}

func main() {

	// create a empty "instance" of struct
	// address1 is created as a pointer
	address1 := new(Address)
	address1.Street = "Benjamin Constant"
	address1.Number = 61
	fmt.Println("address1:", address1)

	// with "positional" parameters
	address2 := Address{"Benjamin Constant", 61}
	fmt.Println("address2:", address2)

	// with "named" parameters
	address3 := Address{Number: 61, Street: "Benjamin Constant"}
	fmt.Println("address3:", address3)

	// assing pointer
	fmt.Println("assign pointer ---------------------------------------------------")

	address4 := address1
	// it changes both
	address4.Number = 50
	fmt.Println("address1:", address1)
	fmt.Println("address4:", address4)

	// assign values
	fmt.Println("assign values ---------------------------------------------------")

	address5 := *address1
	// changes only address5
	address5.Number = 80
	address5.Street = "Outra"
	fmt.Println("address1:", address1)
	fmt.Println("address5:", address5)

	// in slices
	// init a slice of Address type
	addresses := make([]Address, 0)

	addresses = append(addresses, *address1, address2, address3)
}
