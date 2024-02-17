package main

import (
	"fmt"
)

func main() {
	var a int8 = 0

	fmt.Println("operation with a valid size", a+127)

	// overflow
	// fmt.Println("operation with a valid size", a+128)

}
