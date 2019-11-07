package main

import "fmt"

func main() {

	// with no init values
	// values will be the default of item type
	var a [5]int
	fmt.Println("emp:", a, "capacity:", cap(a), "length:", len(a))

	// set by position
	a[4] = 100
	// access
	fmt.Println("emp:", a[4])

	// b, init values
	b := [5]int{1, 2, 4, 6, 50}
	fmt.Println("emp:", b, "capacity:", cap(b), "length:", len(b))

	// append not will work (only in slice)
	// uncomment and it will not build
	// c := append(b, 3)
	// fmt.Println("emp:", c, "capacity:", cap(c), "length:", len(c))

}
