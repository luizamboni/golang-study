package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("empty:", s)

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)

	// the "slice" operator (as Python)
	l := s[2:5]
	fmt.Println("slice 1:", l)

	l = s[:5]
	fmt.Println("slice 2", l)

	l = s[2:]
	fmt.Println("slice 3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("a string slice:", t)

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
