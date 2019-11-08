package main

import "fmt"

// One interface that not implemente any methods is implemented
// by all structs
// because all implements all yours 0 methods
type Any interface {
}

// for example
// this implements Any interface
type Point struct {
	X int
	Y int
}

type Vector struct {
	a Point
	b Point
}

func acceptAnyType(obj Any) {
	fmt.Println(obj)
}

func main() {
	pt := Point{10, 5}
	acceptAnyType(pt)

	vt := Vector{
		Point{1, 5},
		Point{5, 6},
	}

	acceptAnyType(vt)

}
